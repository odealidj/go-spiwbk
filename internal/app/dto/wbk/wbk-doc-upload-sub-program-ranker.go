package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkDocUploadSubProgramRankerUpsertRequest struct {
	abstraction.ID
	wbk.WbkDocUploadSubProgramRankerEntity
}

//Get
type WbkDocUploadSubProgramRankerGetRequest struct {
	abstraction.Pagination
	wbk.WbkDocUploadSubProgramRankerFilter
}

type WbkDocUploadSubProgramRankerGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkDocUploadSubProgramRankerEntity
}

type WbkDocUploadSubProgramRankerGetInfoResponse struct {
	Datas          *[]WbkDocUploadSubProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkDocUploadSubProgramRankerResponse struct {
	abstraction.ID
	wbk.WbkDocUploadSubProgramRankerEntity
}
