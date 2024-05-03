// Code generated by ent, DO NOT EDIT.

package ent

import (
	"PopcornMovie/ent/comment"
	"PopcornMovie/ent/movie"
	"PopcornMovie/ent/predicate"
	"PopcornMovie/ent/showtime"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// MovieUpdate is the builder for updating Movie entities.
type MovieUpdate struct {
	config
	hooks    []Hook
	mutation *MovieMutation
}

// Where appends a list predicates to the MovieUpdate builder.
func (mu *MovieUpdate) Where(ps ...predicate.Movie) *MovieUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetTitle sets the "title" field.
func (mu *MovieUpdate) SetTitle(s string) *MovieUpdate {
	mu.mutation.SetTitle(s)
	return mu
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableTitle(s *string) *MovieUpdate {
	if s != nil {
		mu.SetTitle(*s)
	}
	return mu
}

// SetGenre sets the "genre" field.
func (mu *MovieUpdate) SetGenre(s string) *MovieUpdate {
	mu.mutation.SetGenre(s)
	return mu
}

// SetNillableGenre sets the "genre" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableGenre(s *string) *MovieUpdate {
	if s != nil {
		mu.SetGenre(*s)
	}
	return mu
}

// SetStatus sets the "status" field.
func (mu *MovieUpdate) SetStatus(m movie.Status) *MovieUpdate {
	mu.mutation.SetStatus(m)
	return mu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableStatus(m *movie.Status) *MovieUpdate {
	if m != nil {
		mu.SetStatus(*m)
	}
	return mu
}

// SetLanguage sets the "language" field.
func (mu *MovieUpdate) SetLanguage(s string) *MovieUpdate {
	mu.mutation.SetLanguage(s)
	return mu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableLanguage(s *string) *MovieUpdate {
	if s != nil {
		mu.SetLanguage(*s)
	}
	return mu
}

// SetDirector sets the "director" field.
func (mu *MovieUpdate) SetDirector(s string) *MovieUpdate {
	mu.mutation.SetDirector(s)
	return mu
}

// SetNillableDirector sets the "director" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableDirector(s *string) *MovieUpdate {
	if s != nil {
		mu.SetDirector(*s)
	}
	return mu
}

// SetCast sets the "cast" field.
func (mu *MovieUpdate) SetCast(s string) *MovieUpdate {
	mu.mutation.SetCast(s)
	return mu
}

// SetNillableCast sets the "cast" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableCast(s *string) *MovieUpdate {
	if s != nil {
		mu.SetCast(*s)
	}
	return mu
}

// SetPoster sets the "poster" field.
func (mu *MovieUpdate) SetPoster(s string) *MovieUpdate {
	mu.mutation.SetPoster(s)
	return mu
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (mu *MovieUpdate) SetNillablePoster(s *string) *MovieUpdate {
	if s != nil {
		mu.SetPoster(*s)
	}
	return mu
}

// SetRated sets the "rated" field.
func (mu *MovieUpdate) SetRated(s string) *MovieUpdate {
	mu.mutation.SetRated(s)
	return mu
}

// SetNillableRated sets the "rated" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableRated(s *string) *MovieUpdate {
	if s != nil {
		mu.SetRated(*s)
	}
	return mu
}

// SetDuration sets the "duration" field.
func (mu *MovieUpdate) SetDuration(i int) *MovieUpdate {
	mu.mutation.ResetDuration()
	mu.mutation.SetDuration(i)
	return mu
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableDuration(i *int) *MovieUpdate {
	if i != nil {
		mu.SetDuration(*i)
	}
	return mu
}

// AddDuration adds i to the "duration" field.
func (mu *MovieUpdate) AddDuration(i int) *MovieUpdate {
	mu.mutation.AddDuration(i)
	return mu
}

// SetTrailer sets the "trailer" field.
func (mu *MovieUpdate) SetTrailer(s string) *MovieUpdate {
	mu.mutation.SetTrailer(s)
	return mu
}

// SetNillableTrailer sets the "trailer" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableTrailer(s *string) *MovieUpdate {
	if s != nil {
		mu.SetTrailer(*s)
	}
	return mu
}

// SetOpeningDay sets the "opening_day" field.
func (mu *MovieUpdate) SetOpeningDay(t time.Time) *MovieUpdate {
	mu.mutation.SetOpeningDay(t)
	return mu
}

// SetNillableOpeningDay sets the "opening_day" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableOpeningDay(t *time.Time) *MovieUpdate {
	if t != nil {
		mu.SetOpeningDay(*t)
	}
	return mu
}

// SetStory sets the "story" field.
func (mu *MovieUpdate) SetStory(s string) *MovieUpdate {
	mu.mutation.SetStory(s)
	return mu
}

