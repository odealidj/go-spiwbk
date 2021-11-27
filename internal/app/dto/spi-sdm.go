package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type SpiSdmResponse struct {
	abstraction.ID
	model.SpiSdmEntity
	FilePath string `json:"file_path"`
}

type SpiSdmResponses struct {
	abstraction.ID
	model.SpiSdmEntity
	ThnAngYear string `json:"thn_ang_year"`
	SatkerName string `json:"satker_name"`
}

//Save
type SpiSdmSaveRequest struct {
	model.SpiSdmEntity
	FilePath string `json:"file_path" form:"file_path"`
}

//Update
type SpiSdmUpdateRequest struct {
	abstraction.ID
	model.SpiSdmEntity
}

//Delete
type SpiSdmDeleteRequest struct {
	abstraction.ID
}

//Get
type SpiSdmGetRequest struct {
	abstraction.Pagination
	model.SpiSdmFilter
}

//GetByID
type SpiSdmGetByIDRequest struct {
	abstraction.ID
}

type SpiSdmGetResponse struct {
	Datas          *[]SpiSdmResponses
	PaginationInfo *abstraction.PaginationInfo
}
