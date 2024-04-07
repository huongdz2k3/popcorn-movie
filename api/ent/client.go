// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"PopcornMovie/ent/migrate"

	"PopcornMovie/ent/room"
	"PopcornMovie/ent/session"
	"PopcornMovie/ent/theater"
	"PopcornMovie/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Room is the client for interacting with the Room builders.
	Room *RoomClient
	// Session is the client for interacting with the Session builders.
	Session *SessionClient
	// Theater is the client for interacting with the Theater builders.
	Theater *TheaterClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Room = NewRoomClient(c.config)
	c.Session = NewSessionClient(c.config)
	c.Theater = NewTheaterClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Room:    NewRoomClient(cfg),
		Session: NewSessionClient(cfg),
		Theater: NewTheaterClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Room:    NewRoomClient(cfg),
		Session: NewSessionClient(cfg),
		Theater: NewTheaterClient(cfg),
		User:    NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Room.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Room.Use(hooks...)
	c.Session.Use(hooks...)
	c.Theater.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Room.Intercept(interceptors...)
	c.Session.Intercept(interceptors...)
	c.Theater.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *RoomMutation:
		return c.Room.mutate(ctx, m)
	case *SessionMutation:
		return c.Session.mutate(ctx, m)
	case *TheaterMutation:
		return c.Theater.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// RoomClient is a client for the Room schema.
type RoomClient struct {
	config
}

