package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Get
type WbkKomponenGetRequest struct {
	abstraction.Pagination
	model.WbkKomponenFilter
}

//Get
type WbkKomponenGetResponse struct {
	Row int `json:"row"`
	model.WbkKomponenEntity
}

//Get Info
type WbkKomponenGetInfoResponse struct {
	Datas          *[]WbkKomponenGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
