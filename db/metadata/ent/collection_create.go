// Code generated by ent, DO NOT EDIT.

package ent

import (
	"Vectory/db/metadata/ent/collection"
	"Vectory/db/metadata/ent/file"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CollectionCreate is the builder for creating a Collection entity.
type CollectionCreate struct {
	config
	mutation *CollectionMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (cc *CollectionCreate) SetName(s string) *CollectionCreate {
	cc.mutation.SetName(s)
	return cc
}

// SetIndexType sets the "index_type" field.
func (cc *CollectionCreate) SetIndexType(s string) *CollectionCreate {
	cc.mutation.SetIndexType(s)
	return cc
}

// SetDataType sets the "data_type" field.
func (cc *CollectionCreate) SetDataType(s string) *CollectionCreate {
	cc.mutation.SetDataType(s)
	return cc
}

// SetEmbedder sets the "embedder" field.
func (cc *CollectionCreate) SetEmbedder(s string) *CollectionCreate {
	cc.mutation.SetEmbedder(s)
	return cc
}

// SetIndexParams sets the "index_params" field.
func (cc *CollectionCreate) SetIndexParams(m map[string]interface{}) *CollectionCreate {
	cc.mutation.SetIndexParams(m)
	return cc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (cc *CollectionCreate) AddFileIDs(ids ...int) *CollectionCreate {
	cc.mutation.AddFileIDs(ids...)
	return cc
}

// AddFiles adds the "files" edges to the File entity.
func (cc *CollectionCreate) AddFiles(f ...*File) *CollectionCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return cc.AddFileIDs(ids...)
}

// Mutation returns the CollectionMutation object of the builder.
func (cc *CollectionCreate) Mutation() *CollectionMutation {
	return cc.mutation
}

// Save creates the Collection in the database.
func (cc *CollectionCreate) Save(ctx context.Context) (*Collection, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CollectionCreate) SaveX(ctx context.Context) *Collection {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CollectionCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CollectionCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CollectionCreate) check() error {
	if _, ok := cc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Collection.name"`)}
	}
	if _, ok := cc.mutation.IndexType(); !ok {
		return &ValidationError{Name: "index_type", err: errors.New(`ent: missing required field "Collection.index_type"`)}
	}
	if _, ok := cc.mutation.DataType(); !ok {
		return &ValidationError{Name: "data_type", err: errors.New(`ent: missing required field "Collection.data_type"`)}
	}
	if _, ok := cc.mutation.Embedder(); !ok {
		return &ValidationError{Name: "embedder", err: errors.New(`ent: missing required field "Collection.embedder"`)}
	}
	if _, ok := cc.mutation.IndexParams(); !ok {
		return &ValidationError{Name: "index_params", err: errors.New(`ent: missing required field "Collection.index_params"`)}
	}
	return nil
}

func (cc *CollectionCreate) sqlSave(ctx context.Context) (*Collection, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CollectionCreate) createSpec() (*Collection, *sqlgraph.CreateSpec) {
	var (
		_node = &Collection{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(collection.Table, sqlgraph.NewFieldSpec(collection.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.Name(); ok {
		_spec.SetField(collection.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := cc.mutation.IndexType(); ok {
		_spec.SetField(collection.FieldIndexType, field.TypeString, value)
		_node.IndexType = value
	}
	if value, ok := cc.mutation.DataType(); ok {
		_spec.SetField(collection.FieldDataType, field.TypeString, value)
		_node.DataType = value
	}
	if value, ok := cc.mutation.Embedder(); ok {
		_spec.SetField(collection.FieldEmbedder, field.TypeString, value)
		_node.Embedder = value
	}
	if value, ok := cc.mutation.IndexParams(); ok {
		_spec.SetField(collection.FieldIndexParams, field.TypeJSON, value)
		_node.IndexParams = value
	}
	if nodes := cc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   collection.FilesTable,
			Columns: []string{collection.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CollectionCreateBulk is the builder for creating many Collection entities in bulk.
type CollectionCreateBulk struct {
	config
	builders []*CollectionCreate
}

// Save creates the Collection entities in the database.
func (ccb *CollectionCreateBulk) Save(ctx context.Context) ([]*Collection, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Collection, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CollectionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CollectionCreateBulk) SaveX(ctx context.Context) []*Collection {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CollectionCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CollectionCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
