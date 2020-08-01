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

//GetUsers like the name
func (u *UserRepo) GetUsers(filter *model.FilterUser, limit, offset *int) ([]*model.User, error) {

	var users []*model.User

	query := u.DB.Model(&users)

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ? ", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
	if err != nil {
		return nil, err
	}

	return users, nil
}
