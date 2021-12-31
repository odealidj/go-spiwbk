package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
)

type SpiPbjPaketSaveRequest struct {
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
	spi_pbj.SpiPbjPaketEntity
}

//upsert
type SpiPbjPaketUpsertRequest struct {
	abstraction.ID
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
	spi_pbj.SpiPbjPaketEntity
}

//Get
type SpiPbjPaketGetRequest struct {
	abstraction.Pagination
	spi_pbj.SpiPbjPaketFilter
}

type SpiPbjPaketGetResponse struct {
	Row                 int     `json:"row"`
	ID                  int     `json:"ID"`
	SpiAngID            int     `json:"spiAngID"`
	JenisRekapitulasiID int     `json:"jenisRekapitulasiID"`
	ThnAngID            int     `json:"thnAngID"`
	Year                string  `json:"year"`
	SatkerID            int     `json:"satkerID"`
	SatkerName          string  `json:"satkerName"`
	PelaksanaanKegiatan string  `json:"pelaksanaanKegiatan"`
	Total               float64 `json:"total"`
	B01                 float64 `json:"b01"`
	B02                 float64 `json:"b02"`
	B03                 float64 `json:"b03"`
	B04                 float64 `json:"b04"`
	B05                 float64 `json:"b05"`
	B06                 float64 `json:"b06"`
	B07                 float64 `json:"b07"`
	B08                 float64 `json:"b08"`
	B09                 float64 `json:"b09"`
	B10                 float64 `json:"b10"`
	B11                 float64 `json:"b11"`
	B12                 float64 `json:"b21"`
}

type SpiPbjPaketGetInfoResponse struct {
	Datas          *[]SpiPbjRekapitulasiGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type SpiPbjPaketResponse struct {
	abstraction.ID
	spi_pbj.SpiPbjPaketEntity
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
}
