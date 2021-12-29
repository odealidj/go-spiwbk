package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//upsert
type WbkSubProgramUraianUpsertRequest struct {
	abstraction.ID
	model.WbkSubProgramUraianEntity
}

//Get
type WbkSubProgramUraianGetRequest struct {
	abstraction.Pagination
	model.WbkSubProgramUraianFilter
}

type WbkSubProgramUraianGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	model.WbkSubProgramUraianEntity
	FrekuensiRankerName string `json:"frekuensi_ranker_name"`
}

type WbkSubProgramUraianGetInfoResponse struct {
	Datas          *[]WbkSubProgramUraianGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkSubProgramUraianResponse struct {
	abstraction.ID
	model.WbkSubProgramUraianEntity
	FrekuensiRankerName string `json:"frekuensi_ranker_name"`
}
