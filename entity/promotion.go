package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Promotion struct {
	gorm.Model           `json:"-"`
	ID                   uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	CreatedBy            string    `gorm:"column:created_by; not null"`
	PromotionType        string    `gorm:"column:promotion_type;size:10"`
	PromotionBrand       string    `gorm:"column:promotion_brand;size:255"`
	PromotionDetail      string    `gorm:"column:promotion_detail;size:255"`
	PromotionWarrantyDay int       `gorm:"column:promotion_warranty_day;"`
	PromotionStatus      string    `gorm:"column:promotion_status;size:10"`
	PromotionFrom        time.Time `gorm:"column:promotion_from;"`
	PromotionTo          time.Time `gorm:"column:promotion_to;"`
}

func (Promotion) TableName() string {
	return "wt_promotion"
}
