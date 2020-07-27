package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/sony-nurdianto/go-pedia/graph/model"
)

//UserRepo for users
type UserRepo struct {
	DB *pg.DB
}

//GetUserByField get user with specific value
func (u *UserRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()

	return &user, err
}

//GetUserByID controleer id
func (u *UserRepo) GetUserByID(id string) (*model.User, error) {
	return u.GetUserByField("id", id)
}

//GetUserByEmail like the name
func (u *UserRepo) GetUserByEmail(email string) (*model.User, error) {
	return u.GetUserByField("email", email)
}

//GetUserByName like the name
func (u *UserRepo) GetUserByName(name string) (*model.User, error) {
	return u.GetUserByField("user_name", name)
}

//CreateUser like the name
func (u *UserRepo) CreateUser(tx *pg.Tx, user *model.User) (*model.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}
