package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type SpiAngResponse struct {
	abstraction.ID
	model.SpiAngEntity
	Year       *string `json:"year,omitempty"`
	SatkerName *string `json:"satker_name,omitempty"`
}

type SpiAngResponses struct {
	abstraction.ID
	model.SpiAngEntity
	ThnAngYear string `json:"thn_ang_year"`
	SatkerName string `json:"satker_name"`
}

//Save
type SpiAngSaveRequest struct {
	model.SpiAngEntity
	//FilePath string `json:"file_path" form:"file_path"`
}

//Update
type SpiAngUpdateRequest struct {
	abstraction.ID
	model.SpiAngEntity
}

//Delete
type SpiAngDeleteRequest struct {
	abstraction.ID
}

//Get
type SpiAngGetRequest struct {
	abstraction.Pagination
	model.SpiAngFilter
}

//GetByID
type SpiAngGetByIDRequest struct {
	abstraction.ID
}

type SpiAngGetResponse struct {
	Datas          *[]SpiAngResponses
	PaginationInfo *abstraction.PaginationInfo
}

type SpiAngGetResponses struct {
	Datas          *[]SpiAngResponse
	PaginationInfo *abstraction.PaginationInfo
}