// SetNillableStory sets the "story" field if the given value is not nil.
func (mu *MovieUpdate) SetNillableStory(s *string) *MovieUpdate {
	if s != nil {
		mu.SetStory(*s)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MovieUpdate) SetUpdatedAt(t time.Time) *MovieUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// AddShowTimeIDs adds the "showTimes" edge to the ShowTime entity by IDs.
func (mu *MovieUpdate) AddShowTimeIDs(ids ...uuid.UUID) *MovieUpdate {
	mu.mutation.AddShowTimeIDs(ids...)
	return mu
}

// AddShowTimes adds the "showTimes" edges to the ShowTime entity.
func (mu *MovieUpdate) AddShowTimes(s ...*ShowTime) *MovieUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.AddShowTimeIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (mu *MovieUpdate) AddCommentIDs(ids ...uuid.UUID) *MovieUpdate {
	mu.mutation.AddCommentIDs(ids...)
	return mu
}

// AddComments adds the "comments" edges to the Comment entity.
func (mu *MovieUpdate) AddComments(c ...*Comment) *MovieUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.AddCommentIDs(ids...)
}

// Mutation returns the MovieMutation object of the builder.
func (mu *MovieUpdate) Mutation() *MovieMutation {
	return mu.mutation
}

// ClearShowTimes clears all "showTimes" edges to the ShowTime entity.
func (mu *MovieUpdate) ClearShowTimes() *MovieUpdate {
	mu.mutation.ClearShowTimes()
	return mu
}

// RemoveShowTimeIDs removes the "showTimes" edge to ShowTime entities by IDs.
func (mu *MovieUpdate) RemoveShowTimeIDs(ids ...uuid.UUID) *MovieUpdate {
	mu.mutation.RemoveShowTimeIDs(ids...)
	return mu
}

// RemoveShowTimes removes "showTimes" edges to ShowTime entities.
func (mu *MovieUpdate) RemoveShowTimes(s ...*ShowTime) *MovieUpdate {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return mu.RemoveShowTimeIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (mu *MovieUpdate) ClearComments() *MovieUpdate {
	mu.mutation.ClearComments()
	return mu
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (mu *MovieUpdate) RemoveCommentIDs(ids ...uuid.UUID) *MovieUpdate {
	mu.mutation.RemoveCommentIDs(ids...)
	return mu
}

// RemoveComments removes "comments" edges to Comment entities.
func (mu *MovieUpdate) RemoveComments(c ...*Comment) *MovieUpdate {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return mu.RemoveCommentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MovieUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MovieUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MovieUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MovieUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MovieUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := movie.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MovieUpdate) check() error {
	if v, ok := mu.mutation.Status(); ok {
		if err := movie.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Movie.status": %w`, err)}
		}
	}
	return nil
}

func (mu *MovieUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(movie.Table, movie.Columns, sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Title(); ok {
		_spec.SetField(movie.FieldTitle, field.TypeString, value)
	}
	if value, ok := mu.mutation.Genre(); ok {
		_spec.SetField(movie.FieldGenre, field.TypeString, value)
	}
	if value, ok := mu.mutation.Status(); ok {
		_spec.SetField(movie.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := mu.mutation.Language(); ok {
		_spec.SetField(movie.FieldLanguage, field.TypeString, value)
	}
	if value, ok := mu.mutation.Director(); ok {
		_spec.SetField(movie.FieldDirector, field.TypeString, value)
	}
	if value, ok := mu.mutation.Cast(); ok {
		_spec.SetField(movie.FieldCast, field.TypeString, value)
	}
	if value, ok := mu.mutation.Poster(); ok {
		_spec.SetField(movie.FieldPoster, field.TypeString, value)
	}
	if value, ok := mu.mutation.Rated(); ok {
		_spec.SetField(movie.FieldRated, field.TypeString, value)
	}
	if value, ok := mu.mutation.Duration(); ok {
		_spec.SetField(movie.FieldDuration, field.TypeInt, value)
	}
	if value, ok := mu.mutation.AddedDuration(); ok {
		_spec.AddField(movie.FieldDuration, field.TypeInt, value)
	}
	if value, ok := mu.mutation.Trailer(); ok {
		_spec.SetField(movie.FieldTrailer, field.TypeString, value)
	}
	if value, ok := mu.mutation.OpeningDay(); ok {
		_spec.SetField(movie.FieldOpeningDay, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Story(); ok {
		_spec.SetField(movie.FieldStory, field.TypeString, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(movie.FieldUpdatedAt, field.TypeTime, value)
	}
	if mu.mutation.ShowTimesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedShowTimesIDs(); len(nodes) > 0 && !mu.mutation.ShowTimesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.ShowTimesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !mu.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MovieUpdateOne is the builder for updating a single Movie entity.
type MovieUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MovieMutation
}

// SetTitle sets the "title" field.
func (muo *MovieUpdateOne) SetTitle(s string) *MovieUpdateOne {
	muo.mutation.SetTitle(s)
	return muo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableTitle(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetTitle(*s)
	}
	return muo
}

// SetGenre sets the "genre" field.
func (muo *MovieUpdateOne) SetGenre(s string) *MovieUpdateOne {
	muo.mutation.SetGenre(s)
	return muo
}

// SetNillableGenre sets the "genre" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableGenre(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetGenre(*s)
	}
	return muo
}

// SetStatus sets the "status" field.
func (muo *MovieUpdateOne) SetStatus(m movie.Status) *MovieUpdateOne {
	muo.mutation.SetStatus(m)
	return muo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableStatus(m *movie.Status) *MovieUpdateOne {
	if m != nil {
		muo.SetStatus(*m)
	}
	return muo
}

// SetLanguage sets the "language" field.
func (muo *MovieUpdateOne) SetLanguage(s string) *MovieUpdateOne {
	muo.mutation.SetLanguage(s)
	return muo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableLanguage(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetLanguage(*s)
	}
	return muo
}

// SetDirector sets the "director" field.
func (muo *MovieUpdateOne) SetDirector(s string) *MovieUpdateOne {
	muo.mutation.SetDirector(s)
	return muo
}

// SetNillableDirector sets the "director" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableDirector(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetDirector(*s)
	}
	return muo
}

// SetCast sets the "cast" field.
func (muo *MovieUpdateOne) SetCast(s string) *MovieUpdateOne {
	muo.mutation.SetCast(s)
	return muo
}

// SetNillableCast sets the "cast" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableCast(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetCast(*s)
	}
	return muo
}

// SetPoster sets the "poster" field.
func (muo *MovieUpdateOne) SetPoster(s string) *MovieUpdateOne {
	muo.mutation.SetPoster(s)
	return muo
}

// SetNillablePoster sets the "poster" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillablePoster(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetPoster(*s)
	}
	return muo
}

