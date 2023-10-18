package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Config struct {
	gorm.Model   `json:"-"`
	ID           uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy    string    `gorm:"column:created_by; not null"`
	ConfigCode   string    `gorm:"column:config_code;size:50"`
	ConfigDetail string    `gorm:"column:config_detail;size:255"`
	ConfigValue  string    `gorm:"column:config_value;size:255"`
}

func (Config) TableName() string {
	return "wt_config"
}
