package domain

import (
	"github.com/sony-nurdianto/go-pedia/graph/model"
	"github.com/sony-nurdianto/go-pedia/graph/postgres"
)

//Domain Struct same function to store Data from DataBase
type Domain struct {
	UserRepo    postgres.UserRepo
	ProductRepo postgres.ProductRepo
	BucketRepo  postgres.BucketRepo
}

//NewDomain This function is to get code from domain.go
func NewDomain(userRepo postgres.UserRepo, bucketRepo postgres.BucketRepo, productRepo postgres.ProductRepo) *Domain {
	return &Domain{UserRepo: userRepo, BucketRepo: bucketRepo, ProductRepo: productRepo}
}

//Ownable this is to store func from model
type Ownable interface {
	IsOwner(user *model.User) bool
}

func checkOwnerShip(o Ownable, user *model.User) bool {
	return o.IsOwner(user)
}
