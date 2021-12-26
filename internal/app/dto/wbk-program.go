package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Get
type WbkProgramGetRequest struct {
	abstraction.Pagination
	model.WbkProgramFilter
}

//Get
type WbkProgramGetResponse struct {
	Row int `json:"row"`
	model.WbkProgramEntity
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
	model.WbkProgramEntity
	Nilai float64 `json:"nilai"`
}

//Get Info
type WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse struct {
	Datas          *[]WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse
	PaginationInfo *abstraction.PaginationInfo
}
