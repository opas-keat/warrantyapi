package entity

import (
	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model           `json:"-"`
	ID                   uint   `gorm:"column:id; not null"`
	CreatedBy            string `gorm:"column:created_by; not null"`
	CustomerName         string `gorm:"column:customer_name;size:255; not null"`
	CustomerEmail        string `gorm:"column:customer_email;size:100; not null"`
	CustomerPhone        string `gorm:"column:customer_phone;size:20; not null"`
	CustomerLicensePlate string `gorm:"column:customer_license_plate;size:20; not null"`
	CustomerMile         string `gorm:"column:customer_mile;size:20; not null"`
	InstallationDate     string `gorm:"column:installation_date; not null"`
	WarrantyNo           string `gorm:"column:warranty_no; not null"`
}

func (Customer) TableName() string {
	return "wt_customer"
}
