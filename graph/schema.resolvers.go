package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/aberyotaro/gql_sample_api/graph/generated"
	"github.com/aberyotaro/gql_sample_api/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	u := &model.User{
		Name: input.Name,
	}

	res := r.ORM.Create(u)
	if res.Error != nil {
		return nil, res.Error
	}

	return u, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	res := r.ORM.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}

	return users, nil
}

func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	var user model.User
	res := r.ORM.First(&user, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
