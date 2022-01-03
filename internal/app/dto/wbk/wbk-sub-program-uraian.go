package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSubProgramUraianUpsertRequest struct {
	abstraction.ID
	wbk.WbkSubProgramUraianEntity
	BulanID string `json:"bulan_id" validate:"required"`
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
	B1              bool   `json:"b1"`
	B2              bool   `json:"b2"`
	B3              bool   `json:"b3"`
	B4              bool   `json:"b4"`
	B5              bool   `json:"b5"`
	B6              bool   `json:"b6"`
	B7              bool   `json:"b7"`
	B8              bool   `json:"b8"`
	B9              bool   `json:"b9"`
	B10             bool   `json:"b10"`
	B11             bool   `json:"b11"`
	B12             bool   `json:"b12"`
	FrekuensiWaktu  string `json:"frekuensi_waktu"`
	Komponen        string `json:"komponen,omitempty"`
	Program         string `json:"program,omitempty"`
	ProgramRenja    string `json:"program_renja,omitempty"`
	SubProgramRenja string `json:"sub_program_renja,omitempty"`
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
