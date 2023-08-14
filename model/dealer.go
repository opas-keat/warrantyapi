package model

type DealerRequest struct {
	DealerCode    string `json:"dealer_code"`
	DealerName    string `json:"dealer_name"`
	DealerAddress string `json:"dealer_address"`
	DealerPhone   string `json:"dealer_phone"`
	DealerTax     string `json:"dealer_tax"`
}

type DealerResponse struct {
	ID            uint   `json:"id"`
	DealerCode    string `json:"dealer_code"`
	DealerName    string `json:"dealer_name"`
	DealerAddress string `json:"dealer_address"`
	DealerPhone   string `json:"dealer_phone"`
	DealerTax     string `json:"dealer_tax"`
}
