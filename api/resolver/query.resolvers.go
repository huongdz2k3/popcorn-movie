package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.44

import (
	"PopcornMovie/ent"
	graphql1 "PopcornMovie/graphql"
	"PopcornMovie/model"
	"context"
	"fmt"
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

// Rooms is the resolver for the Rooms field.
func (r *queryResolver) Rooms(ctx context.Context, input model.ListRoomInput) (*model.ListRoomOutput, error) {
	listRooms, count, err := r.service.Room().ListRooms(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListRoomOutput{
		Data: listRooms,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// GetAvailableRooms is the resolver for the GetAvailableRooms field.
func (r *queryResolver) GetAvailableRooms(ctx context.Context, input model.ListAvailableRoomInput) (*model.AvailableRoomOutput, error) {
	check, err := r.service.Room().CheckAvailableRoom(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.AvailableRoomOutput{
		IsAvailableRoom: check,
	}, nil
}

// Foods is the resolver for the Foods field.
func (r *queryResolver) Foods(ctx context.Context, input model.ListFoodInput) (*model.ListFoodOutput, error) {
	listFood, count, err := r.service.Food().ListFood(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListFoodOutput{
		Data: listFood,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Movies is the resolver for the Movies field.
func (r *queryResolver) Movies(ctx context.Context, input model.ListMovieInput) (*model.ListMovieOutput, error) {
	listMovies, count, err := r.service.Movie().ListMovies(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListMovieOutput{
		Data: listMovies,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// GetMovieByID is the resolver for the GetMovieByID field.
func (r *queryResolver) GetMovieByID(ctx context.Context, input string) (*ent.Movie, error) {
	movie, err := r.service.Movie().GetMovieByID(ctx, input)
	if err != nil {
		return nil, err
	}

	return movie, nil
}

// Comments is the resolver for the Comments field.
func (r *queryResolver) Comments(ctx context.Context, input model.ListCommentInput) (*model.ListCommentOutput, error) {
	panic(fmt.Errorf("not implemented: Comments - Comments"))
}

// ShowTimes is the resolver for the ShowTimes field.
func (r *queryResolver) ShowTimes(ctx context.Context, input model.ListShowTimeInput) (*model.ListShowTimeOutput, error) {
	listShowTimes, count, err := r.service.ShowTime().ListShowTimes(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListShowTimeOutput{
		Data: listShowTimes,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Seats is the resolver for the Seats field.
func (r *queryResolver) Seats(ctx context.Context, input model.ListSeatInput) (*model.ListSeatOutput, error) {
	seats, count, err := r.service.Seat().GetAllSeats(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListSeatOutput{
		Data: seats,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// GetAvailableSeats is the resolver for the GetAvailableSeats field.
func (r *queryResolver) GetAvailableSeats(ctx context.Context, input model.ListAvailableSeatInput) (*model.ListAvailableSeatOutput, error) {
	availableSeats, count, err := r.service.Seat().ListAvailableSeats(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListAvailableSeatOutput{
		Data: availableSeats,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Users is the resolver for the Users field.
func (r *queryResolver) Users(ctx context.Context, input model.ListUserInput) (*model.ListUserOutput, error) {
	listUsers, count, err := r.service.User().GetAllUsers(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListUserOutput{
		Data: listUsers,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Transactions is the resolver for the Transactions field.
func (r *queryResolver) Transactions(ctx context.Context, input model.ListTransactionInput) (*model.ListTransactionOutput, error) {
	listTransactions, count, err := r.service.Transaction().GetAllTransactions(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListTransactionOutput{
		Data: listTransactions,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// GetRevenue is the resolver for the GetRevenue field.
func (r *queryResolver) GetRevenue(ctx context.Context, input model.RevenueInput) (*model.YearlyRevenueOutput, error) {
	return r.service.Transaction().GetRevenue(ctx, input)
}

// Tickets is the resolver for the Tickets field.
func (r *queryResolver) Tickets(ctx context.Context, input model.ListTicketInput) (*model.ListTicketOutput, error) {
	tickets, count, err := r.service.Ticket().GetAllTickets(ctx, input)
	if err != nil {
		return nil, err
	}

	return &model.ListTicketOutput{
		Data: tickets,
		Pagination: &model.PaginationOutput{
			Total: count,
		},
	}, nil
}

// Query returns graphql1.QueryResolver implementation.
func (r *Resolver) Query() graphql1.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
