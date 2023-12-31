package model

type DealerRequest struct {
	ID            string `json:"id"`
	DealerCode    string `json:"dealer_code"`
	DealerName    string `json:"dealer_name"`
	DealerAddress string `json:"dealer_address"`
	DealerPhone   string `json:"dealer_phone"`
	DealerTax     string `json:"dealer_tax"`
	DealerArea    int    `json:"dealer_area"`
}

type DealerResponse struct {
	ID            string `json:"id"`
	DealerCode    string `json:"dealer_code"`
	DealerName    string `json:"dealer_name"`
	DealerAddress string `json:"dealer_address"`
	DealerPhone   string `json:"dealer_phone"`
	DealerTax     string `json:"dealer_tax"`
	DealerArea    int    `json:"dealer_area"`
}
