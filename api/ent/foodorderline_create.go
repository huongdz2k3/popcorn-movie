// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/food"
	"PopcornMovie/ent/foodorderline"
	"PopcornMovie/ent/transaction"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// FoodOrderLineCreate is the builder for creating a FoodOrderLine entity.
type FoodOrderLineCreate struct {
	config
	mutation *FoodOrderLineMutation
	hooks    []Hook
}

// SetQuantity sets the "quantity" field.
func (folc *FoodOrderLineCreate) SetQuantity(i int) *FoodOrderLineCreate {
	folc.mutation.SetQuantity(i)
	return folc
}

// SetFoodID sets the "food_id" field.
func (folc *FoodOrderLineCreate) SetFoodID(u uuid.UUID) *FoodOrderLineCreate {
	folc.mutation.SetFoodID(u)
	return folc
}

// SetTransactionID sets the "transaction_id" field.
func (folc *FoodOrderLineCreate) SetTransactionID(u uuid.UUID) *FoodOrderLineCreate {
	folc.mutation.SetTransactionID(u)
	return folc
}

// SetCreatedAt sets the "created_at" field.
func (folc *FoodOrderLineCreate) SetCreatedAt(t time.Time) *FoodOrderLineCreate {
	folc.mutation.SetCreatedAt(t)
	return folc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (folc *FoodOrderLineCreate) SetNillableCreatedAt(t *time.Time) *FoodOrderLineCreate {
	if t != nil {
		folc.SetCreatedAt(*t)
	}
	return folc
}

// SetUpdatedAt sets the "updated_at" field.
func (folc *FoodOrderLineCreate) SetUpdatedAt(t time.Time) *FoodOrderLineCreate {
	folc.mutation.SetUpdatedAt(t)
	return folc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (folc *FoodOrderLineCreate) SetNillableUpdatedAt(t *time.Time) *FoodOrderLineCreate {
	if t != nil {
		folc.SetUpdatedAt(*t)
	}
	return folc
}

// SetID sets the "id" field.
func (folc *FoodOrderLineCreate) SetID(u uuid.UUID) *FoodOrderLineCreate {
	folc.mutation.SetID(u)
	return folc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (folc *FoodOrderLineCreate) SetNillableID(u *uuid.UUID) *FoodOrderLineCreate {
	if u != nil {
		folc.SetID(*u)
	}
	return folc
}

// SetFood sets the "food" edge to the Food entity.
func (folc *FoodOrderLineCreate) SetFood(f *Food) *FoodOrderLineCreate {
	return folc.SetFoodID(f.ID)
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (folc *FoodOrderLineCreate) SetTransaction(t *Transaction) *FoodOrderLineCreate {
	return folc.SetTransactionID(t.ID)
}

// Mutation returns the FoodOrderLineMutation object of the builder.
func (folc *FoodOrderLineCreate) Mutation() *FoodOrderLineMutation {
	return folc.mutation
}

// Save creates the FoodOrderLine in the database.
func (folc *FoodOrderLineCreate) Save(ctx context.Context) (*FoodOrderLine, error) {
	folc.defaults()
	return withHooks(ctx, folc.sqlSave, folc.mutation, folc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (folc *FoodOrderLineCreate) SaveX(ctx context.Context) *FoodOrderLine {
	v, err := folc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (folc *FoodOrderLineCreate) Exec(ctx context.Context) error {
	_, err := folc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (folc *FoodOrderLineCreate) ExecX(ctx context.Context) {
	if err := folc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (folc *FoodOrderLineCreate) defaults() {
	if _, ok := folc.mutation.CreatedAt(); !ok {
		v := foodorderline.DefaultCreatedAt()
		folc.mutation.SetCreatedAt(v)
	}
	if _, ok := folc.mutation.UpdatedAt(); !ok {
		v := foodorderline.DefaultUpdatedAt()
		folc.mutation.SetUpdatedAt(v)
	}
	if _, ok := folc.mutation.ID(); !ok {
		v := foodorderline.DefaultID()
		folc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (folc *FoodOrderLineCreate) check() error {
	if _, ok := folc.mutation.Quantity(); !ok {
		return &ValidationError{Name: "quantity", err: errors.New(`ent: missing required field "FoodOrderLine.quantity"`)}
	}
	if _, ok := folc.mutation.FoodID(); !ok {
		return &ValidationError{Name: "food_id", err: errors.New(`ent: missing required field "FoodOrderLine.food_id"`)}
	}
	if _, ok := folc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction_id", err: errors.New(`ent: missing required field "FoodOrderLine.transaction_id"`)}
	}
	if _, ok := folc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "FoodOrderLine.created_at"`)}
	}
	if _, ok := folc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "FoodOrderLine.updated_at"`)}
	}
	if _, ok := folc.mutation.FoodID(); !ok {
		return &ValidationError{Name: "food", err: errors.New(`ent: missing required edge "FoodOrderLine.food"`)}
	}
	if _, ok := folc.mutation.TransactionID(); !ok {
		return &ValidationError{Name: "transaction", err: errors.New(`ent: missing required edge "FoodOrderLine.transaction"`)}
	}
	return nil
}

func (folc *FoodOrderLineCreate) sqlSave(ctx context.Context) (*FoodOrderLine, error) {
	if err := folc.check(); err != nil {
		return nil, err
	}
	_node, _spec := folc.createSpec()
	if err := sqlgraph.CreateNode(ctx, folc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	folc.mutation.id = &_node.ID
	folc.mutation.done = true
	return _node, nil
}

func (folc *FoodOrderLineCreate) createSpec() (*FoodOrderLine, *sqlgraph.CreateSpec) {
	var (
		_node = &FoodOrderLine{config: folc.config}
		_spec = sqlgraph.NewCreateSpec(foodorderline.Table, sqlgraph.NewFieldSpec(foodorderline.FieldID, field.TypeUUID))
	)
	if id, ok := folc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := folc.mutation.Quantity(); ok {
		_spec.SetField(foodorderline.FieldQuantity, field.TypeInt, value)
		_node.Quantity = value
	}
	if value, ok := folc.mutation.CreatedAt(); ok {
		_spec.SetField(foodorderline.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := folc.mutation.UpdatedAt(); ok {
		_spec.SetField(foodorderline.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := folc.mutation.FoodIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.FoodTable,
			Columns: []string{foodorderline.FoodColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(food.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.FoodID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := folc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   foodorderline.TransactionTable,
			Columns: []string{foodorderline.TransactionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(transaction.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TransactionID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FoodOrderLineCreateBulk is the builder for creating many FoodOrderLine entities in bulk.
type FoodOrderLineCreateBulk struct {
	config
	err      error
	builders []*FoodOrderLineCreate
}

// Save creates the FoodOrderLine entities in the database.
func (folcb *FoodOrderLineCreateBulk) Save(ctx context.Context) ([]*FoodOrderLine, error) {
	if folcb.err != nil {
		return nil, folcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(folcb.builders))
	nodes := make([]*FoodOrderLine, len(folcb.builders))
	mutators := make([]Mutator, len(folcb.builders))
	for i := range folcb.builders {
		func(i int, root context.Context) {
			builder := folcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FoodOrderLineMutation)
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
					_, err = mutators[i+1].Mutate(root, folcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, folcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, folcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (folcb *FoodOrderLineCreateBulk) SaveX(ctx context.Context) []*FoodOrderLine {
	v, err := folcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (folcb *FoodOrderLineCreateBulk) Exec(ctx context.Context) error {
	_, err := folcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (folcb *FoodOrderLineCreateBulk) ExecX(ctx context.Context) {
	if err := folcb.Exec(ctx); err != nil {
		panic(err)
	}
}
