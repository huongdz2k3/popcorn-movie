// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/seat"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/ent/transaction"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TicketCreate is the builder for creating a Ticket entity.
type TicketCreate struct {
	config
	mutation *TicketMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetIsBooked sets the "is_booked" field.
func (tc *TicketCreate) SetIsBooked(b bool) *TicketCreate {
	tc.mutation.SetIsBooked(b)
	return tc
}

// SetNillableIsBooked sets the "is_booked" field if the given value is not nil.
func (tc *TicketCreate) SetNillableIsBooked(b *bool) *TicketCreate {
	if b != nil {
		tc.SetIsBooked(*b)
	}
	return tc
}

// SetPrice sets the "price" field.
func (tc *TicketCreate) SetPrice(f float64) *TicketCreate {
	tc.mutation.SetPrice(f)
	return tc
}

// SetTransactionID sets the "transaction_id" field.
func (tc *TicketCreate) SetTransactionID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetTransactionID(u)
	return tc
}

// SetNillableTransactionID sets the "transaction_id" field if the given value is not nil.
func (tc *TicketCreate) SetNillableTransactionID(u *uuid.UUID) *TicketCreate {
	if u != nil {
		tc.SetTransactionID(*u)
	}
	return tc
}

// SetSeatID sets the "seat_id" field.
func (tc *TicketCreate) SetSeatID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetSeatID(u)
	return tc
}

