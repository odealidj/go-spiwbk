package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisKesesuaianResponse struct {
	abstraction.ID
	model.JenisKesesuaianEntity
}

//Get
type JenisKesesuaianGetRequest struct {
	abstraction.Pagination
	model.JenisKesesuaianFilter
}

type JenisKesesuaianGetResponses struct {
	Datas          *[]JenisKesesuaianResponse
	PaginationInfo *abstraction.PaginationInfo
}