// SetRated sets the "rated" field.
func (muo *MovieUpdateOne) SetRated(s string) *MovieUpdateOne {
	muo.mutation.SetRated(s)
	return muo
}

// SetNillableRated sets the "rated" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableRated(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetRated(*s)
	}
	return muo
}

// SetDuration sets the "duration" field.
func (muo *MovieUpdateOne) SetDuration(i int) *MovieUpdateOne {
	muo.mutation.ResetDuration()
	muo.mutation.SetDuration(i)
	return muo
}

// SetNillableDuration sets the "duration" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableDuration(i *int) *MovieUpdateOne {
	if i != nil {
		muo.SetDuration(*i)
	}
	return muo
}

// AddDuration adds i to the "duration" field.
func (muo *MovieUpdateOne) AddDuration(i int) *MovieUpdateOne {
	muo.mutation.AddDuration(i)
	return muo
}

// SetTrailer sets the "trailer" field.
func (muo *MovieUpdateOne) SetTrailer(s string) *MovieUpdateOne {
	muo.mutation.SetTrailer(s)
	return muo
}

// SetNillableTrailer sets the "trailer" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableTrailer(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetTrailer(*s)
	}
	return muo
}

// SetOpeningDay sets the "opening_day" field.
func (muo *MovieUpdateOne) SetOpeningDay(t time.Time) *MovieUpdateOne {
	muo.mutation.SetOpeningDay(t)
	return muo
}

// SetNillableOpeningDay sets the "opening_day" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableOpeningDay(t *time.Time) *MovieUpdateOne {
	if t != nil {
		muo.SetOpeningDay(*t)
	}
	return muo
}

// SetStory sets the "story" field.
func (muo *MovieUpdateOne) SetStory(s string) *MovieUpdateOne {
	muo.mutation.SetStory(s)
	return muo
}

// SetNillableStory sets the "story" field if the given value is not nil.
func (muo *MovieUpdateOne) SetNillableStory(s *string) *MovieUpdateOne {
	if s != nil {
		muo.SetStory(*s)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MovieUpdateOne) SetUpdatedAt(t time.Time) *MovieUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// AddShowTimeIDs adds the "showTimes" edge to the ShowTime entity by IDs.
func (muo *MovieUpdateOne) AddShowTimeIDs(ids ...uuid.UUID) *MovieUpdateOne {
	muo.mutation.AddShowTimeIDs(ids...)
	return muo
}

// AddShowTimes adds the "showTimes" edges to the ShowTime entity.
func (muo *MovieUpdateOne) AddShowTimes(s ...*ShowTime) *MovieUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.AddShowTimeIDs(ids...)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (muo *MovieUpdateOne) AddCommentIDs(ids ...uuid.UUID) *MovieUpdateOne {
	muo.mutation.AddCommentIDs(ids...)
	return muo
}

// AddComments adds the "comments" edges to the Comment entity.
func (muo *MovieUpdateOne) AddComments(c ...*Comment) *MovieUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.AddCommentIDs(ids...)
}

