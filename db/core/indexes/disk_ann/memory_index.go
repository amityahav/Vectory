package disk_ann

import (
	"Vectory/db/core/indexes/utils"
	"container/heap"
	"encoding/binary"
	"sync"
	"sync/atomic"
)

type MemoryIndex struct {
	sync.RWMutex
	graph             *Graph
	calculateDistance func([]float32, []float32) float32
	deletedObjIds     *sync.Map

	// index size
	size uint32

	// starting point
	s uint32

	// vectors dimension
	dim uint32

	// max vertex degree
	maxDegree uint32

	// the first vertex id that was inserted into the size
	firstId uint32

	readOnly     bool
	snapshotPath string
}

func newMemoryIndex(deletedObjIds *sync.Map, firstId uint32) *MemoryIndex {
	mi := MemoryIndex{
		graph:             nil,
		calculateDistance: nil,
		deletedObjIds:     deletedObjIds,
		firstId:           firstId,
		readOnly:          false,
		snapshotPath:      "",
	}

	return &mi
}

// TODO: beam search support as well
func (mi *MemoryIndex) Search(q []float32, k int, listSize int, onlySearch bool) ([]utils.Element, []utils.Element) {
	// TODO: locking
	sVertex := mi.graph.vertices[mi.s]
	e := utils.Element{
		Id:       int64(mi.s),
		Distance: mi.calculateDistance(sVertex.vector, q),
	}

	if onlySearch {
		e.DataId = sVertex.objId
	}

	resultSet := utils.NewMinMaxHeapFromSlice([]utils.Element{e})
	visited := map[uint32]struct{}{}

	var candidatesVisited []utils.Element

	for resultSet.Len() != 0 {
		min := utils.Pop(resultSet).(utils.Element)

		if _, ok := visited[uint32(min.Id)]; ok {
			continue
		}

		visited[uint32(min.Id)] = struct{}{}

		minVertex := mi.graph.vertices[uint32(min.Id)]

		// filter deleted vertices from the result
		if _, ok := mi.deletedObjIds.Load(minVertex.objId); !ok {
			candidatesVisited = append(candidatesVisited, min)
		}

		for _, n := range minVertex.neighbors {
			nVertex := mi.graph.vertices[n]

			e = utils.Element{
				Id:       int64(n),
				Distance: mi.calculateDistance(nVertex.vector, q),
			}

			if onlySearch {
				e.DataId = nVertex.objId
			}

			utils.Push(resultSet, e)

			for resultSet.Len() > listSize {
				utils.PopMax(resultSet)
			}
		}
	}

	// K-NN is returned when search is performed
	if onlySearch {
		candidatesHeap := utils.NewMinHeapFromSliceDeep(candidatesVisited, len(candidatesVisited))

		knn := make([]utils.Element, 0, k)
		for i := 0; i < k && candidatesHeap.Len() > 0; i++ {
			knn = append(knn, heap.Pop(candidatesHeap).(utils.Element))
		}

		return knn, nil
	}

	return nil, candidatesVisited
}

func (mi *MemoryIndex) Insert(v []float32, listSize int, distanceThreshold int, currId, dataId uint32) error {
	mi.Lock() // TODO: optimize locking
	defer mi.Unlock()

	if mi.readOnly {
		return ErrReadOnlyIndex
	}
	vertex := Vertex{
		id:     currId,
		objId:  dataId,
		vector: v,
	}

	mi.graph.addVertex(&vertex)

	_, visited := mi.Search(v, 1, listSize, false)

	vertex.neighbors = mi.RobustPrune(&vertex, visited, distanceThreshold)

	for _, n := range vertex.neighbors {
		neighbor := mi.graph.vertices[n]

		neighbor.neighbors = append(neighbor.neighbors, vertex.id)

		if len(neighbor.neighbors) > int(mi.maxDegree) {
			distances := make([]utils.Element, 0, len(neighbor.neighbors))

			for _, nn := range neighbor.neighbors {
				nnVertex := mi.graph.vertices[nn]
				distances = append(distances, utils.Element{
					Id:       int64(nn),
					Distance: mi.calculateDistance(nnVertex.vector, neighbor.vector),
				})
			}
			neighbor.neighbors = mi.RobustPrune(neighbor, distances, distanceThreshold)
		}
	}

	atomic.AddUint32(&mi.size, 1)

	return nil
}

func (mi *MemoryIndex) Delete(id uint32) error {
	// TODO: delete consolidation?
	return nil
}

func (mi *MemoryIndex) RobustPrune(v *Vertex, candidates []utils.Element, distanceThreshold int) []uint32 {
	// TODO locking
	deletedCandidates := map[uint32]struct{}{v.id: {}} // excluding vertex v

	for _, n := range v.neighbors {
		nVertex := mi.graph.vertices[n]
		e := utils.Element{
			Id:       int64(n),
			Distance: mi.calculateDistance(v.vector, nVertex.vector),
		}

		candidates = append(candidates, e)
	}

	candidatesHeap := utils.NewMinHeapFromSlice(candidates)
	newNeighbors := make([]uint32, 0, mi.graph.maxDegree)

	for candidatesHeap.Len() != 0 {
		min := heap.Pop(candidatesHeap).(utils.Element)
		if _, ok := deletedCandidates[uint32(min.Id)]; ok {
			continue
		}

		newNeighbors = append(newNeighbors, uint32(min.Id))

		if len(v.neighbors) == int(mi.maxDegree) {
			break
		}

		for _, c := range candidatesHeap.Elements {
			minVertex := mi.graph.vertices[uint32(min.Id)]
			cVertex := mi.graph.vertices[uint32(c.Id)]

			if float32(distanceThreshold)*mi.calculateDistance(minVertex.vector, cVertex.vector) <= c.Distance {
				deletedCandidates[cVertex.id] = struct{}{}
			}
		}
	}

	return newNeighbors
}

func (mi *MemoryIndex) Size() uint32 {
	return mi.size
}

func (mi *MemoryIndex) Snapshot() {

}

func (mi *MemoryIndex) serializeMetadata(buff []byte) int {
	var offset int

	binary.LittleEndian.PutUint32(buff[offset:], mi.dim)
	offset += 4

	binary.LittleEndian.PutUint32(buff[offset:], mi.maxDegree)
	offset += 4

	binary.LittleEndian.PutUint32(buff[offset:], mi.firstId)
	offset += 4

	binary.LittleEndian.PutUint32(buff[offset:], mi.size)
	offset += 4

	return offset
}

func (mi *MemoryIndex) deserializeMetadata(buff []byte) {
	var offset int

	mi.dim = binary.LittleEndian.Uint32(buff[offset:])
	offset += 4

	mi.maxDegree = binary.LittleEndian.Uint32(buff[offset:])
	offset += 4

	mi.firstId = binary.LittleEndian.Uint32(buff[offset:])
	offset += 4

	mi.size = binary.LittleEndian.Uint32(buff[offset:])
	offset += 4

}

func (mi *MemoryIndex) ReadOnly() {
	mi.readOnly = true
}