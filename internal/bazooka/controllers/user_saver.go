package controllers

import (
	"bazooka/internal/pkg/models"
	"github.com/jinzhu/gorm"
)

type IUserSaver interface {
	Save(user *models.User) error
	Load(user *models.User) error
}

type UserSaver struct {
	db *gorm.DB
}

func (s UserSaver) Save(user *models.User) error {
	if s.db.NewRecord(*user) {
		return s.db.Create(user).Error
	}

	return s.db.Update(user).Error
}

func (s UserSaver) Load(user *models.User) error {
	err := s.db.First(user).Error
	return err
}
