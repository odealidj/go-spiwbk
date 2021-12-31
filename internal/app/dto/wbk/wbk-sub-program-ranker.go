package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSubProgramRankerUpsertRequest struct {
	abstraction.ID
	wbk.WbkSubProgramRankerEntity
}

//Get
type WbkSubProgramRankerGetRequest struct {
	abstraction.Pagination
	wbk.WbkSubProgramRankerFilter
}

type WbkSubProgramRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkSubProgramRankerEntity
}

type WbkSubProgramRankerGetInfoResponse struct {
	Datas          *[]WbkSubProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkSubProgramRankerResponse struct {
	abstraction.ID
	wbk.WbkSubProgramRankerEntity
}