// Mutation returns the MovieMutation object of the builder.
func (muo *MovieUpdateOne) Mutation() *MovieMutation {
	return muo.mutation
}

// ClearShowTimes clears all "showTimes" edges to the ShowTime entity.
func (muo *MovieUpdateOne) ClearShowTimes() *MovieUpdateOne {
	muo.mutation.ClearShowTimes()
	return muo
}

// RemoveShowTimeIDs removes the "showTimes" edge to ShowTime entities by IDs.
func (muo *MovieUpdateOne) RemoveShowTimeIDs(ids ...uuid.UUID) *MovieUpdateOne {
	muo.mutation.RemoveShowTimeIDs(ids...)
	return muo
}

// RemoveShowTimes removes "showTimes" edges to ShowTime entities.
func (muo *MovieUpdateOne) RemoveShowTimes(s ...*ShowTime) *MovieUpdateOne {
	ids := make([]uuid.UUID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return muo.RemoveShowTimeIDs(ids...)
}

// ClearComments clears all "comments" edges to the Comment entity.
func (muo *MovieUpdateOne) ClearComments() *MovieUpdateOne {
	muo.mutation.ClearComments()
	return muo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (muo *MovieUpdateOne) RemoveCommentIDs(ids ...uuid.UUID) *MovieUpdateOne {
	muo.mutation.RemoveCommentIDs(ids...)
	return muo
}

// RemoveComments removes "comments" edges to Comment entities.
func (muo *MovieUpdateOne) RemoveComments(c ...*Comment) *MovieUpdateOne {
	ids := make([]uuid.UUID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return muo.RemoveCommentIDs(ids...)
}

// Where appends a list predicates to the MovieUpdate builder.
func (muo *MovieUpdateOne) Where(ps ...predicate.Movie) *MovieUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MovieUpdateOne) Select(field string, fields ...string) *MovieUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Movie entity.
func (muo *MovieUpdateOne) Save(ctx context.Context) (*Movie, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MovieUpdateOne) SaveX(ctx context.Context) *Movie {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MovieUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MovieUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MovieUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := movie.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MovieUpdateOne) check() error {
	if v, ok := muo.mutation.Status(); ok {
		if err := movie.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Movie.status": %w`, err)}
		}
	}
	return nil
}

func (muo *MovieUpdateOne) sqlSave(ctx context.Context) (_node *Movie, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(movie.Table, movie.Columns, sqlgraph.NewFieldSpec(movie.FieldID, field.TypeUUID))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Movie.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, movie.FieldID)
		for _, f := range fields {
			if !movie.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != movie.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Title(); ok {
		_spec.SetField(movie.FieldTitle, field.TypeString, value)
	}
	if value, ok := muo.mutation.Genre(); ok {
		_spec.SetField(movie.FieldGenre, field.TypeString, value)
	}
	if value, ok := muo.mutation.Status(); ok {
		_spec.SetField(movie.FieldStatus, field.TypeEnum, value)
	}
	if value, ok := muo.mutation.Language(); ok {
		_spec.SetField(movie.FieldLanguage, field.TypeString, value)
	}
	if value, ok := muo.mutation.Director(); ok {
		_spec.SetField(movie.FieldDirector, field.TypeString, value)
	}
	if value, ok := muo.mutation.Cast(); ok {
		_spec.SetField(movie.FieldCast, field.TypeString, value)
	}
	if value, ok := muo.mutation.Poster(); ok {
		_spec.SetField(movie.FieldPoster, field.TypeString, value)
	}
	if value, ok := muo.mutation.Rated(); ok {
		_spec.SetField(movie.FieldRated, field.TypeString, value)
	}
	if value, ok := muo.mutation.Duration(); ok {
		_spec.SetField(movie.FieldDuration, field.TypeInt, value)
	}
	if value, ok := muo.mutation.AddedDuration(); ok {
		_spec.AddField(movie.FieldDuration, field.TypeInt, value)
	}
	if value, ok := muo.mutation.Trailer(); ok {
		_spec.SetField(movie.FieldTrailer, field.TypeString, value)
	}
	if value, ok := muo.mutation.OpeningDay(); ok {
		_spec.SetField(movie.FieldOpeningDay, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Story(); ok {
		_spec.SetField(movie.FieldStory, field.TypeString, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(movie.FieldUpdatedAt, field.TypeTime, value)
	}
	if muo.mutation.ShowTimesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedShowTimesIDs(); len(nodes) > 0 && !muo.mutation.ShowTimesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.ShowTimesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.ShowTimesTable,
			Columns: []string{movie.ShowTimesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(showtime.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !muo.mutation.CommentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   movie.CommentsTable,
			Columns: []string{movie.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Movie{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{movie.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
