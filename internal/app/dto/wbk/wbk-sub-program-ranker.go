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
	Komponen     string `json:"komponen,omitempty"`
	Program      string `json:"program,omitempty"`
	ProgramRenja string `json:"program_renja,omitempty"`
}

type WbkSubProgramRankerGetInfoResponse struct {
	Datas          *[]WbkSubProgramRankerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkSubProgramRankerResponse struct {
	abstraction.ID
	wbk.WbkSubProgramRankerEntity
}
