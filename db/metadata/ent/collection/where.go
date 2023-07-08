// Code generated by ent, DO NOT EDIT.

package collection

import (
	"Vectory/db/metadata/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldName, v))
}

// IndexType applies equality check predicate on the "index_type" field. It's identical to IndexTypeEQ.
func IndexType(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldIndexType, v))
}

// DataType applies equality check predicate on the "data_type" field. It's identical to DataTypeEQ.
func DataType(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDataType, v))
}

// Embedder applies equality check predicate on the "embedder" field. It's identical to EmbedderEQ.
func Embedder(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldEmbedder, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldName, v))
}

// IndexTypeEQ applies the EQ predicate on the "index_type" field.
func IndexTypeEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldIndexType, v))
}

// IndexTypeNEQ applies the NEQ predicate on the "index_type" field.
func IndexTypeNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldIndexType, v))
}

// IndexTypeIn applies the In predicate on the "index_type" field.
func IndexTypeIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldIndexType, vs...))
}

// IndexTypeNotIn applies the NotIn predicate on the "index_type" field.
func IndexTypeNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldIndexType, vs...))
}

// IndexTypeGT applies the GT predicate on the "index_type" field.
func IndexTypeGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldIndexType, v))
}

// IndexTypeGTE applies the GTE predicate on the "index_type" field.
func IndexTypeGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldIndexType, v))
}

// IndexTypeLT applies the LT predicate on the "index_type" field.
func IndexTypeLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldIndexType, v))
}

// IndexTypeLTE applies the LTE predicate on the "index_type" field.
func IndexTypeLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldIndexType, v))
}

// IndexTypeContains applies the Contains predicate on the "index_type" field.
func IndexTypeContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldIndexType, v))
}

// IndexTypeHasPrefix applies the HasPrefix predicate on the "index_type" field.
func IndexTypeHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldIndexType, v))
}

// IndexTypeHasSuffix applies the HasSuffix predicate on the "index_type" field.
func IndexTypeHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldIndexType, v))
}

// IndexTypeEqualFold applies the EqualFold predicate on the "index_type" field.
func IndexTypeEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldIndexType, v))
}

// IndexTypeContainsFold applies the ContainsFold predicate on the "index_type" field.
func IndexTypeContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldIndexType, v))
}

// DataTypeEQ applies the EQ predicate on the "data_type" field.
func DataTypeEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldDataType, v))
}

// DataTypeNEQ applies the NEQ predicate on the "data_type" field.
func DataTypeNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldDataType, v))
}

// DataTypeIn applies the In predicate on the "data_type" field.
func DataTypeIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldDataType, vs...))
}

// DataTypeNotIn applies the NotIn predicate on the "data_type" field.
func DataTypeNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldDataType, vs...))
}

// DataTypeGT applies the GT predicate on the "data_type" field.
func DataTypeGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldDataType, v))
}

// DataTypeGTE applies the GTE predicate on the "data_type" field.
func DataTypeGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldDataType, v))
}

// DataTypeLT applies the LT predicate on the "data_type" field.
func DataTypeLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldDataType, v))
}

// DataTypeLTE applies the LTE predicate on the "data_type" field.
func DataTypeLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldDataType, v))
}

// DataTypeContains applies the Contains predicate on the "data_type" field.
func DataTypeContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldDataType, v))
}

// DataTypeHasPrefix applies the HasPrefix predicate on the "data_type" field.
func DataTypeHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldDataType, v))
}

// DataTypeHasSuffix applies the HasSuffix predicate on the "data_type" field.
func DataTypeHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldDataType, v))
}

// DataTypeEqualFold applies the EqualFold predicate on the "data_type" field.
func DataTypeEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldDataType, v))
}

// DataTypeContainsFold applies the ContainsFold predicate on the "data_type" field.
func DataTypeContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldDataType, v))
}

// EmbedderEQ applies the EQ predicate on the "embedder" field.
func EmbedderEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEQ(FieldEmbedder, v))
}

// EmbedderNEQ applies the NEQ predicate on the "embedder" field.
func EmbedderNEQ(v string) predicate.Collection {
	return predicate.Collection(sql.FieldNEQ(FieldEmbedder, v))
}

// EmbedderIn applies the In predicate on the "embedder" field.
func EmbedderIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldIn(FieldEmbedder, vs...))
}

// EmbedderNotIn applies the NotIn predicate on the "embedder" field.
func EmbedderNotIn(vs ...string) predicate.Collection {
	return predicate.Collection(sql.FieldNotIn(FieldEmbedder, vs...))
}

// EmbedderGT applies the GT predicate on the "embedder" field.
func EmbedderGT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGT(FieldEmbedder, v))
}

// EmbedderGTE applies the GTE predicate on the "embedder" field.
func EmbedderGTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldGTE(FieldEmbedder, v))
}

// EmbedderLT applies the LT predicate on the "embedder" field.
func EmbedderLT(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLT(FieldEmbedder, v))
}

// EmbedderLTE applies the LTE predicate on the "embedder" field.
func EmbedderLTE(v string) predicate.Collection {
	return predicate.Collection(sql.FieldLTE(FieldEmbedder, v))
}

// EmbedderContains applies the Contains predicate on the "embedder" field.
func EmbedderContains(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContains(FieldEmbedder, v))
}

// EmbedderHasPrefix applies the HasPrefix predicate on the "embedder" field.
func EmbedderHasPrefix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasPrefix(FieldEmbedder, v))
}

// EmbedderHasSuffix applies the HasSuffix predicate on the "embedder" field.
func EmbedderHasSuffix(v string) predicate.Collection {
	return predicate.Collection(sql.FieldHasSuffix(FieldEmbedder, v))
}

// EmbedderEqualFold applies the EqualFold predicate on the "embedder" field.
func EmbedderEqualFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldEqualFold(FieldEmbedder, v))
}

// EmbedderContainsFold applies the ContainsFold predicate on the "embedder" field.
func EmbedderContainsFold(v string) predicate.Collection {
	return predicate.Collection(sql.FieldContainsFold(FieldEmbedder, v))
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Collection) predicate.Collection {
	return predicate.Collection(func(s *sql.Selector) {
		p(s.Not())
	})
}
