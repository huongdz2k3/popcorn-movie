// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"PopcornMovie/ent"
	"PopcornMovie/ent/comment"
	"PopcornMovie/ent/food"
	"PopcornMovie/ent/foodorderline"
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/resetpassword"
	"PopcornMovie/ent/room"
	"PopcornMovie/ent/seat"
	"PopcornMovie/ent/session"
	"PopcornMovie/ent/showtime"
	"PopcornMovie/ent/theater"
	"PopcornMovie/ent/ticket"
	"PopcornMovie/ent/transaction"
	"PopcornMovie/ent/user"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The CommentFunc type is an adapter to allow the use of ordinary function as a Querier.
type CommentFunc func(context.Context, *ent.CommentQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f CommentFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.CommentQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.CommentQuery", q)
}

// The TraverseComment type is an adapter to allow the use of ordinary function as Traverser.
type TraverseComment func(context.Context, *ent.CommentQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseComment) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseComment) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.CommentQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.CommentQuery", q)
}

// The FoodFunc type is an adapter to allow the use of ordinary function as a Querier.
type FoodFunc func(context.Context, *ent.FoodQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f FoodFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.FoodQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.FoodQuery", q)
}

// The TraverseFood type is an adapter to allow the use of ordinary function as Traverser.
type TraverseFood func(context.Context, *ent.FoodQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFood) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFood) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FoodQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.FoodQuery", q)
}

// The FoodOrderLineFunc type is an adapter to allow the use of ordinary function as a Querier.
type FoodOrderLineFunc func(context.Context, *ent.FoodOrderLineQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f FoodOrderLineFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.FoodOrderLineQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.FoodOrderLineQuery", q)
}

// The TraverseFoodOrderLine type is an adapter to allow the use of ordinary function as Traverser.
type TraverseFoodOrderLine func(context.Context, *ent.FoodOrderLineQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFoodOrderLine) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFoodOrderLine) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.FoodOrderLineQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.FoodOrderLineQuery", q)
}

// The MovieFunc type is an adapter to allow the use of ordinary function as a Querier.
type MovieFunc func(context.Context, *ent.MovieQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f MovieFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.MovieQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.MovieQuery", q)
}

// The TraverseMovie type is an adapter to allow the use of ordinary function as Traverser.
type TraverseMovie func(context.Context, *ent.MovieQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseMovie) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseMovie) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.MovieQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.MovieQuery", q)
}

// The ResetPasswordFunc type is an adapter to allow the use of ordinary function as a Querier.
type ResetPasswordFunc func(context.Context, *ent.ResetPasswordQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ResetPasswordFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ResetPasswordQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ResetPasswordQuery", q)
}

// The TraverseResetPassword type is an adapter to allow the use of ordinary function as Traverser.
type TraverseResetPassword func(context.Context, *ent.ResetPasswordQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseResetPassword) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseResetPassword) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ResetPasswordQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ResetPasswordQuery", q)
}

// The RoomFunc type is an adapter to allow the use of ordinary function as a Querier.
type RoomFunc func(context.Context, *ent.RoomQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f RoomFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.RoomQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.RoomQuery", q)
}

// The TraverseRoom type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRoom func(context.Context, *ent.RoomQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRoom) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRoom) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RoomQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.RoomQuery", q)
}

// The SeatFunc type is an adapter to allow the use of ordinary function as a Querier.
type SeatFunc func(context.Context, *ent.SeatQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SeatFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SeatQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SeatQuery", q)
}

// The TraverseSeat type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSeat func(context.Context, *ent.SeatQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSeat) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSeat) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SeatQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SeatQuery", q)
}

// The SessionFunc type is an adapter to allow the use of ordinary function as a Querier.
type SessionFunc func(context.Context, *ent.SessionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f SessionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.SessionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.SessionQuery", q)
}

// The TraverseSession type is an adapter to allow the use of ordinary function as Traverser.
type TraverseSession func(context.Context, *ent.SessionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseSession) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseSession) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.SessionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.SessionQuery", q)
}

// The ShowTimeFunc type is an adapter to allow the use of ordinary function as a Querier.
type ShowTimeFunc func(context.Context, *ent.ShowTimeQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ShowTimeFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ShowTimeQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ShowTimeQuery", q)
}

