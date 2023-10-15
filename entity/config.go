package entity

import (
	"github.com/google/uuid"
)

type Config struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	ConfigCode  string    `gorm:"column:config_code;size:50"`
	ConfigValue string    `gorm:"column:config_value;size:255"`
}

func (Config) TableName() string {
	return "wt_config"
}
