package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type WbkProgramRankerSaveRequest struct {
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
	//model.SpiPbjRekapitulasiEntity
}

type WbkProgramRankerGetSatkerNilaiResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.SatkerEntity
	Nilai float64 `json:"nilai"`
}

type WbkProgramRankerGetSatkerNilaiInfoResponse struct {
	Datas          *[]WbkProgramRankerGetSatkerNilaiResponse
	PaginationInfo *abstraction.PaginationInfo
}

//upsert
type WbkProgramRankerUpsertRequest struct {
	abstraction.ID
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
	model.WbkProgramRankerEntity
}

//Get
type WbkProgramRankerGetRequest struct {
	abstraction.Pagination
	model.WbkProgramRankerFilter
}

type WbkProgramRankerGetResponse struct {
	Row            int    `json:"row"`
	WbkProgramID   int    `json:"wbk_program_id"`
	WbkProgramCode string `json:"wbk_program_code"`
	WbkProgramName string `json:"wbk_program_name"`
	abstraction.ID
	model.WbkProgramRankerEntity
	Nilai float64 `json:"nilai"`
}

type WbkProgramRankerGetInfoResponse struct {
	Datas          *[]WbkProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkProgramRankerResponse struct {
	abstraction.ID
	model.WbkProgramRankerEntity
	SatkerID int     `json:"satkerID"`
	ThnAngID int     `json:"thnAngID"`
	Nilai    float64 `json:"nilai"`
}
