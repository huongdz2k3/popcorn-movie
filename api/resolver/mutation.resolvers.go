package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	graphql1 "PopcornMovie/graphql"
	"PopcornMovie/model"
	"context"
)

// Signup is the resolver for the Signup field.
func (r *mutationResolver) Signup(ctx context.Context, input model.RegisterInput) (string, error) {
	_, err := r.service.Auth().Register(ctx, input)
	if err != nil {
		return "", err
	}

	return "Register Successfully", nil
}

// Login is the resolver for the Login field.
func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.Jwt, error) {
	jwt, err := r.service.Auth().Login(ctx, input)
	if err != nil {
		return nil, err
	}

	return jwt, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }