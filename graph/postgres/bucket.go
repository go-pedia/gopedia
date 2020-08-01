package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//BucketRepo this struct to colect data from database
type BucketRepo struct {
	DB *pg.DB
}

//GetBucket to get all bucket
func (ub *BucketRepo) GetBucket() ([]*model.Bucket, error) {
	var bucket []*model.Bucket
	err := ub.DB.Model(&bucket).Select()
	if err != nil {
		return nil, err
	}
	return bucket, nil
}

//GetUserBucket this function is to get user in Bucket
func (ub *BucketRepo) GetUserBucket(user *model.User) ([]*model.Bucket, error) {
	var users []*model.Bucket

	err := ub.DB.Model(&users).Where("user = ?", user.ID).Select()

	return users, err
}

//GetProductBucket Bucket is a function to product in Bucket
func (ub *BucketRepo) GetProductBucket(product *model.Product) ([]*model.Bucket, error) {
	var products []*model.Bucket

	err := ub.DB.Model(&products).Where("product = ?", product.ID).Select()

	return products, err
}

//CreateBucket this is just to simplyfy function real function is domain.bucket.go
func (ub *BucketRepo) CreateBucket(bucket *model.Bucket) (*model.Bucket, error) {
	_, err := ub.DB.Model(bucket).Returning("*").Insert()

	return bucket, err
}

//GetBucketByID Bucket
func (ub *BucketRepo) GetBucketByID(id string) (*model.Bucket, error) {
	var bucket model.Bucket

	err := ub.DB.Model(&bucket).Where("id = ?", id).First()

	return &bucket, err

}

//DeleteBucket to Delete Bucket
func (ub *BucketRepo) DeleteBucket(bucket *model.Bucket) error {
	_, err := ub.DB.Model(bucket).Where("id = ?", bucket.ID).Delete()
	return err
}
