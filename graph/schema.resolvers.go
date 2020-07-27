package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sony-nurdianto/go-pedia/graph/generated"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//RegisterUser func
func (r *Resolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.RegisterUser(ctx, input)
}

//LoginUser user
func (r *Resolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.LoginUser(ctx, input)
}

//CreateProduct
func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	return r.Domain.CreateProduct(ctx, input)
}

//UpdateProduct
func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
	return r.Domain.UpdateProduct(ctx, id, input)
}

//DeleteProduct
func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteProduct(ctx, id)
}

func (r *productResolver) User(ctx context.Context, obj *model.Product) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Products(ctx context.Context, filter *model.FilterProduct, limit *int, offset *int) ([]*model.Product, error) {
	return r.Domain.ProductRepo.GetProduct(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.Domain.UserRepo.GetUserByID(id)
}

func (r *userResolver) ProductID(ctx context.Context, obj *model.User) ([]*model.Product, error) {
	return r.Domain.ProductRepo.GetUserProduct(obj)
}

func (r *userResolver) UpdataeAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type productResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *productResolver) Users(ctx context.Context, obj *model.Product) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.User)
}
