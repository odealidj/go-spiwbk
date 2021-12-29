package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//upsert
type BulanRankerUpsertRequest struct {
	abstraction.ID
	model.BulanRankerEntity
}

//Get
type BulanRankerGetRequest struct {
	abstraction.Pagination
	model.BulanRankerFilter
}

type BulanRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.BulanRankerEntity
}

type BulanRankerGetInfoResponse struct {
	Datas          *[]BulanRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type BulanRankerResponse struct {
	abstraction.ID
	model.BulanRankerEntity
}