// NewRoomClient returns a client for the Room from the given config.
func NewRoomClient(c config) *RoomClient {
	return &RoomClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `room.Hooks(f(g(h())))`.
func (c *RoomClient) Use(hooks ...Hook) {
	c.hooks.Room = append(c.hooks.Room, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `room.Intercept(f(g(h())))`.
func (c *RoomClient) Intercept(interceptors ...Interceptor) {
	c.inters.Room = append(c.inters.Room, interceptors...)
}

// Create returns a builder for creating a Room entity.
func (c *RoomClient) Create() *RoomCreate {
	mutation := newRoomMutation(c.config, OpCreate)
	return &RoomCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Room entities.
func (c *RoomClient) CreateBulk(builders ...*RoomCreate) *RoomCreateBulk {
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RoomClient) MapCreateBulk(slice any, setFunc func(*RoomCreate, int)) *RoomCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RoomCreateBulk{err: fmt.Errorf("calling to RoomClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RoomCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RoomCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Room.
func (c *RoomClient) Update() *RoomUpdate {
	mutation := newRoomMutation(c.config, OpUpdate)
	return &RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RoomClient) UpdateOne(r *Room) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoom(r))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RoomClient) UpdateOneID(id int) *RoomUpdateOne {
	mutation := newRoomMutation(c.config, OpUpdateOne, withRoomID(id))
	return &RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Room.
func (c *RoomClient) Delete() *RoomDelete {
	mutation := newRoomMutation(c.config, OpDelete)
	return &RoomDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RoomClient) DeleteOne(r *Room) *RoomDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RoomClient) DeleteOneID(id int) *RoomDeleteOne {
	builder := c.Delete().Where(room.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RoomDeleteOne{builder}
}

// Query returns a query builder for Room.
func (c *RoomClient) Query() *RoomQuery {
	return &RoomQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRoom},
		inters: c.Interceptors(),
	}
}

// Get returns a Room entity by its id.
func (c *RoomClient) Get(ctx context.Context, id int) (*Room, error) {
	return c.Query().Where(room.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RoomClient) GetX(ctx context.Context, id int) *Room {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTheater queries the theater edge of a Room.
func (c *RoomClient) QueryTheater(r *Room) *TheaterQuery {
	query := (&TheaterClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(room.Table, room.FieldID, id),
			sqlgraph.To(theater.Table, theater.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, room.TheaterTable, room.TheaterColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RoomClient) Hooks() []Hook {
	return c.hooks.Room
}

// Interceptors returns the client interceptors.
func (c *RoomClient) Interceptors() []Interceptor {
	return c.inters.Room
}

func (c *RoomClient) mutate(ctx context.Context, m *RoomMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RoomCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RoomUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RoomUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RoomDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Room mutation op: %q", m.Op())
	}
}

// SessionClient is a client for the Session schema.
type SessionClient struct {
	config
}

// NewSessionClient returns a client for the Session from the given config.
func NewSessionClient(c config) *SessionClient {
	return &SessionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `session.Hooks(f(g(h())))`.
func (c *SessionClient) Use(hooks ...Hook) {
	c.hooks.Session = append(c.hooks.Session, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `session.Intercept(f(g(h())))`.
func (c *SessionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Session = append(c.inters.Session, interceptors...)
}

// Create returns a builder for creating a Session entity.
func (c *SessionClient) Create() *SessionCreate {
	mutation := newSessionMutation(c.config, OpCreate)
	return &SessionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Session entities.
func (c *SessionClient) CreateBulk(builders ...*SessionCreate) *SessionCreateBulk {
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *SessionClient) MapCreateBulk(slice any, setFunc func(*SessionCreate, int)) *SessionCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &SessionCreateBulk{err: fmt.Errorf("calling to SessionClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*SessionCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &SessionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Session.
func (c *SessionClient) Update() *SessionUpdate {
	mutation := newSessionMutation(c.config, OpUpdate)
	return &SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SessionClient) UpdateOne(s *Session) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSession(s))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SessionClient) UpdateOneID(id uuid.UUID) *SessionUpdateOne {
	mutation := newSessionMutation(c.config, OpUpdateOne, withSessionID(id))
	return &SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Session.
func (c *SessionClient) Delete() *SessionDelete {
	mutation := newSessionMutation(c.config, OpDelete)
	return &SessionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SessionClient) DeleteOne(s *Session) *SessionDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SessionClient) DeleteOneID(id uuid.UUID) *SessionDeleteOne {
	builder := c.Delete().Where(session.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SessionDeleteOne{builder}
}

// Query returns a query builder for Session.
func (c *SessionClient) Query() *SessionQuery {
	return &SessionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSession},
		inters: c.Interceptors(),
	}
}

// Get returns a Session entity by its id.
func (c *SessionClient) Get(ctx context.Context, id uuid.UUID) (*Session, error) {
	return c.Query().Where(session.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SessionClient) GetX(ctx context.Context, id uuid.UUID) *Session {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *SessionClient) Hooks() []Hook {
	return c.hooks.Session
}

// Interceptors returns the client interceptors.
func (c *SessionClient) Interceptors() []Interceptor {
	return c.inters.Session
}

func (c *SessionClient) mutate(ctx context.Context, m *SessionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SessionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SessionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SessionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SessionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Session mutation op: %q", m.Op())
	}
}

// TheaterClient is a client for the Theater schema.
type TheaterClient struct {
	config
}

// NewTheaterClient returns a client for the Theater from the given config.
func NewTheaterClient(c config) *TheaterClient {
	return &TheaterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `theater.Hooks(f(g(h())))`.
func (c *TheaterClient) Use(hooks ...Hook) {
	c.hooks.Theater = append(c.hooks.Theater, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `theater.Intercept(f(g(h())))`.
func (c *TheaterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Theater = append(c.inters.Theater, interceptors...)
}

// Create returns a builder for creating a Theater entity.
func (c *TheaterClient) Create() *TheaterCreate {
	mutation := newTheaterMutation(c.config, OpCreate)
	return &TheaterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Theater entities.
func (c *TheaterClient) CreateBulk(builders ...*TheaterCreate) *TheaterCreateBulk {
	return &TheaterCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TheaterClient) MapCreateBulk(slice any, setFunc func(*TheaterCreate, int)) *TheaterCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TheaterCreateBulk{err: fmt.Errorf("calling to TheaterClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TheaterCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TheaterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Theater.
func (c *TheaterClient) Update() *TheaterUpdate {
	mutation := newTheaterMutation(c.config, OpUpdate)
	return &TheaterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TheaterClient) UpdateOne(t *Theater) *TheaterUpdateOne {
	mutation := newTheaterMutation(c.config, OpUpdateOne, withTheater(t))
	return &TheaterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TheaterClient) UpdateOneID(id uuid.UUID) *TheaterUpdateOne {
	mutation := newTheaterMutation(c.config, OpUpdateOne, withTheaterID(id))
	return &TheaterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Theater.
func (c *TheaterClient) Delete() *TheaterDelete {
	mutation := newTheaterMutation(c.config, OpDelete)
	return &TheaterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TheaterClient) DeleteOne(t *Theater) *TheaterDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TheaterClient) DeleteOneID(id uuid.UUID) *TheaterDeleteOne {
	builder := c.Delete().Where(theater.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TheaterDeleteOne{builder}
}

// Query returns a query builder for Theater.
func (c *TheaterClient) Query() *TheaterQuery {
	return &TheaterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTheater},
		inters: c.Interceptors(),
	}
}

// Get returns a Theater entity by its id.
func (c *TheaterClient) Get(ctx context.Context, id uuid.UUID) (*Theater, error) {
	return c.Query().Where(theater.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TheaterClient) GetX(ctx context.Context, id uuid.UUID) *Theater {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRooms queries the rooms edge of a Theater.
func (c *TheaterClient) QueryRooms(t *Theater) *RoomQuery {
	query := (&RoomClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(theater.Table, theater.FieldID, id),
			sqlgraph.To(room.Table, room.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, theater.RoomsTable, theater.RoomsColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TheaterClient) Hooks() []Hook {
	return c.hooks.Theater
}

// Interceptors returns the client interceptors.
func (c *TheaterClient) Interceptors() []Interceptor {
	return c.inters.Theater
}

func (c *TheaterClient) mutate(ctx context.Context, m *TheaterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TheaterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TheaterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TheaterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TheaterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Theater mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id uuid.UUID) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id uuid.UUID) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id uuid.UUID) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id uuid.UUID) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Room, Session, Theater, User []ent.Hook
	}
	inters struct {
		Room, Session, Theater, User []ent.Interceptor
	}
)
