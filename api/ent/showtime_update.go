// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ShowTimeUpdate is the builder for updating ShowTime entities.
type ShowTimeUpdate struct {
	config
	hooks    []Hook
	mutation *ShowTimeMutation
}

// Where appends a list predicates to the ShowTimeUpdate builder.
func (stu *ShowTimeUpdate) Where(ps ...predicate.ShowTime) *ShowTimeUpdate {
	stu.mutation.Where(ps...)
	return stu
}

// SetStartAt sets the "start_at" field.
func (stu *ShowTimeUpdate) SetStartAt(t time.Time) *ShowTimeUpdate {
	stu.mutation.SetStartAt(t)
	return stu
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (stu *ShowTimeUpdate) SetNillableStartAt(t *time.Time) *ShowTimeUpdate {
	if t != nil {
		stu.SetStartAt(*t)
	}
	return stu
}

// SetEndAt sets the "end_at" field.
func (stu *ShowTimeUpdate) SetEndAt(t time.Time) *ShowTimeUpdate {
	stu.mutation.SetEndAt(t)
	return stu
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (stu *ShowTimeUpdate) SetNillableEndAt(t *time.Time) *ShowTimeUpdate {
	if t != nil {
		stu.SetEndAt(*t)
	}
	return stu
}

// SetMovieID sets the "movie_id" field.
func (stu *ShowTimeUpdate) SetMovieID(u uuid.UUID) *ShowTimeUpdate {
	stu.mutation.SetMovieID(u)
	return stu
}

// SetNillableMovieID sets the "movie_id" field if the given value is not nil.
func (stu *ShowTimeUpdate) SetNillableMovieID(u *uuid.UUID) *ShowTimeUpdate {
	if u != nil {
		stu.SetMovieID(*u)
	}
	return stu
}

// SetRoomID sets the "room_id" field.
func (stu *ShowTimeUpdate) SetRoomID(u uuid.UUID) *ShowTimeUpdate {
	stu.mutation.SetRoomID(u)
	return stu
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (stu *ShowTimeUpdate) SetNillableRoomID(u *uuid.UUID) *ShowTimeUpdate {
	if u != nil {
		stu.SetRoomID(*u)
	}
	return stu
}

// SetUpdatedAt sets the "updated_at" field.
func (stu *ShowTimeUpdate) SetUpdatedAt(t time.Time) *ShowTimeUpdate {
	stu.mutation.SetUpdatedAt(t)
	return stu
}

// SetRoom sets the "room" edge to the Room entity.
func (stu *ShowTimeUpdate) SetRoom(r *Room) *ShowTimeUpdate {
	return stu.SetRoomID(r.ID)
}

// SetMovie sets the "movie" edge to the Movie entity.
func (stu *ShowTimeUpdate) SetMovie(m *Movie) *ShowTimeUpdate {
	return stu.SetMovieID(m.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (stu *ShowTimeUpdate) AddTicketIDs(ids ...uuid.UUID) *ShowTimeUpdate {
	stu.mutation.AddTicketIDs(ids...)
	return stu
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (stu *ShowTimeUpdate) AddTickets(t ...*Ticket) *ShowTimeUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stu.AddTicketIDs(ids...)
}

// Mutation returns the ShowTimeMutation object of the builder.
func (stu *ShowTimeUpdate) Mutation() *ShowTimeMutation {
	return stu.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (stu *ShowTimeUpdate) ClearRoom() *ShowTimeUpdate {
	stu.mutation.ClearRoom()
	return stu
}

// ClearMovie clears the "movie" edge to the Movie entity.
func (stu *ShowTimeUpdate) ClearMovie() *ShowTimeUpdate {
	stu.mutation.ClearMovie()
	return stu
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (stu *ShowTimeUpdate) ClearTickets() *ShowTimeUpdate {
	stu.mutation.ClearTickets()
	return stu
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (stu *ShowTimeUpdate) RemoveTicketIDs(ids ...uuid.UUID) *ShowTimeUpdate {
	stu.mutation.RemoveTicketIDs(ids...)
	return stu
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (stu *ShowTimeUpdate) RemoveTickets(t ...*Ticket) *ShowTimeUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stu.RemoveTicketIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (stu *ShowTimeUpdate) Save(ctx context.Context) (int, error) {
	stu.defaults()
	return withHooks(ctx, stu.sqlSave, stu.mutation, stu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stu *ShowTimeUpdate) SaveX(ctx context.Context) int {
	affected, err := stu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (stu *ShowTimeUpdate) Exec(ctx context.Context) error {
	_, err := stu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stu *ShowTimeUpdate) ExecX(ctx context.Context) {
	if err := stu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stu *ShowTimeUpdate) defaults() {
	if _, ok := stu.mutation.UpdatedAt(); !ok {
		v := showtime.UpdateDefaultUpdatedAt()
		stu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stu *ShowTimeUpdate) check() error {
	if _, ok := stu.mutation.RoomID(); stu.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShowTime.room"`)
	}
	if _, ok := stu.mutation.MovieID(); stu.mutation.MovieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShowTime.movie"`)
	}
	return nil
}

func (stu *ShowTimeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := stu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(showtime.Table, showtime.Columns, sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID))
	if ps := stu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stu.mutation.StartAt(); ok {
		_spec.SetField(showtime.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := stu.mutation.EndAt(); ok {
		_spec.SetField(showtime.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := stu.mutation.UpdatedAt(); ok {
		_spec.SetField(showtime.FieldUpdatedAt, field.TypeTime, value)
	}
	if stu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.RoomTable,
			Columns: []string{showtime.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.RoomTable,
			Columns: []string{showtime.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stu.mutation.MovieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.MovieTable,
			Columns: []string{showtime.MovieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.MovieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.MovieTable,
			Columns: []string{showtime.MovieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stu.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !stu.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stu.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, stu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{showtime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	stu.mutation.done = true
	return n, nil
}

// ShowTimeUpdateOne is the builder for updating a single ShowTime entity.
type ShowTimeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ShowTimeMutation
}

// SetStartAt sets the "start_at" field.
func (stuo *ShowTimeUpdateOne) SetStartAt(t time.Time) *ShowTimeUpdateOne {
	stuo.mutation.SetStartAt(t)
	return stuo
}

// SetNillableStartAt sets the "start_at" field if the given value is not nil.
func (stuo *ShowTimeUpdateOne) SetNillableStartAt(t *time.Time) *ShowTimeUpdateOne {
	if t != nil {
		stuo.SetStartAt(*t)
	}
	return stuo
}

// SetEndAt sets the "end_at" field.
func (stuo *ShowTimeUpdateOne) SetEndAt(t time.Time) *ShowTimeUpdateOne {
	stuo.mutation.SetEndAt(t)
	return stuo
}

// SetNillableEndAt sets the "end_at" field if the given value is not nil.
func (stuo *ShowTimeUpdateOne) SetNillableEndAt(t *time.Time) *ShowTimeUpdateOne {
	if t != nil {
		stuo.SetEndAt(*t)
	}
	return stuo
}

// SetMovieID sets the "movie_id" field.
func (stuo *ShowTimeUpdateOne) SetMovieID(u uuid.UUID) *ShowTimeUpdateOne {
	stuo.mutation.SetMovieID(u)
	return stuo
}

// SetNillableMovieID sets the "movie_id" field if the given value is not nil.
func (stuo *ShowTimeUpdateOne) SetNillableMovieID(u *uuid.UUID) *ShowTimeUpdateOne {
	if u != nil {
		stuo.SetMovieID(*u)
	}
	return stuo
}

// SetRoomID sets the "room_id" field.
func (stuo *ShowTimeUpdateOne) SetRoomID(u uuid.UUID) *ShowTimeUpdateOne {
	stuo.mutation.SetRoomID(u)
	return stuo
}

// SetNillableRoomID sets the "room_id" field if the given value is not nil.
func (stuo *ShowTimeUpdateOne) SetNillableRoomID(u *uuid.UUID) *ShowTimeUpdateOne {
	if u != nil {
		stuo.SetRoomID(*u)
	}
	return stuo
}

// SetUpdatedAt sets the "updated_at" field.
func (stuo *ShowTimeUpdateOne) SetUpdatedAt(t time.Time) *ShowTimeUpdateOne {
	stuo.mutation.SetUpdatedAt(t)
	return stuo
}

// SetRoom sets the "room" edge to the Room entity.
func (stuo *ShowTimeUpdateOne) SetRoom(r *Room) *ShowTimeUpdateOne {
	return stuo.SetRoomID(r.ID)
}

// SetMovie sets the "movie" edge to the Movie entity.
func (stuo *ShowTimeUpdateOne) SetMovie(m *Movie) *ShowTimeUpdateOne {
	return stuo.SetMovieID(m.ID)
}

// AddTicketIDs adds the "tickets" edge to the Ticket entity by IDs.
func (stuo *ShowTimeUpdateOne) AddTicketIDs(ids ...uuid.UUID) *ShowTimeUpdateOne {
	stuo.mutation.AddTicketIDs(ids...)
	return stuo
}

// AddTickets adds the "tickets" edges to the Ticket entity.
func (stuo *ShowTimeUpdateOne) AddTickets(t ...*Ticket) *ShowTimeUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stuo.AddTicketIDs(ids...)
}

// Mutation returns the ShowTimeMutation object of the builder.
func (stuo *ShowTimeUpdateOne) Mutation() *ShowTimeMutation {
	return stuo.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (stuo *ShowTimeUpdateOne) ClearRoom() *ShowTimeUpdateOne {
	stuo.mutation.ClearRoom()
	return stuo
}

// ClearMovie clears the "movie" edge to the Movie entity.
func (stuo *ShowTimeUpdateOne) ClearMovie() *ShowTimeUpdateOne {
	stuo.mutation.ClearMovie()
	return stuo
}

// ClearTickets clears all "tickets" edges to the Ticket entity.
func (stuo *ShowTimeUpdateOne) ClearTickets() *ShowTimeUpdateOne {
	stuo.mutation.ClearTickets()
	return stuo
}

// RemoveTicketIDs removes the "tickets" edge to Ticket entities by IDs.
func (stuo *ShowTimeUpdateOne) RemoveTicketIDs(ids ...uuid.UUID) *ShowTimeUpdateOne {
	stuo.mutation.RemoveTicketIDs(ids...)
	return stuo
}

// RemoveTickets removes "tickets" edges to Ticket entities.
func (stuo *ShowTimeUpdateOne) RemoveTickets(t ...*Ticket) *ShowTimeUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return stuo.RemoveTicketIDs(ids...)
}

// Where appends a list predicates to the ShowTimeUpdate builder.
func (stuo *ShowTimeUpdateOne) Where(ps ...predicate.ShowTime) *ShowTimeUpdateOne {
	stuo.mutation.Where(ps...)
	return stuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (stuo *ShowTimeUpdateOne) Select(field string, fields ...string) *ShowTimeUpdateOne {
	stuo.fields = append([]string{field}, fields...)
	return stuo
}

// Save executes the query and returns the updated ShowTime entity.
func (stuo *ShowTimeUpdateOne) Save(ctx context.Context) (*ShowTime, error) {
	stuo.defaults()
	return withHooks(ctx, stuo.sqlSave, stuo.mutation, stuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (stuo *ShowTimeUpdateOne) SaveX(ctx context.Context) *ShowTime {
	node, err := stuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (stuo *ShowTimeUpdateOne) Exec(ctx context.Context) error {
	_, err := stuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (stuo *ShowTimeUpdateOne) ExecX(ctx context.Context) {
	if err := stuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (stuo *ShowTimeUpdateOne) defaults() {
	if _, ok := stuo.mutation.UpdatedAt(); !ok {
		v := showtime.UpdateDefaultUpdatedAt()
		stuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (stuo *ShowTimeUpdateOne) check() error {
	if _, ok := stuo.mutation.RoomID(); stuo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShowTime.room"`)
	}
	if _, ok := stuo.mutation.MovieID(); stuo.mutation.MovieCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "ShowTime.movie"`)
	}
	return nil
}

func (stuo *ShowTimeUpdateOne) sqlSave(ctx context.Context) (_node *ShowTime, err error) {
	if err := stuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(showtime.Table, showtime.Columns, sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID))
	id, ok := stuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ShowTime.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := stuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, showtime.FieldID)
		for _, f := range fields {
			if !showtime.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != showtime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := stuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := stuo.mutation.StartAt(); ok {
		_spec.SetField(showtime.FieldStartAt, field.TypeTime, value)
	}
	if value, ok := stuo.mutation.EndAt(); ok {
		_spec.SetField(showtime.FieldEndAt, field.TypeTime, value)
	}
	if value, ok := stuo.mutation.UpdatedAt(); ok {
		_spec.SetField(showtime.FieldUpdatedAt, field.TypeTime, value)
	}
	if stuo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.RoomTable,
			Columns: []string{showtime.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.RoomTable,
			Columns: []string{showtime.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stuo.mutation.MovieCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.MovieTable,
			Columns: []string{showtime.MovieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.MovieIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   showtime.MovieTable,
			Columns: []string{showtime.MovieColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if stuo.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.RemovedTicketsIDs(); len(nodes) > 0 && !stuo.mutation.TicketsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := stuo.mutation.TicketsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   showtime.TicketsTable,
			Columns: []string{showtime.TicketsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ticket.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ShowTime{config: stuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, stuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{showtime.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	stuo.mutation.done = true
	return _node, nil
}
