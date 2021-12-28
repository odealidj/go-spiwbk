package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//upsert
type WbkSubProgramRankerUpsertRequest struct {
	abstraction.ID
	model.WbkSubProgramRankerEntity
	SatkerID *int `json:"satkerID,omitempty"`
	ThnAngID *int `json:"thnAngID,omitempty"`
}

//Get
type WbkSubProgramRankerGetRequest struct {
	abstraction.Pagination
	model.WbkSubProgramRankerFilter
}

type WbkSubProgramRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.WbkSubProgramRankerEntity
}

type WbkSubProgramRankerGetInfoResponse struct {
	Datas          *[]WbkSubProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkSubProgramRankerResponse struct {
	abstraction.ID
	model.WbkSubProgramRankerEntity
}
