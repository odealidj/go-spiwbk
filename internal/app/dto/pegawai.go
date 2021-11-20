package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type PegawaiResponse struct {
	abstraction.ID
	model.PegawaiEntity
}

//Save
type PegawaiSaveRequest struct {
	model.PegawaiEntity
}

//Update
type PegawaiUpdateRequest struct {
	abstraction.ID
	model.PegawaiEntity
}

//Update
type PegawaiDeleteRequest struct {
	abstraction.ID
}

//Get
type PegawaiGetRequest struct {
	abstraction.Pagination
	model.PegawaiFilter
}

type PegawaiGetResponse struct {
	Datas          *[]PegawaiResponse
	PaginationInfo *abstraction.PaginationInfo
}
