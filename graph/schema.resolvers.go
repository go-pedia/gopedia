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

func (r *bucketResolver) Users(ctx context.Context, obj *model.Bucket) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bucketResolver) Product(ctx context.Context, obj *model.Bucket) (*model.Product, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bucketResolver) User(ctx context.Context, obj *model.Bucket) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *Resolver) RegisterUser(ctx context.Context, input model.RegisterUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.RegisterUser(ctx, input)
}
func (r *Resolver) LoginUser(ctx context.Context, input model.LoginUser) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, errors.New("input errors")
	}

	return r.Domain.LoginUser(ctx, input)
}
func (r *mutationResolver) CreateBucket(ctx context.Context, input *model.NewBucket) (*model.Bucket, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateProduct(ctx context.Context, input model.NewProduct) (*model.Product, error) {
	return r.Domain.CreateProduct(ctx, input)
}

func (r *mutationResolver) UpdateProduct(ctx context.Context, id string, input model.UpdateProduct) (*model.Product, error) {
	return r.Domain.UpdateProduct(ctx, id, input)
}

func (r *mutationResolver) DeleteProduct(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteProduct(ctx, id)
}

func (r *productResolver) User(ctx context.Context, obj *model.Product) (*model.User, error) {
	return getUserLoader(ctx).Load(obj.User)
}

func (r *queryResolver) Products(ctx context.Context, filter *model.FilterProduct, limit *int, offset *int) ([]*model.Product, error) {
	return r.Domain.ProductRepo.GetProduct(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.Domain.UserRepo.GetUserByID(id)
}

func (r *queryResolver) Users(ctx context.Context, filter *model.FilterUser, limit *int, offset *int) ([]*model.User, error) {
	return r.Domain.UserRepo.GetUsers(filter, limit, offset)
}

func (r *queryResolver) Buckets(ctx context.Context) ([]*model.Bucket, error) {
	return r.Domain.BucketRepo.GetBucket()
}

func (r *userResolver) ProductID(ctx context.Context, obj *model.User) ([]*model.Product, error) {
	return r.Domain.ProductRepo.GetUserProduct(obj)
}

func (r *userResolver) UpdataeAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Bucket returns generated.BucketResolver implementation.
func (r *Resolver) Bucket() generated.BucketResolver { return &bucketResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Product returns generated.ProductResolver implementation.
func (r *Resolver) Product() generated.ProductResolver { return &productResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type bucketResolver struct{ *Resolver }
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
