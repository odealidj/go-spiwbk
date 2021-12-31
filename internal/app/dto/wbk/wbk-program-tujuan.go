package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//Upsert
type WbkProgramTujuanUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	wbk.WbkProgramTujuanEntity
}

type WbkProgramTujuanResponse struct {
	abstraction.ID
	wbk.WbkProgramTujuanEntity
}

//Get
type WbkProgramTujuanGetRequest struct {
	abstraction.Pagination
	wbk.WbkProgramTujuanFilter
}

//Get
type WbkProgramTujuanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkProgramTujuanEntity
	WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkProgramTujuanGetInfoResponse struct {
	Datas          *[]WbkProgramTujuanGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