// SetShowTimeID sets the "show_time_id" field.
func (tc *TicketCreate) SetShowTimeID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetShowTimeID(u)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TicketCreate) SetCreatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TicketCreate) SetNillableCreatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// SetUpdatedAt sets the "updated_at" field.
func (tc *TicketCreate) SetUpdatedAt(t time.Time) *TicketCreate {
	tc.mutation.SetUpdatedAt(t)
	return tc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (tc *TicketCreate) SetNillableUpdatedAt(t *time.Time) *TicketCreate {
	if t != nil {
		tc.SetUpdatedAt(*t)
	}
	return tc
}

// SetID sets the "id" field.
func (tc *TicketCreate) SetID(u uuid.UUID) *TicketCreate {
	tc.mutation.SetID(u)
	return tc
}

// SetTransaction sets the "transaction" edge to the Transaction entity.
func (tc *TicketCreate) SetTransaction(t *Transaction) *TicketCreate {
	return tc.SetTransactionID(t.ID)
}

// SetSeat sets the "seat" edge to the Seat entity.
func (tc *TicketCreate) SetSeat(s *Seat) *TicketCreate {
	return tc.SetSeatID(s.ID)
}

// SetShowTime sets the "show_time" edge to the ShowTime entity.
func (tc *TicketCreate) SetShowTime(s *ShowTime) *TicketCreate {
	return tc.SetShowTimeID(s.ID)
}

// Mutation returns the TicketMutation object of the builder.
func (tc *TicketCreate) Mutation() *TicketMutation {
	return tc.mutation
}

// Save creates the Ticket in the database.
func (tc *TicketCreate) Save(ctx context.Context) (*Ticket, error) {
	tc.defaults()
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TicketCreate) SaveX(ctx context.Context) *Ticket {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TicketCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TicketCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TicketCreate) defaults() {
	if _, ok := tc.mutation.IsBooked(); !ok {
		v := ticket.DefaultIsBooked
		tc.mutation.SetIsBooked(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := ticket.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		v := ticket.DefaultUpdatedAt()
		tc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TicketCreate) check() error {
	if _, ok := tc.mutation.IsBooked(); !ok {
		return &ValidationError{Name: "is_booked", err: errors.New(`ent: missing required field "Ticket.is_booked"`)}
	}
	if _, ok := tc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Ticket.price"`)}
	}
	if v, ok := tc.mutation.Price(); ok {
		if err := ticket.PriceValidator(v); err != nil {
			return &ValidationError{Name: "price", err: fmt.Errorf(`ent: validator failed for field "Ticket.price": %w`, err)}
		}
	}
	if _, ok := tc.mutation.SeatID(); !ok {
		return &ValidationError{Name: "seat_id", err: errors.New(`ent: missing required field "Ticket.seat_id"`)}
	}
	if _, ok := tc.mutation.ShowTimeID(); !ok {
		return &ValidationError{Name: "show_time_id", err: errors.New(`ent: missing required field "Ticket.show_time_id"`)}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Ticket.created_at"`)}
	}
	if _, ok := tc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Ticket.updated_at"`)}
	}
	if _, ok := tc.mutation.SeatID(); !ok {
		return &ValidationError{Name: "seat", err: errors.New(`ent: missing required edge "Ticket.seat"`)}
	}
	if _, ok := tc.mutation.ShowTimeID(); !ok {
		return &ValidationError{Name: "show_time", err: errors.New(`ent: missing required edge "Ticket.show_time"`)}
	}
	return nil
}

func (tc *TicketCreate) sqlSave(ctx context.Context) (*Ticket, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TicketCreate) createSpec() (*Ticket, *sqlgraph.CreateSpec) {
	var (
		_node = &Ticket{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(ticket.Table, sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID))
	)
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.IsBooked(); ok {
		_spec.SetField(ticket.FieldIsBooked, field.TypeBool, value)
		_node.IsBooked = value
	}
	if value, ok := tc.mutation.Price(); ok {
		_spec.SetField(ticket.FieldPrice, field.TypeFloat64, value)
		_node.Price = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.SetField(ticket.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := tc.mutation.UpdatedAt(); ok {
		_spec.SetField(ticket.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if nodes := tc.mutation.TransactionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.TransactionTable,
			Columns: []string{ticket.TransactionColumn},
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
	if nodes := tc.mutation.SeatIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.SeatTable,
			Columns: []string{ticket.SeatColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(seat.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SeatID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := tc.mutation.ShowTimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   ticket.ShowTimeTable,
			Columns: []string{ticket.ShowTimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ShowTimeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Ticket.Create().
//		SetIsBooked(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TicketUpsert) {
//			SetIsBooked(v+v).
//		}).
//		Exec(ctx)
func (tc *TicketCreate) OnConflict(opts ...sql.ConflictOption) *TicketUpsertOne {
	tc.conflict = opts
	return &TicketUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Ticket.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TicketCreate) OnConflictColumns(columns ...string) *TicketUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TicketUpsertOne{
		create: tc,
	}
}

type (
	// TicketUpsertOne is the builder for "upsert"-ing
	//  one Ticket node.
	TicketUpsertOne struct {
		create *TicketCreate
	}

	// TicketUpsert is the "OnConflict" setter.
	TicketUpsert struct {
		*sql.UpdateSet
	}
)

// SetIsBooked sets the "is_booked" field.
func (u *TicketUpsert) SetIsBooked(v bool) *TicketUpsert {
	u.Set(ticket.FieldIsBooked, v)
	return u
}

// UpdateIsBooked sets the "is_booked" field to the value that was provided on create.
func (u *TicketUpsert) UpdateIsBooked() *TicketUpsert {
	u.SetExcluded(ticket.FieldIsBooked)
	return u
}

// SetPrice sets the "price" field.
func (u *TicketUpsert) SetPrice(v float64) *TicketUpsert {
	u.Set(ticket.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *TicketUpsert) UpdatePrice() *TicketUpsert {
	u.SetExcluded(ticket.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *TicketUpsert) AddPrice(v float64) *TicketUpsert {
	u.Add(ticket.FieldPrice, v)
	return u
}

// SetTransactionID sets the "transaction_id" field.
func (u *TicketUpsert) SetTransactionID(v uuid.UUID) *TicketUpsert {
	u.Set(ticket.FieldTransactionID, v)
	return u
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TicketUpsert) UpdateTransactionID() *TicketUpsert {
	u.SetExcluded(ticket.FieldTransactionID)
	return u
}

// ClearTransactionID clears the value of the "transaction_id" field.
func (u *TicketUpsert) ClearTransactionID() *TicketUpsert {
	u.SetNull(ticket.FieldTransactionID)
	return u
}

// SetSeatID sets the "seat_id" field.
func (u *TicketUpsert) SetSeatID(v uuid.UUID) *TicketUpsert {
	u.Set(ticket.FieldSeatID, v)
	return u
}

// UpdateSeatID sets the "seat_id" field to the value that was provided on create.
func (u *TicketUpsert) UpdateSeatID() *TicketUpsert {
	u.SetExcluded(ticket.FieldSeatID)
	return u
}

// SetShowTimeID sets the "show_time_id" field.
func (u *TicketUpsert) SetShowTimeID(v uuid.UUID) *TicketUpsert {
	u.Set(ticket.FieldShowTimeID, v)
	return u
}

// UpdateShowTimeID sets the "show_time_id" field to the value that was provided on create.
func (u *TicketUpsert) UpdateShowTimeID() *TicketUpsert {
	u.SetExcluded(ticket.FieldShowTimeID)
	return u
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TicketUpsert) SetUpdatedAt(v time.Time) *TicketUpsert {
	u.Set(ticket.FieldUpdatedAt, v)
	return u
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TicketUpsert) UpdateUpdatedAt() *TicketUpsert {
	u.SetExcluded(ticket.FieldUpdatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Ticket.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ticket.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TicketUpsertOne) UpdateNewValues() *TicketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(ticket.FieldID)
		}
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(ticket.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Ticket.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TicketUpsertOne) Ignore() *TicketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TicketUpsertOne) DoNothing() *TicketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TicketCreate.OnConflict
// documentation for more info.
func (u *TicketUpsertOne) Update(set func(*TicketUpsert)) *TicketUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TicketUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsBooked sets the "is_booked" field.
func (u *TicketUpsertOne) SetIsBooked(v bool) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetIsBooked(v)
	})
}

// UpdateIsBooked sets the "is_booked" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdateIsBooked() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateIsBooked()
	})
}

// SetPrice sets the "price" field.
func (u *TicketUpsertOne) SetPrice(v float64) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *TicketUpsertOne) AddPrice(v float64) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdatePrice() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdatePrice()
	})
}

