package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Save
type RkaklSaveRequest struct {
	model.RkaklEntity
	FilePath  string `json:"file_path" form:"file_path"`
	Sheetname string `json:"sheetname" form:"sheetname"`
}

//Update
type RkaklUpdateRequest struct {
	abstraction.ID
	model.RkaklEntity
	FilePath  string `json:"file_path" form:"file" validate:"required"`
	Sheetname string `json:"sheetname" form:"sheetname" validate:"required"`
}

//Delete
type RkaklDeleteRequest struct {
	abstraction.ID
}

type RkaklDeleteResponse struct {
	abstraction.ID
}

//GetByID
type RkaklGetByIDRequest struct {
	abstraction.ID
}

//Get
type RkaklGetRequest struct {
	abstraction.Pagination
	model.RkaklFilter
}

type RkaklGetInfoResponse struct {
	Datas          *[]RkaklResponse
	PaginationInfo *abstraction.PaginationInfo
}

type RkaklResponse struct {
	abstraction.ID
	model.RkaklEntity
	ThnAngYear *string `json:"thn_ang_year,omitempty"`
	SatkerName *string `json:"satker_name,omitempty"`
	Filepath   *string `json:"filepath,omitempty"`
}
