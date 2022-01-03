package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSubProgramRankerBulanUpsertRequest struct {
	abstraction.ID
	wbk.WbkSubProgramRankerBulanEntity
}

//Get
type WbkSubProgramRankerBulanGetRequest struct {
	abstraction.PaginationArr
	wbk.WbkSubProgramRankerBulanFilter
}

type WbkSubProgramRankerBulanGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkSubProgramRankerBulanEntity
	BulanName string `json:"bulan_name,omitempty"`
}

type WbkSubProgramRankerBulanGetInfoResponse struct {
	Datas          *[]WbkSubProgramRankerBulanGetResponse
	PaginationInfo *abstraction.PaginationInfoArr
}

type WbkSubProgramRankerBulanResponse struct {
	abstraction.ID
	wbk.WbkSubProgramRankerBulanEntity
	BulanName string `json:"bulan_name,omitempty"`
}
