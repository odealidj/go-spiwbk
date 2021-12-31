package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSubProgramUraianUpsertRequest struct {
	abstraction.ID
	wbk.WbkSubProgramUraianEntity
}

//Get
type WbkSubProgramUraianGetRequest struct {
	abstraction.PaginationArr
	wbk.WbkSubProgramUraianFilter
}

type WbkSubProgramUraianGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkSubProgramUraianEntity
	FrekuensiRankerName string `json:"frekuensi_ranker_name"`
}

type WbkSubProgramUraianGetInfoResponse struct {
	Datas          *[]WbkSubProgramUraianGetResponse
	PaginationInfo *abstraction.PaginationInfoArr
}

type WbkSubProgramUraianResponse struct {
	abstraction.ID
	wbk.WbkSubProgramUraianEntity
	FrekuensiRankerName string `json:"frekuensi_ranker_name"`
}
