package resolvers

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"strconv"

	"github.com/OswinZheng/gin-web-F/internal/dto/auth_dto"
	"github.com/OswinZheng/gin-web-F/internal/services/auth"

	"github.com/OswinZheng/gin-web-F/internal/graph/generated"
	"github.com/OswinZheng/gin-web-F/internal/graph/model"
)

type Resolver struct{}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.UserInput) (*model.User, error) {
	username, password := input.Username, input.Password
	_, userInfo := auth.Register(&auth_dto.AddAuthDto{username, password})

	return &model.User{
		ID:       strconv.Itoa(userInfo.Id),
		Username: userInfo.UserName,
	}, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) ([]*model.User, error) {
	return []*model.User{
		{"1", "test"},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
