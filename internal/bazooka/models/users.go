package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Uid       uint64 `gorm:"type:bigint;unique;not null" json:"uid"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name"`
	LastName  string `gorm:"type:varchar(255)" json:"last_name"`
	UserName  string `gorm:"type:varchar(255)" json:"username"`
	Lang      string `gorm:"type:varchar(32)" json:"lang"`
}

func (u User) TableName() string {
	return "bazooka_user"
}

func (u User) IsValid() bool {
	if 0 >= u.Uid {
		return false
	}

	return true
}
