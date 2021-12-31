package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//Upsert
type WbkPertanyaanUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	wbk.WbkPertanyaanEntity
}

//Upsert
type WbkPertanyaanResponse struct {
	ID int `json:"id" param:"id"`
	wbk.WbkPertanyaanEntity
}

//Get
type WbkPertanyaanGetRequest struct {
	abstraction.Pagination
	wbk.WbkPertanyaanFilter
}

//Get
type WbkPertanyaanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkPertanyaanEntity
	//WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkPertanyaanGetInfoResponse struct {
	Datas          *[]WbkPertanyaanGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
