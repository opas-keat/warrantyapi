package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Warranty struct {
	gorm.Model           `json:"-"`
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy            string    `gorm:"column:created_by; not null"`
	WarrantyNo           string    `gorm:"column:warranty_no;size:100; not null"`
	WarrantyDateTime     string    `gorm:"column:warranty_date;size:20; not null"`
	DealerCode           string    `gorm:"column:dealer_code;size:50"`
	DealerName           string    `gorm:"column:dealer_name;size:255"`
	CustomerName         string    `gorm:"column:customer_name;size:255; not null"`
	CustomerPhone        string    `gorm:"column:customer_phone;size:20; not null"`
	CustomerLicensePlate string    `gorm:"column:customer_license_plate;size:20; not null"`
	CustomerEmail        string    `gorm:"column:customer_email;size:100; not null"`
	CustomerMile         string    `gorm:"column:customer_mile;size:20; not null"`
}

func (Warranty) TableName() string {
	return "wt_warranty"
}
