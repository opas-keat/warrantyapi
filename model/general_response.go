package model

type GeneralResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginateResponse struct {
	Data      any `json:"data"`
	Page      int `json:"page"`
	Limit     int `json:"limit"`
	TotalPage int `json:"total_page"`
	TotalItem int `json:"total_item"`
}
