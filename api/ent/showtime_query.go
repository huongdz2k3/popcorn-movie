// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/ticket"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ShowTimeQuery is the builder for querying ShowTime entities.
type ShowTimeQuery struct {
	config
	ctx         *QueryContext
	order       []showtime.OrderOption
	inters      []Interceptor
	predicates  []predicate.ShowTime
	withRoom    *RoomQuery
	withMovie   *MovieQuery
	withTickets *TicketQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ShowTimeQuery builder.
func (stq *ShowTimeQuery) Where(ps ...predicate.ShowTime) *ShowTimeQuery {
	stq.predicates = append(stq.predicates, ps...)
	return stq
}

// Limit the number of records to be returned by this query.
func (stq *ShowTimeQuery) Limit(limit int) *ShowTimeQuery {
	stq.ctx.Limit = &limit
	return stq
}

// Offset to start from.
func (stq *ShowTimeQuery) Offset(offset int) *ShowTimeQuery {
	stq.ctx.Offset = &offset
	return stq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (stq *ShowTimeQuery) Unique(unique bool) *ShowTimeQuery {
	stq.ctx.Unique = &unique
	return stq
}

// Order specifies how the records should be ordered.
func (stq *ShowTimeQuery) Order(o ...showtime.OrderOption) *ShowTimeQuery {
	stq.order = append(stq.order, o...)
	return stq
}

// QueryRoom chains the current query on the "room" edge.
func (stq *ShowTimeQuery) QueryRoom() *RoomQuery {
	query := (&RoomClient{config: stq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(showtime.Table, showtime.FieldID, selector),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, showtime.RoomTable, showtime.RoomColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryMovie chains the current query on the "movie" edge.
func (stq *ShowTimeQuery) QueryMovie() *MovieQuery {
	query := (&MovieClient{config: stq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(showtime.Table, showtime.FieldID, selector),
			sqlgraph.To(movie.Table, movie.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, showtime.MovieTable, showtime.MovieColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTickets chains the current query on the "tickets" edge.
func (stq *ShowTimeQuery) QueryTickets() *TicketQuery {
	query := (&TicketClient{config: stq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := stq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := stq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(showtime.Table, showtime.FieldID, selector),
			sqlgraph.To(ticket.Table, ticket.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, showtime.TicketsTable, showtime.TicketsColumn),
		)
		fromU = sqlgraph.SetNeighbors(stq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ShowTime entity from the query.
// Returns a *NotFoundError when no ShowTime was found.
func (stq *ShowTimeQuery) First(ctx context.Context) (*ShowTime, error) {
	nodes, err := stq.Limit(1).All(setContextOp(ctx, stq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{showtime.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (stq *ShowTimeQuery) FirstX(ctx context.Context) *ShowTime {
	node, err := stq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ShowTime ID from the query.
// Returns a *NotFoundError when no ShowTime ID was found.
func (stq *ShowTimeQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = stq.Limit(1).IDs(setContextOp(ctx, stq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{showtime.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (stq *ShowTimeQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := stq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ShowTime entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ShowTime entity is found.
// Returns a *NotFoundError when no ShowTime entities are found.
func (stq *ShowTimeQuery) Only(ctx context.Context) (*ShowTime, error) {
	nodes, err := stq.Limit(2).All(setContextOp(ctx, stq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{showtime.Label}
	default:
		return nil, &NotSingularError{showtime.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (stq *ShowTimeQuery) OnlyX(ctx context.Context) *ShowTime {
	node, err := stq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ShowTime ID in the query.
// Returns a *NotSingularError when more than one ShowTime ID is found.
// Returns a *NotFoundError when no entities are found.
func (stq *ShowTimeQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = stq.Limit(2).IDs(setContextOp(ctx, stq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{showtime.Label}
	default:
		err = &NotSingularError{showtime.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (stq *ShowTimeQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := stq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ShowTimes.
func (stq *ShowTimeQuery) All(ctx context.Context) ([]*ShowTime, error) {
	ctx = setContextOp(ctx, stq.ctx, "All")
	if err := stq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ShowTime, *ShowTimeQuery]()
	return withInterceptors[[]*ShowTime](ctx, stq, qr, stq.inters)
}

// AllX is like All, but panics if an error occurs.
func (stq *ShowTimeQuery) AllX(ctx context.Context) []*ShowTime {
	nodes, err := stq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ShowTime IDs.
func (stq *ShowTimeQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if stq.ctx.Unique == nil && stq.path != nil {
		stq.Unique(true)
	}
	ctx = setContextOp(ctx, stq.ctx, "IDs")
	if err = stq.Select(showtime.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (stq *ShowTimeQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := stq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (stq *ShowTimeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, stq.ctx, "Count")
	if err := stq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, stq, querierCount[*ShowTimeQuery](), stq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (stq *ShowTimeQuery) CountX(ctx context.Context) int {
	count, err := stq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (stq *ShowTimeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, stq.ctx, "Exist")
	switch _, err := stq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (stq *ShowTimeQuery) ExistX(ctx context.Context) bool {
	exist, err := stq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ShowTimeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (stq *ShowTimeQuery) Clone() *ShowTimeQuery {
	if stq == nil {
		return nil
	}
	return &ShowTimeQuery{
		config:      stq.config,
		ctx:         stq.ctx.Clone(),
		order:       append([]showtime.OrderOption{}, stq.order...),
		inters:      append([]Interceptor{}, stq.inters...),
		predicates:  append([]predicate.ShowTime{}, stq.predicates...),
		withRoom:    stq.withRoom.Clone(),
		withMovie:   stq.withMovie.Clone(),
		withTickets: stq.withTickets.Clone(),
		// clone intermediate query.
		sql:  stq.sql.Clone(),
		path: stq.path,
	}
}

// WithRoom tells the query-builder to eager-load the nodes that are connected to
// the "room" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShowTimeQuery) WithRoom(opts ...func(*RoomQuery)) *ShowTimeQuery {
	query := (&RoomClient{config: stq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	stq.withRoom = query
	return stq
}

// WithMovie tells the query-builder to eager-load the nodes that are connected to
// the "movie" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShowTimeQuery) WithMovie(opts ...func(*MovieQuery)) *ShowTimeQuery {
	query := (&MovieClient{config: stq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	stq.withMovie = query
	return stq
}

// WithTickets tells the query-builder to eager-load the nodes that are connected to
// the "tickets" edge. The optional arguments are used to configure the query builder of the edge.
func (stq *ShowTimeQuery) WithTickets(opts ...func(*TicketQuery)) *ShowTimeQuery {
	query := (&TicketClient{config: stq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	stq.withTickets = query
	return stq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		StartAt time.Time `json:"start_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ShowTime.Query().
//		GroupBy(showtime.FieldStartAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (stq *ShowTimeQuery) GroupBy(field string, fields ...string) *ShowTimeGroupBy {
	stq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ShowTimeGroupBy{build: stq}
	grbuild.flds = &stq.ctx.Fields
	grbuild.label = showtime.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		StartAt time.Time `json:"start_at,omitempty"`
//	}
//
//	client.ShowTime.Query().
//		Select(showtime.FieldStartAt).
//		Scan(ctx, &v)
func (stq *ShowTimeQuery) Select(fields ...string) *ShowTimeSelect {
	stq.ctx.Fields = append(stq.ctx.Fields, fields...)
	sbuild := &ShowTimeSelect{ShowTimeQuery: stq}
	sbuild.label = showtime.Label
	sbuild.flds, sbuild.scan = &stq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ShowTimeSelect configured with the given aggregations.
func (stq *ShowTimeQuery) Aggregate(fns ...AggregateFunc) *ShowTimeSelect {
	return stq.Select().Aggregate(fns...)
}

func (stq *ShowTimeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range stq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, stq); err != nil {
				return err
			}
		}
	}
	for _, f := range stq.ctx.Fields {
		if !showtime.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if stq.path != nil {
		prev, err := stq.path(ctx)
		if err != nil {
			return err
		}
		stq.sql = prev
	}
	return nil
}

func (stq *ShowTimeQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ShowTime, error) {
	var (
		nodes       = []*ShowTime{}
		_spec       = stq.querySpec()
		loadedTypes = [3]bool{
			stq.withRoom != nil,
			stq.withMovie != nil,
			stq.withTickets != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ShowTime).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ShowTime{config: stq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, stq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := stq.withRoom; query != nil {
		if err := stq.loadRoom(ctx, query, nodes, nil,
			func(n *ShowTime, e *Room) { n.Edges.Room = e }); err != nil {
			return nil, err
		}
	}
	if query := stq.withMovie; query != nil {
		if err := stq.loadMovie(ctx, query, nodes, nil,
			func(n *ShowTime, e *Movie) { n.Edges.Movie = e }); err != nil {
			return nil, err
		}
	}
	if query := stq.withTickets; query != nil {
		if err := stq.loadTickets(ctx, query, nodes,
			func(n *ShowTime) { n.Edges.Tickets = []*Ticket{} },
			func(n *ShowTime, e *Ticket) { n.Edges.Tickets = append(n.Edges.Tickets, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (stq *ShowTimeQuery) loadRoom(ctx context.Context, query *RoomQuery, nodes []*ShowTime, init func(*ShowTime), assign func(*ShowTime, *Room)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShowTime)
	for i := range nodes {
		fk := nodes[i].RoomID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(room.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "room_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (stq *ShowTimeQuery) loadMovie(ctx context.Context, query *MovieQuery, nodes []*ShowTime, init func(*ShowTime), assign func(*ShowTime, *Movie)) error {
	ids := make([]uuid.UUID, 0, len(nodes))
	nodeids := make(map[uuid.UUID][]*ShowTime)
	for i := range nodes {
		fk := nodes[i].MovieID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(movie.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "movie_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (stq *ShowTimeQuery) loadTickets(ctx context.Context, query *TicketQuery, nodes []*ShowTime, init func(*ShowTime), assign func(*ShowTime, *Ticket)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[uuid.UUID]*ShowTime)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(ticket.FieldShowTimeID)
	}
	query.Where(predicate.Ticket(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(showtime.TicketsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.ShowTimeID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "show_time_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (stq *ShowTimeQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := stq.querySpec()
	_spec.Node.Columns = stq.ctx.Fields
	if len(stq.ctx.Fields) > 0 {
		_spec.Unique = stq.ctx.Unique != nil && *stq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, stq.driver, _spec)
}

func (stq *ShowTimeQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(showtime.Table, showtime.Columns, sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID))
	_spec.From = stq.sql
	if unique := stq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if stq.path != nil {
		_spec.Unique = true
	}
	if fields := stq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, showtime.FieldID)
		for i := range fields {
			if fields[i] != showtime.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if stq.withRoom != nil {
			_spec.Node.AddColumnOnce(showtime.FieldRoomID)
		}
		if stq.withMovie != nil {
			_spec.Node.AddColumnOnce(showtime.FieldMovieID)
		}
	}
	if ps := stq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := stq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := stq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := stq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (stq *ShowTimeQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(stq.driver.Dialect())
	t1 := builder.Table(showtime.Table)
	columns := stq.ctx.Fields
	if len(columns) == 0 {
		columns = showtime.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if stq.sql != nil {
		selector = stq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if stq.ctx.Unique != nil && *stq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range stq.predicates {
		p(selector)
	}
	for _, p := range stq.order {
		p(selector)
	}
	if offset := stq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := stq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShowTimeGroupBy is the group-by builder for ShowTime entities.
type ShowTimeGroupBy struct {
	selector
	build *ShowTimeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (stgb *ShowTimeGroupBy) Aggregate(fns ...AggregateFunc) *ShowTimeGroupBy {
	stgb.fns = append(stgb.fns, fns...)
	return stgb
}

// Scan applies the selector query and scans the result into the given value.
func (stgb *ShowTimeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, stgb.build.ctx, "GroupBy")
	if err := stgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShowTimeQuery, *ShowTimeGroupBy](ctx, stgb.build, stgb, stgb.build.inters, v)
}

func (stgb *ShowTimeGroupBy) sqlScan(ctx context.Context, root *ShowTimeQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(stgb.fns))
	for _, fn := range stgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*stgb.flds)+len(stgb.fns))
		for _, f := range *stgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*stgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := stgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ShowTimeSelect is the builder for selecting fields of ShowTime entities.
type ShowTimeSelect struct {
	*ShowTimeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sts *ShowTimeSelect) Aggregate(fns ...AggregateFunc) *ShowTimeSelect {
	sts.fns = append(sts.fns, fns...)
	return sts
}

// Scan applies the selector query and scans the result into the given value.
func (sts *ShowTimeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sts.ctx, "Select")
	if err := sts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ShowTimeQuery, *ShowTimeSelect](ctx, sts.ShowTimeQuery, sts, sts.inters, v)
}

func (sts *ShowTimeSelect) sqlScan(ctx context.Context, root *ShowTimeQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(sts.fns))
	for _, fn := range sts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*sts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
