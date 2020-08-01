package postgres

import (
	"github.com/go-pg/pg/v9"
)

//BucketRepo this struct to colect data from database
type BucketRepo struct {
	DB *pg.DB
}

//GetBucketUser This function indormation user Bucket
// func (ub *BucketRepo) GetBucketUser(user *model.User) ([]*model.Bucket, error) {
// 	var users []*model.Bucket

// 	err := ub.DB.Model(&users).Where("user = ? ", user.ID).Select()
// 	return users, err
// }

//GetBucketProduct This function is to Get information product user
// func (ub *BucketRepo) GetBucketProduct(product *model.Product) ([]*model.Bucket, error) {
// 	var products []*model.Bucket

// 	err := ub.DB.Model(&products).Where("product = ?", product.ID).Select()
// 	return products, err
// }
