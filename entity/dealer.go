package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Dealer struct {
	gorm.Model    `json:"-"`
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy     string    `gorm:"column:created_by; not null"`
	DealerCode    string    `gorm:"column:dealer_code;size:50"`
	DealerName    string    `gorm:"column:dealer_name;size:255"`
	DealerAddress string    `gorm:"column:dealer_address;size:1000"`
	DealerPhone   string    `gorm:"column:dealer_phone;size:255"`
	DealerTax     string    `gorm:"column:dealer_tax;size:20"`
	DealerArea    int       `gorm:"column:dealer_area;"`
}

func (Dealer) TableName() string {
	return "wt_dealer"
}
