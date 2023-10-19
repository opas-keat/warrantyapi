package model

type User struct {
	ID         uint   `json:"id"`
	CreatedBy  string `json:"-"`
	UserName   string `json:"user_name" validate:"required"`
	Passwords  string `json:"user_pass" validate:"required"`
	UserType   string `json:"user_type" validate:"required"`
	UserStatus string `json:"user_status" validate:"required"`
	Code       string `json:"code"`
	FirstName  string `json:"first_name" validate:"required"`
	LastName   string `json:"last_name" validate:"required"`
}

type UserResponse struct {
	ID         uint   `json:"id"`
	UserType   string `json:"user_type"`
	UserStatus string `json:"user_status"`
	UserName   string `json:"user_name"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
}