// The TraverseShowTime type is an adapter to allow the use of ordinary function as Traverser.
type TraverseShowTime func(context.Context, *ent.ShowTimeQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseShowTime) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseShowTime) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ShowTimeQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ShowTimeQuery", q)
}

// The TheaterFunc type is an adapter to allow the use of ordinary function as a Querier.
type TheaterFunc func(context.Context, *ent.TheaterQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TheaterFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TheaterQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TheaterQuery", q)
}

// The TraverseTheater type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTheater func(context.Context, *ent.TheaterQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTheater) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTheater) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TheaterQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TheaterQuery", q)
}

// The TicketFunc type is an adapter to allow the use of ordinary function as a Querier.
type TicketFunc func(context.Context, *ent.TicketQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TicketFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TicketQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TicketQuery", q)
}

// The TraverseTicket type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTicket func(context.Context, *ent.TicketQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTicket) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTicket) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TicketQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TicketQuery", q)
}

// The TransactionFunc type is an adapter to allow the use of ordinary function as a Querier.
type TransactionFunc func(context.Context, *ent.TransactionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f TransactionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.TransactionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.TransactionQuery", q)
}

// The TraverseTransaction type is an adapter to allow the use of ordinary function as Traverser.
type TraverseTransaction func(context.Context, *ent.TransactionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseTransaction) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseTransaction) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.TransactionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.TransactionQuery", q)
}

// The UserFunc type is an adapter to allow the use of ordinary function as a Querier.
type UserFunc func(context.Context, *ent.UserQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f UserFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// The TraverseUser type is an adapter to allow the use of ordinary function as Traverser.
type TraverseUser func(context.Context, *ent.UserQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseUser) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseUser) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.UserQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.UserQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.CommentQuery:
		return &query[*ent.CommentQuery, predicate.Comment, comment.OrderOption]{typ: ent.TypeComment, tq: q}, nil
	case *ent.FoodQuery:
		return &query[*ent.FoodQuery, predicate.Food, food.OrderOption]{typ: ent.TypeFood, tq: q}, nil
	case *ent.FoodOrderLineQuery:
		return &query[*ent.FoodOrderLineQuery, predicate.FoodOrderLine, foodorderline.OrderOption]{typ: ent.TypeFoodOrderLine, tq: q}, nil
	case *ent.MovieQuery:
		return &query[*ent.MovieQuery, predicate.Movie, movie.OrderOption]{typ: ent.TypeMovie, tq: q}, nil
	case *ent.ResetPasswordQuery:
		return &query[*ent.ResetPasswordQuery, predicate.ResetPassword, resetpassword.OrderOption]{typ: ent.TypeResetPassword, tq: q}, nil
	case *ent.RoomQuery:
		return &query[*ent.RoomQuery, predicate.Room, room.OrderOption]{typ: ent.TypeRoom, tq: q}, nil
	case *ent.SeatQuery:
		return &query[*ent.SeatQuery, predicate.Seat, seat.OrderOption]{typ: ent.TypeSeat, tq: q}, nil
	case *ent.SessionQuery:
		return &query[*ent.SessionQuery, predicate.Session, session.OrderOption]{typ: ent.TypeSession, tq: q}, nil
	case *ent.ShowTimeQuery:
		return &query[*ent.ShowTimeQuery, predicate.ShowTime, showtime.OrderOption]{typ: ent.TypeShowTime, tq: q}, nil
	case *ent.TheaterQuery:
		return &query[*ent.TheaterQuery, predicate.Theater, theater.OrderOption]{typ: ent.TypeTheater, tq: q}, nil
	case *ent.TicketQuery:
		return &query[*ent.TicketQuery, predicate.Ticket, ticket.OrderOption]{typ: ent.TypeTicket, tq: q}, nil
	case *ent.TransactionQuery:
		return &query[*ent.TransactionQuery, predicate.Transaction, transaction.OrderOption]{typ: ent.TypeTransaction, tq: q}, nil
	case *ent.UserQuery:
		return &query[*ent.UserQuery, predicate.User, user.OrderOption]{typ: ent.TypeUser, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
