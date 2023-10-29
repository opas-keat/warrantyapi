package model

type PromotionRequest struct {
	ID                   string `json:"id"`
	PromotionType        string `json:"promotion_type"`
	PromotionBrand       string `json:"promotion_brand"`
	PromotionDetail      string `json:"promotion_detail"`
	PromotionWarrantyDay int    `json:"promotion_warranty_day"`
	PromotionStatus      string `json:"promotion_status"`
	PromotionFrom        string `json:"promotion_from"`
	PromotionTo          string `json:"promotion_to"`
}

type PromotionResponse struct {
	ID                   string `json:"id"`
	PromotionType        string `json:"promotion_type"`
	PromotionBrand       string `json:"promotion_brand"`
	PromotionDetail      string `json:"promotion_detail"`
	PromotionWarrantyDay int    `json:"promotion_warranty_day"`
	PromotionStatus      string `json:"promotion_status"`
	PromotionFrom        string `json:"promotion_from"`
	PromotionTo          string `json:"promotion_to"`
}
