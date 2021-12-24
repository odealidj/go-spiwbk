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
	Row int `json:"row"`
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
