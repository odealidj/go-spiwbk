package dto

import "code-boiler/internal/abstractions"

//TODO: use this pattern {Module}{Method}{Category(Request/Response)}

type User struct {
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Status   string `json:"status"`
	IsActive bool   `json:"is_active"`
}

//region Login
type AuthLoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type AuthLoginResponse struct {
	Token string `json:"token"`
	User  User   `json:"user"`
}

//endregion

//region Register
type AuthRegisterRequest struct {
	Name     string `json:"name" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
	Status   string `json:"status" validate:"required"`
	IsActive bool   `json:"is_active"  validate:"required"`
}

type AuthRegisterResponse struct {
	abstractions.Model
	User
}

//endregion
