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
func (b *BucketRepo) GetBucket() ([]*model.Bucket, error) {
	var buckets []*model.Bucket

	err := b.DB.Model(&buckets).Select()
	if err != nil {
		return nil, err
	}

	return buckets, err

}
