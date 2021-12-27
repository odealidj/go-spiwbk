package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Upsert
type WbkProgramTujuanUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	model.WbkProgramTujuanEntity
}

type WbkProgramTujuanResponse struct {
	abstraction.ID
	model.WbkProgramTujuanEntity
}

//Get
type WbkProgramTujuanGetRequest struct {
	abstraction.Pagination
	model.WbkProgramTujuanFilter
}

//Get
type WbkProgramTujuanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.WbkProgramTujuanEntity
	WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkProgramTujuanGetInfoResponse struct {
	Datas          *[]WbkProgramTujuanGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
