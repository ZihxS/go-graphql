package mysql

import (
	"github.com/ZihxS/go-graphql/graph/model"
	"gorm.io/gorm"
)

type UsersRepo struct {
	DB *gorm.DB
}

func (r *UsersRepo) GetUserByID(id int) (*model.User, error) {
	var user *model.User

	err := r.DB.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return user, err
}
