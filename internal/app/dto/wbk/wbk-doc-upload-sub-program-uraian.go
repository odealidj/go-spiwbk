package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkDocUploadSubProgramUraianUpsertRequest struct {
	abstraction.ID
	wbk.WbkDocUploadSubProgramUraianEntity
}

//Get
type WbkDocUploadSubProgramUraianGetRequest struct {
	abstraction.Pagination
	wbk.WbkDocUploadSubProgramUraianFilter
}

type WbkDocUploadSubProgramUraianGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkDocUploadSubProgramUraianEntity
}

type WbkDocUploadSubProgramUraianGetInfoResponse struct {
	Datas          *[]WbkDocUploadSubProgramUraianGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkDocUploadSubProgramUraianResponse struct {
	abstraction.ID
	wbk.WbkDocUploadSubProgramUraianEntity
}
