package model

type ConfigRequest struct {
	ID          string `json:"id"`
	ConfigCode  string `json:"config_code"`
	ConfigValue string `json:"config_value"`
}

type ConfigResponse struct {
	ID          string `json:"id"`
	ConfigCode  string `json:"config_code"`
	ConfigValue string `json:"config_value"`
}
