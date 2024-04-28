package model

type ExcelsResponse struct {
	CustomerName           string `json:"customer_name"`
	CustomerPhone          string `json:"customer_phone"`
	CustomerLicensePlate   string `json:"customer_license_plate"`
	CustomerEmail          string `json:"customer_email"`
	WarrantyNo             string `json:"warranty_no"`
	DealerName             string `json:"dealer_name"`
	WarrantyDateTime       string `json:"warranty_date"`
	ProductType            string `json:"product_type"`
	ProductBrand           string `json:"product_brand"`
	ProductAmount          int    `json:"product_amount"`
	ProductStructureExpire string `json:"product_structure_expire"`
	ProductColorExpire     string `json:"product_color_expire"`
	ProductTireExpire      string `json:"product_tire_expire"`
	ProductMileExpire      string `json:"product_mile_expire"`
	Campagne               string `json:"campagne"`
}
