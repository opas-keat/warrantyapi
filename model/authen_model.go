package model

type Auth struct {
	UserName string `json:"user_name"`
	UserPass string `json:"user_pass"`
}

type AuthResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
}
