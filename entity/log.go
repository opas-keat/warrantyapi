package entity

import (
	"time"

	"gorm.io/gorm"
)

type Log struct {
	gorm.Model `json:"-"`
	ID         uint      `json:"id"`
	CreatedAt  time.Time `json:"createdAt"`
	CreatedBy  string    `json:"createdBy"`
	Module     string    `json:"module"`
	Detail     string    `json:"detail"`
}

func (Log) TableName() string {
	return "wt_log"
}
