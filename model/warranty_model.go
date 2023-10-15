package model

type WarrantyRequest struct {
	ID                   string           `json:"id"`
	WarrantyNo           string           `json:"warranty_no"`
	WarrantyDateTime     string           `json:"warranty_date"`
	DealerCode           string           `json:"dealer_code"`
	DealerName           string           `json:"dealer_name"`
	CustomerName         string           `json:"customer_name"`
	CustomerPhone        string           `json:"customer_phone"`
	CustomerLicensePlate string           `json:"customer_license_plate"`
	CustomerEmail        string           `json:"customer_email"`
	CustomerMile         string           `json:"customer_mile"`
	ProductRequest       []ProductRequest `json:"products"`
}

type WarrantyResponse struct {
	ID                   string            `json:"id"`
	WarrantyNo           string            `json:"warranty_no"`
	WarrantyDateTime     string            `json:"warranty_date"`
	DealerCode           string            `json:"dealer_code"`
	DealerName           string            `json:"dealer_name"`
	CustomerName         string            `json:"customer_name"`
	CustomerPhone        string            `json:"customer_phone"`
	CustomerLicensePlate string            `json:"customer_license_plate"`
	CustomerEmail        string            `json:"customer_email"`
	CustomerMile         string            `json:"customer_mile"`
	Campagne             string            `json:"campagne"`
	ProductResponse      []ProductResponse `json:"products"`
}
