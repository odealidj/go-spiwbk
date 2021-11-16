package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	res "codeid-boiler/pkg/util/response"
)

// Login
type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type LoginResponse struct {
	Token string `json:"token"`
	abstraction.ID
	model.UserAppEntity
}
type LoginResponseDoc struct {
	Body struct {
		Meta res.Meta      `json:"meta"`
		Data LoginResponse `json:"data"`
	} `json:"body"`
}

// Register
type RegisterRequest struct {
	LoginRequest
	model.UserAppEntity
}

type RegisterResponse struct {
	abstraction.ID
	model.UserAppEntity
}
type RegisterResponseDoc struct {
	Body struct {
		Meta res.Meta         `json:"meta"`
		Data RegisterResponse `json:"data"`
	} `json:"body"`
}
