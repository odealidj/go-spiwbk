package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Upsert
type WbkPertanyaanUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	model.WbkPertanyaanEntity
}

//Upsert
type WbkPertanyaanResponse struct {
	ID int `json:"id" param:"id"`
	model.WbkPertanyaanEntity
}

//Get
type WbkPertanyaanGetRequest struct {
	abstraction.Pagination
	model.WbkPertanyaanFilter
}

//Get
type WbkPertanyaanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.WbkPertanyaanEntity
	//WbkProgramName string `json:"wbk_program_name"`
}

//Get Info
type WbkPertanyaanGetInfoResponse struct {
	Datas          *[]WbkPertanyaanGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
