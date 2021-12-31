package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type FrekuensiRankerUpsertRequest struct {
	abstraction.ID
	wbk.FrekuensiRankerEntity
}

//Get
type FrekuensiRankerGetRequest struct {
	abstraction.Pagination
	wbk.FrekuensiRankerFilter
}

type FrekuensiRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.FrekuensiRankerEntity
}

type FrekuensiRankerGetInfoResponse struct {
	Datas          *[]FrekuensiRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type FrekuensiRankerResponse struct {
	abstraction.ID
	wbk.FrekuensiRankerEntity
}
