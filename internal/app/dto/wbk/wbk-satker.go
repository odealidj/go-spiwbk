package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/wbk"
)

//upsert
type WbkSatkerUpsertRequest struct {
	abstraction.ID
	ThnAngID int `json:"thn_ang_id"`
	SatkerID int `json:"satker_id"`
}

//Get ProgramByThnAngIDAndSatkerID
type WbkSatkerGetProgramByThnAngAndSatkerAndWbkKomponenRequest struct {
	abstraction.Pagination
	ThnAngID      int  `json:"thn_ang_id" query:"thn_ang_id"`
	SatkerID      int  `json:"satker_id" query:"satker_id"`
	WbkKomponenID *int `json:"wbk_komponen_id,omitempty" query:"wbk_komponen_id"`
}

//Get ProgramByThnAngIDAndSatkerID
type WbkSatkerGetProgramResponse struct {
	Row int `json:"row"`
	abstraction.ID
	ThnAngSatkerID  int    `json:"thn_ang_satker_id"`
	WbkKomponenID   int    `json:"wbk_komponen_id"`
	WbkKomponenCode string `json:"wbk_komponen_code"`
	WbkKomponenName string `json:"wbk_komponen_name"`
	WbkKomponen     string `json:"wbk_komponen"`
	WbkProgramID    int    `json:"wbk_program_id"`
	WbkProgramCode  string `json:"wbk_program_code"`
	WbkProgramName  string `json:"wbk_program_name"`
	WbkProgram      string `json:"wbk_program"`
	ThnAngID        int    `json:"thn_ang_id"`
	Year            string `json:"year"`
	SatkerID        int    `json:"satker_id"`
	SatkerName      string `json:"satker_name"`
}

//Get Info ProgramByThnAngIDAndSatkerID
type WbkSatkerGetProgramInfoResponse struct {
	Datas *[]WbkSatkerGetProgramResponse
	WbkSatkerPogramGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

//Get
type WbkSatkerGetRequest struct {
	abstraction.Pagination
	wbk.WbkSatkerFilter
}

type WbkSatkerGetResponse struct {
	abstraction.ID
	wbk.WbkSatkerEntity
	B1             *bool   `json:"b1"`
	B2             *bool   `json:"b2"`
	B3             *bool   `json:"b3"`
	B4             *bool   `json:"b4"`
	B5             *bool   `json:"b5"`
	B6             *bool   `json:"b6"`
	B7             *bool   `json:"b7"`
	B8             *bool   `json:"b8"`
	B9             *bool   `json:"b9"`
	B10            *bool   `json:"b10"`
	B11            *bool   `json:"b11"`
	B12            *bool   `json:"b12"`
	FrekuensiWaktu *string `json:"frekuensi_waktu"`
	Komponen       string  `json:"komponen,omitempty"`
	Program        string  `json:"program,omitempty"`
	ProgramRenja   string  `json:"program_renja,omitempty"`
}

type WbkSatkerGetInfoResponse struct {
	Datas          *[]WbkSatkerGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type WbkSatkerResponse struct {
	abstraction.ID
	wbk.WbkSatkerEntity
}

type WbkSatkerResponses struct {
	Datas []WbkSatkerResponse
}

//Get By Satker-ID and Tahun-ID
type WbkSatkerGetBySatkerIDAndTahunIDResponse struct {
	abstraction.ID
	SatkerID           int    `json:"satker_id"`
	TahunID            int    `json:"tahun_id"`
	Year               string `json:"year"`
	KomponenID         int    `json:"komponen_id"`
	Komponen           string `json:"komponen"`
	ProgramID          int    `json:"program_id"`
	Program            string `json:"program"`
	ProgramRenjaID     int    `json:"program_renja_id"`
	ProgramRenja       string `json:"program_renja"`
	SubProgramRenjaID  int    `json:"sub_program_renja_id"`
	SubProgramRenja    string `json:"sub_program_renja"`
	SubProgramUraianID int    `json:"sub_program_uraian_id"`
	SubProgramUraian   string `json:"sub_program_uraian"`
	B1                 bool   `json:"b1"`
	B2                 bool   `json:"b2"`
	B3                 bool   `json:"b3"`
	B4                 bool   `json:"b4"`
	B5                 bool   `json:"b5"`
	B6                 bool   `json:"b6"`
	B7                 bool   `json:"b7"`
	B8                 bool   `json:"b8"`
	B9                 bool   `json:"b9"`
	B10                bool   `json:"b10"`
	B11                bool   `json:"b11"`
	B12                bool   `json:"b12"`
	FrekuensiWaktu     string `json:"frekuensi_waktu"`
}

type WbkSatkerGetBySatkerIDAndTahunIDInfoResponse struct {
	Datas          []WbkSatkerGetBySatkerIDAndTahunIDResponse
	PaginationInfo *abstraction.PaginationInfo
}

//Get Program
type WbkSatkerPogramGetResponse struct {
	abstraction.ID
	wbk.WbkSatkerEntity
	WbkKomponenName string `json:"wbk_komponen_name"`
	WbkProgramID    int    `json:"wbk_program_id"`
	WbkProgramName  string `json:"wbk_program_name"`
	Komponen        string `json:"komponen"`
	Program         string `json:"program"`
	SatkerID        int    `json:"satker_id"`
	Satker          string `json:"satker"`
	ThnAngID        int    `json:"thn_ang_id"`
	Year            string `json:"year"`
}

type WbkSatkerPogramGetInfoResponse struct {
	Datas          []WbkSatkerPogramGetResponse
	PaginationInfo *abstraction.PaginationInfo
}
