package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//upsert
type FrekuensiRankerUpsertRequest struct {
	abstraction.ID
	model.FrekuensiRankerEntity
}

//Get
type FrekuensiRankerGetRequest struct {
	abstraction.Pagination
	model.FrekuensiRankerFilter
}

type FrekuensiRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.FrekuensiRankerEntity
}

type FrekuensiRankerGetInfoResponse struct {
	Datas          *[]FrekuensiRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type FrekuensiRankerResponse struct {
	abstraction.ID
	model.FrekuensiRankerEntity
}
