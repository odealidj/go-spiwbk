package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//Get
type WbkKomponenGetRequest struct {
	abstraction.Pagination
	wbk.WbkKomponenFilter
}

//Get
type WbkKomponenGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkKomponenEntity
}

//Get Info
type WbkKomponenGetInfoResponse struct {
	Datas          *[]WbkKomponenGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
