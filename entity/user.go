package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model `json:"-"`
	ID         uint   `json:"id"`
	CreatedBy  string `json:"-"`
	UserName   string `gorm:"size:255; not null"`
	Passwords  string `gorm:"size:255;not null"`
	UserType   string `gorm:"size:1;not null"`
	UserStatus string `gorm:"size:1;not null"`
	UserCode   string `gorm:"size:100;not null"`
	FirstName  string `gorm:"size:100;"`
	LastName   string `gorm:"size:100;"`
}

func (User) TableName() string {
	return "wt_user"
}
