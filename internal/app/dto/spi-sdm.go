package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type SpiSdmResponse struct {
	abstraction.ID
	model.SpiSdmEntity
	ThnAngYear string `json:"thn_ang_year"`
	SatkerName string `json:"satker_name"`
}

//Save
type SpiSdmSaveRequest struct {
	model.SpiSdmEntity
}

//Update
type SpiSdmUpdateRequest struct {
	abstraction.ID
	model.SpiSdmEntity
}

//Get
type SpiSdmGetRequest struct {
	abstraction.Pagination
	model.SpiSdmFilter
}

type SpiSdmGetResponse struct {
	Datas          *[]SpiSdmResponse
	PaginationInfo *abstraction.PaginationInfo
}
