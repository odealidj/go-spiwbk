package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Get
type WbkProgramTargetGetRequest struct {
	abstraction.Pagination
	model.WbkProgramTargetFilter
}

//Get
type WbkProgramTargetGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.WbkProgramTargetEntity
	//WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkProgramTargetGetInfoResponse struct {
	Datas          *[]WbkProgramTargetGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
