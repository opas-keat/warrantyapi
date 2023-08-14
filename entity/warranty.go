package entity

import "gorm.io/gorm"

type Warranty struct {
	gorm.Model    `json:"-"`
	ID            uint   `gorm:"column:id; not null"`
	CreatedBy     string `gorm:"column:created_by; not null"`
	WarrantyNo    string `gorm:"column:warranty_no;size:100; not null"`
	ProductCode   string `gorm:"column:product_code;size:100; not null"`
	ProductDetail string `gorm:"column:product_detail;size:100; not null"`
}

func (Warranty) TableName() string {
	return "wt_warranty"
}
