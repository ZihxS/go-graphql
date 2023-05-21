package mysql

import (
	"github.com/ZihxS/go-graphql/graph/model"
	"gorm.io/gorm"
)

type MeetupsRepo struct {
	DB *gorm.DB
}

func (r *MeetupsRepo) GetMeetups() ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	err := r.DB.Find(&meetups).Error

	if err != nil {
		return nil, err
	}

	return meetups, err
}

func (r *MeetupsRepo) GetMeetupsByUserID(id int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup

	err := r.DB.Where("user_id = ?", id).Find(&meetups).Error

	if err != nil {
		return nil, err
	}

	return meetups, err
}
