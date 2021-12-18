package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Save
type SpiAngKesesuaianSaveRequest struct {
	SatkerID          int  `json:"satkerID"`
	ThnAngID          int  `json:"thnAngID"`
	JenisKesesuaianID int  `json:"jenisKesesuaianID"`
	JenisPengendaliID int  `json:"jenisPengendaliID"`
	IsCheck           bool `json:"isCheck"`
}

//Get
type SpiAngKesesuaianGetRequest struct {
	abstraction.Pagination
	model.SpiAngKesesuaianFilter
}

type SpiAngKesesuaianGetResponse struct {
	Row                           int    `json:"row"`
	SpiAngKesesuaianID            int    `json:"spiAngKesesuaianID"`
	SpiAngItemID                  int    `json:"spiAngItemID"`
	JenisKesesuaianID             int    `json:"jenisKesesuaianID"`
	JenisPengendaliID             int    `json:"jenisPengendaliID"`
	SpiAngID                      int    `json:"spiAngID"`
	KomponenID                    int    `json:"komponenID"`
	ThnAngID                      int    `json:"thnAngID"`
	SatkerID                      int    `json:"satkerID"`
	ProgramKegiatanOutputKomponen string `json:"programKegiatanOutputKomponen"`
	JenisKendaliName              string `json:"jenisKendaliName"`
	PengusulYa                    bool   `json:"pengusulYa"`
	PengusulTidak                 bool   `json:"pengusulTidak"`
	KeuSatkerYa                   bool   `json:"keuSatkerYa"`
	KeuSatkerTidak                bool   `json:"keuSatkerTidak"`
	KeuEselon1Ya                  bool   `json:"keuEselon1Ya"`
	KeuEselon1Tidak               bool   `json:"keuEselon1Tidak"`
}

type SpiAngKesesuaianGetInfoResponse struct {
	Datas          *[]SpiAngKesesuaianGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type SpiAngKesesuaianResponse struct {
	abstraction.ID
	model.SpiAngKesesuaianEntity
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
}
