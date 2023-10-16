package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model              `json:"-"`
	ID                      uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy               string    `gorm:"column:created_by; not null"`
	ProductType             string    `gorm:"column:product_type;size:20; not null"`
	ProductBrand            string    `gorm:"column:product_brand;size:255; not null"`
	ProductAmount           int       `gorm:"column:product_amount; default:0"`
	ProductStructureExpire  string    `gorm:"column:product_structure_expire; not null"`
	ProductColorExpire      string    `gorm:"column:product_color_expire; not null"`
	ProductTireExpire       string    `gorm:"column:product_tire_expire; not null"`
	ProductMileExpire       string    `gorm:"column:product_mile_expire; not null"`
	ProductPromotionExpire  string    `gorm:"column:product_promotion_expire; not null"`
	WarrantyNo              string    `gorm:"column:warranty_no;size:100; not null"`
	Promotion               string    `gorm:"column:campagne; not null"`
	PromotionDay            int       `gorm:"column:campagne_day; default 0;"`
	PromotionMile           int       `gorm:"column:campagne_mile; default 0;"`
	WarrantyWheelYear       int       `gorm:"column:warranty_wheel_year; default 0;"`
	WarrantyWheelColor      int       `gorm:"column:warranty_wheel_color; default 0;"`
	WarrantyTireYear        int       `gorm:"column:warranty_tire_year; default 0;"`
	WarrantyTireMile        int       `gorm:"column:warranty_tire_mile; default 0;"`
	WarrantyTireYearZestino int       `gorm:"column:warranty_tire_year_zestino; default 0;"`
	WarrantyTireMileZestino int       `gorm:"column:warranty_tire_mile_zestino; default 0;"`
}

func (Product) TableName() string {
	return "wt_product"
}