// SetTransactionID sets the "transaction_id" field.
func (u *TicketUpsertOne) SetTransactionID(v uuid.UUID) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetTransactionID(v)
	})
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdateTransactionID() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateTransactionID()
	})
}

// ClearTransactionID clears the value of the "transaction_id" field.
func (u *TicketUpsertOne) ClearTransactionID() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.ClearTransactionID()
	})
}

// SetSeatID sets the "seat_id" field.
func (u *TicketUpsertOne) SetSeatID(v uuid.UUID) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetSeatID(v)
	})
}

// UpdateSeatID sets the "seat_id" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdateSeatID() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateSeatID()
	})
}

// SetShowTimeID sets the "show_time_id" field.
func (u *TicketUpsertOne) SetShowTimeID(v uuid.UUID) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetShowTimeID(v)
	})
}

// UpdateShowTimeID sets the "show_time_id" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdateShowTimeID() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateShowTimeID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TicketUpsertOne) SetUpdatedAt(v time.Time) *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TicketUpsertOne) UpdateUpdatedAt() *TicketUpsertOne {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TicketUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TicketCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TicketUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TicketUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: TicketUpsertOne.ID is not supported by MySQL driver. Use TicketUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TicketUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TicketCreateBulk is the builder for creating many Ticket entities in bulk.
type TicketCreateBulk struct {
	config
	err      error
	builders []*TicketCreate
	conflict []sql.ConflictOption
}

// Save creates the Ticket entities in the database.
func (tcb *TicketCreateBulk) Save(ctx context.Context) ([]*Ticket, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Ticket, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TicketMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TicketCreateBulk) SaveX(ctx context.Context) []*Ticket {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TicketCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TicketCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Ticket.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TicketUpsert) {
//			SetIsBooked(v+v).
//		}).
//		Exec(ctx)
func (tcb *TicketCreateBulk) OnConflict(opts ...sql.ConflictOption) *TicketUpsertBulk {
	tcb.conflict = opts
	return &TicketUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Ticket.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TicketCreateBulk) OnConflictColumns(columns ...string) *TicketUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TicketUpsertBulk{
		create: tcb,
	}
}

// TicketUpsertBulk is the builder for "upsert"-ing
// a bulk of Ticket nodes.
type TicketUpsertBulk struct {
	create *TicketCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Ticket.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(ticket.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TicketUpsertBulk) UpdateNewValues() *TicketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(ticket.FieldID)
			}
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(ticket.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Ticket.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TicketUpsertBulk) Ignore() *TicketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TicketUpsertBulk) DoNothing() *TicketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TicketCreateBulk.OnConflict
// documentation for more info.
func (u *TicketUpsertBulk) Update(set func(*TicketUpsert)) *TicketUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TicketUpsert{UpdateSet: update})
	}))
	return u
}

// SetIsBooked sets the "is_booked" field.
func (u *TicketUpsertBulk) SetIsBooked(v bool) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetIsBooked(v)
	})
}

// UpdateIsBooked sets the "is_booked" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdateIsBooked() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateIsBooked()
	})
}

// SetPrice sets the "price" field.
func (u *TicketUpsertBulk) SetPrice(v float64) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *TicketUpsertBulk) AddPrice(v float64) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdatePrice() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdatePrice()
	})
}

// SetTransactionID sets the "transaction_id" field.
func (u *TicketUpsertBulk) SetTransactionID(v uuid.UUID) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetTransactionID(v)
	})
}

// UpdateTransactionID sets the "transaction_id" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdateTransactionID() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateTransactionID()
	})
}

// ClearTransactionID clears the value of the "transaction_id" field.
func (u *TicketUpsertBulk) ClearTransactionID() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.ClearTransactionID()
	})
}

// SetSeatID sets the "seat_id" field.
func (u *TicketUpsertBulk) SetSeatID(v uuid.UUID) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetSeatID(v)
	})
}

// UpdateSeatID sets the "seat_id" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdateSeatID() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateSeatID()
	})
}

// SetShowTimeID sets the "show_time_id" field.
func (u *TicketUpsertBulk) SetShowTimeID(v uuid.UUID) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetShowTimeID(v)
	})
}

// UpdateShowTimeID sets the "show_time_id" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdateShowTimeID() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateShowTimeID()
	})
}

// SetUpdatedAt sets the "updated_at" field.
func (u *TicketUpsertBulk) SetUpdatedAt(v time.Time) *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.SetUpdatedAt(v)
	})
}

// UpdateUpdatedAt sets the "updated_at" field to the value that was provided on create.
func (u *TicketUpsertBulk) UpdateUpdatedAt() *TicketUpsertBulk {
	return u.Update(func(s *TicketUpsert) {
		s.UpdateUpdatedAt()
	})
}

// Exec executes the query.
func (u *TicketUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TicketCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TicketCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TicketUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
