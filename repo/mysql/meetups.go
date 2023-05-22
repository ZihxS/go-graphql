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

func (r *MeetupsRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	result := r.DB.Create(&meetup)

	if result.Error != nil {
		return nil, result.Error
	}

	return meetup, nil
}

func (r *MeetupsRepo) GetMeetupByID(id int) (meetup *model.Meetup, err error) {
	err = r.DB.Where("id = ?", id).Find(&meetup).Error

	if err != nil {
		return nil, err
	}

	return meetup, err
}

func (r *MeetupsRepo) UpdateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	err := r.DB.Model(&meetup).Updates(meetup).Error

	if err != nil {
		return nil, err
	}

	return r.GetMeetupByID(meetup.ID)
}
