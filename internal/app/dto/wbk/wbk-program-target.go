package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//Upsert
type WbkProgramTargetUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	wbk.WbkProgramTargetEntity
}

//Upsert
type WbkProgramTargetResponse struct {
	ID int `json:"id" param:"id"`
	wbk.WbkProgramTargetEntity
}

//Get
type WbkProgramTargetGetRequest struct {
	abstraction.Pagination
	wbk.WbkProgramTargetFilter
}

//Get
type WbkProgramTargetGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkProgramTargetEntity
	//WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkProgramTargetGetInfoResponse struct {
	Datas          *[]WbkProgramTargetGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
