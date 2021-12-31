package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/model/wbk"
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
	wbk.WbkProgramRankerEntity
	SatkerID *int `json:"satkerID,omitempty"`
	ThnAngID *int `json:"thnAngID,omitempty"`
}

//Get
type WbkProgramRankerGetRequest struct {
	abstraction.Pagination
	wbk.WbkProgramRankerFilter
}

type WbkProgramRankerGetResponse struct {
	Row            int    `json:"row"`
	WbkProgramID   int    `json:"wbk_program_id"`
	WbkProgramCode string `json:"wbk_program_code"`
	WbkProgramName string `json:"wbk_program_name"`
	abstraction.ID
	wbk.WbkProgramRankerEntity
	Nilai *float64 `json:"nilai,omitempty"`
}

type WbkProgramRankerGetInfoResponse struct {
	Datas          *[]WbkProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkProgramRankerResponse struct {
	abstraction.ID
	wbk.WbkProgramRankerEntity
	SatkerID *int     `json:"satkerID,omitempty"`
	ThnAngID *int     `json:"thnAngID,omitempty"`
	Nilai    *float64 `json:"nilai,omitempty"`
}
