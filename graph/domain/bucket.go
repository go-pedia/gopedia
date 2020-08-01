package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/sony-nurdianto/go-pedia/graph/middleware1"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//CreateBucket is function to create Bucket
func (d *Domain) CreateBucket(ctx context.Context, input model.NewBucket) (*model.Bucket, error) {
	currentUser, err := middleware1.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, errors.New("unautentichated")
	}

	bucket := &model.Bucket{
		User:    currentUser.ID,
		Product: input.Product,
	}

	return d.BucketRepo.CreateBucket(bucket)
}

//DeleteBucket this is function to delete bucket
func (d *Domain) DeleteBucket(ctx context.Context, id string) (bool, error) {
	currentUser, err := middleware1.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return false, errors.New("unautentichated")
	}

	bucket, err := d.BucketRepo.GetBucketByID(id)
	if err != nil || bucket == nil {
		return false, errors.New("bucket not exist")
	}

	if !bucket.IsBucketOwner(currentUser) {
		return false, errors.New("your not the owner")
	}

	err = d.BucketRepo.DeleteBucket(bucket)
	if err != nil {
		return false, fmt.Errorf("error while delete")
	}

	return true, nil
}
