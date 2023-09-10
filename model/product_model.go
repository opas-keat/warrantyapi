package model

type ProductRequest struct {
	ID            string `json:"id"`
	ProductType   string `json:"product_type"`
	ProductBrand  string `json:"product_brand"`
	ProductAmount int    `json:"product_amount"`
	WarrantyNo    string `json:"warranty_no"`
}

type ProductResponse struct {
	ID            string `json:"id"`
	ProductType   string `json:"product_type"`
	ProductBrand  string `json:"product_brand"`
	ProductAmount int    `json:"product_amount"`
	WarrantyNo    string `json:"warranty_no"`
}
