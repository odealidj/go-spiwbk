package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSubProgramUraianBulanUpsertRequest struct {
	abstraction.ID
	wbk.WbkSubProgramUraianBulanEntity
}

//Get
type WbkSubProgramUraianBulanGetRequest struct {
	abstraction.PaginationArr
	wbk.WbkSubProgramUraianBulanFilter
}

type WbkSubProgramUraianBulanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkSubProgramUraianBulanEntity
	BulanName string `json:"bulan_name,omitempty"`
}

type WbkSubProgramUraianBulanGetInfoResponse struct {
	Datas          *[]WbkSubProgramUraianBulanGetResponse
	PaginationInfo *abstraction.PaginationInfoArr
}

type WbkSubProgramUraianBulanResponse struct {
	abstraction.ID
	wbk.WbkSubProgramUraianBulanEntity
	BulanName string `json:"bulan_name,omitempty"`
}
