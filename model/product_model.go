package model

type ProductRequest struct {
	ID            string `json:"id"`
	ProductType   string `json:"product_type"`
	ProductBrand  string `json:"product_brand"`
	ProductAmount int    `json:"product_amount"`
	WarrantyNo    string `json:"warranty_no"`
}

type ProductResponse struct {
	ID                      string `json:"id"`
	ProductType             string `json:"product_type"`
	ProductBrand            string `json:"product_brand"`
	ProductAmount           int    `json:"product_amount"`
	ProductStructureExpire  string `json:"product_structure_expire"`
	ProductColorExpire      string `json:"product_color_expire"`
	ProductTireExpire       string `json:"product_tire_expire"`
	ProductMileExpire       string `json:"product_mile_expire"`
	ProductPromotionExpire  string `json:"product_promotion_expire"`
	WarrantyNo              string `json:"warranty_no"`
	Promotion               string `json:"promotion"`
	PromotionDay            int    `json:"promotion_day"`
	PromotionMile           int    `json:"promotion_mile"`
	WarrantyWheelYear       int    `json:"warranty_wheel_year"`
	WarrantyWheelColor      int    `json:"warranty_wheel_color"`
	WarrantyTireYear        int    `json:"warranty_tire_year"`
	WarrantyTireMile        int    `json:"warranty_tire_mile"`
	WarrantyTireYearZestino int    `json:"warranty_tire_year_zestino"`
	WarrantyTireMileZestino int    `json:"warranty_tire_mile_zestino"`
}
