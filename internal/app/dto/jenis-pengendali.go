package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisPengendaliResponse struct {
	abstraction.ID
	model.JenisPengendaliEntity
}

//Get
type JenisPengendaliGetRequest struct {
	abstraction.Pagination
	model.JenisPengendaliFilter
}

type JenisPengendaliGetResponses struct {
	Datas          *[]JenisPengendaliResponse
	PaginationInfo *abstraction.PaginationInfo
}
