package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type SpiSdmItemResponse struct {
	abstraction.ID
	model.SpiSdmItemEntity
}

//Save
type SpiSdmItemSaveRequest struct {
	model.SpiSdmItemEntity
}

//Update
type SpiSdmItemUpdateRequest struct {
	abstraction.ID
	model.SpiSdmItemEntity
}

//Delete
type SpiSdmItemDeleteRequest struct {
	abstraction.ID
}

//Get
type SpiSdmItemGetRequest struct {
	abstraction.Pagination
	model.SpiSdmItemFilter
}

type SpiSdmItemGetResponse struct {
	Datas          *[]SpiSdmItemResponse
	PaginationInfo *abstraction.PaginationInfo
}

//Report
type SpiSdmItemViewBySpiSdmIDRequest struct {
	SpiSdmID uint16 `json:"spi_sdm_id"`
}

type SpiSdmItemViewBySpiSdmIDResponse struct {
	SpiSdmID             uint16 `json:"spi_sdm_id"`
	Uraian               string `json:"uraian"`
	KPA                  string `json:"kpa"`
	PPK                  string `json:"ppk"`
	PPSM                 string `json:"ppsm"`
	BendaharaPengeluaran string `json:"bendahara_pengeluaran"`
	BendaharaPenerimaan  string `json:"bendahara_penerimaan"`
}
