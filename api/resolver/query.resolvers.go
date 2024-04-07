package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	graphql1 "PopcornMovie/graphql"
	"PopcornMovie/model"
	"context"
)

// Theaters is the resolver for the Theaters field.
func (r *queryResolver) Theaters(ctx context.Context, input model.ListTheatersInput) (*model.ListTheatersOutput, error) {
	listTheater, count, err := r.service.Theater().ListTheaters(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListTheatersOutput{
		Data: listTheater,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
