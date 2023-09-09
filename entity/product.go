package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model    `json:"-"`
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy     string    `gorm:"column:created_by; not null"`
	ProductType   string    `gorm:"column:product_type;size:20; not null"`
	ProductBrand  string    `gorm:"column:product_brand;size:255; not null"`
	ProductAmount int       `gorm:"column:product_amount; default:0"`
	WarrantyNo    string    `gorm:"column:warranty_no;size:100; not null"`
}

func (Product) TableName() string {
	return "wt_product"
}
