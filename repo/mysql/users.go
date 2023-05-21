package mysql

import (
	"github.com/ZihxS/go-graphql/graph/model"
	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func (r *UsersRepo) GetUsers() ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, err
}

func (r *UsersRepo) GetUserByID(id int) (*model.User, error) {
	var user *model.User

	err := r.DB.Where("id = ?", id).First(&user).Error

	if err != nil {
		return nil, err
	}

	return user, err
}
