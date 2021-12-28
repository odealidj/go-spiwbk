package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Upsert
type WbkProgramTargetUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	model.WbkProgramTargetEntity
}

//Upsert
type WbkProgramTargetResponse struct {
	ID int `json:"id" param:"id"`
	model.WbkProgramTargetEntity
}

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
