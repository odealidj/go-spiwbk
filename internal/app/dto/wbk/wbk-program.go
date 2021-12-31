package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//Upsert
type WbkProgramUpsertRequest struct {
	ID *int `json:"id" param:"id"`
	wbk.WbkProgramEntity
}

type WbkProgramResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkProgramEntity
}

//Get
type WbkProgramGetRequest struct {
	abstraction.Pagination
	wbk.WbkProgramFilter
}

//Get
type WbkProgramGetResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkProgramEntity
	WbkKomponenName string `json:"wbk_komponen_name"`
}

//Get Info
type WbkProgramGetInfoResponse struct {
	Datas          *[]WbkProgramGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

//Get
type WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse struct {
	Row int `json:"row"`
	abstraction.ID
	wbk.WbkProgramEntity
	Nilai float64 `json:"nilai"`
}

//Get Info
type WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse struct {
	Datas          *[]WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse
	PaginationInfo *abstraction.PaginationInfo
}
