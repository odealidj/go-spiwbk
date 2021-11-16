package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisSdmResponse struct {
	abstraction.ID
	model.JenisSdmEntity
}

//Save
type JenisSdmSaveRequest struct {
	model.JenisSdmEntity
}

//Update
type JenisSdmUpdateRequest struct {
	abstraction.ID
	model.JenisSdmEntity
}

//Get
type JenisSdmGetRequest struct {
	abstraction.Pagination
	model.JenisSdmFilter
}

type JenisSdmGetResponse struct {
	Datas          []JenisSdmResponse
	PaginationInfo abstraction.PaginationInfo
}
