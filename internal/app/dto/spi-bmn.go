package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"github.com/shopspring/decimal"
)

type SpiBmnResponse struct {
	abstraction.ID
	model.SpiBmnEntity
	Year       *string `json:"year,omitempty"`
	SatkerName *string `json:"satkerName,omitempty"`
}

type SpiBmnSaveRequest struct {
	ThnAngID int `json:"thnAngID"`
	SatkerID int `json:"satkerID"`
	model.SpiBmnEntity
	//JenisKesesuaianID int  `json:"jenisKesesuaianID"`
	//JenisPengendaliID int  `json:"jenisPengendaliID"`
	//IsCheck           bool `json:"isCheck"`
}

//Get
type SpiBmnGetRequest struct {
	abstraction.Pagination
	model.SpiBmnFilter
}

type SpiBmnGetResponse struct {
	Row                        int             `json:"row"`
	ID                         int             `json:"ID"`
	SpiAngID                   int             `json:"spiAngID"`
	ThnAngID                   int             `json:"thnAngID"`
	SatkerID                   int             `json:"satkerID"`
	JenisBmnID                 int             `json:"jenisBmnID"`
	JenisBmnName               string          `json:"jenisBmnName"`
	JenisBmnUraian             string          `json:"jenisBmnUraian"`
	NilaiBmn                   decimal.Decimal `json:"nilaiBmn"`
	PengelolaSatkerID          int             `json:"pengelolaSatkerID"`
	PengelolaSatkerName        string          `json:"pengelolaSatkerName"`
	PengelolaPihakTigaID       int             `json:"pengelolaPihakTigaID"`
	PengelolaPihakTigaName     string          `json:"PengelolaPihakTigaName"`
	PengelolaKsoID             int             `json:"pengelolaKsoID"`
	PengelolaKsoName           string          `json:"pengelolaKsoName"`
	PermasalahanSengketaID     int             `json:"permasalahanSengketaID"`
	PermasalahanSengketaUraian string          `json:"permasalahanSengketaUraian"`
	PermasalahanDokumenID      string          `json:"permasalahanDokumenID"`
	PermasalahanDokumenUraian  string          `json:"PermasalahanDokumenUraian"`
	PermasalahanHilangID       int             `json:"permasalahanHilangID"`
	PermasalahanHilangUraian   string          `json:"permasalahanHilangUraian"`
	PermasalahanRusakID        int             `json:"permasalahanRusakID"`
	PermasalahanRusakUraian    string          `json:"PermasalahanRusakUraian"`
	PermasalahanLainnyaID      int             `json:"permasalahanLainnyaID"`
	PermasalahanLainnyaUraian  string          `json:"permasalahanLainnyaUraian"`
	RencanaPemecahan           string          `json:"rencanaPemecahan"`
	RealisasiPemecahan         string          `json:"realisasiPemecahan"`
}

type SpiBmnGetReformatResponse struct {
	Num                        int             `json:"num"`
	Row                        string          `json:"row"`
	ID                         int             `json:"ID"`
	SpiAngID                   int             `json:"spiAngID"`
	ThnAngID                   int             `json:"thnAngID"`
	SatkerID                   int             `json:"satkerID"`
	JenisBmnID                 int             `json:"jenisBmnID"`
	JenisBmnName               string          `json:"jenisBmnName"`
	JenisBmnUraian             string          `json:"jenisBmnUraian"`
	NilaiBmn                   decimal.Decimal `json:"nilaiBmn"`
	PengelolaSatkerID          int             `json:"pengelolaSatkerID"`
	PengelolaSatkerName        string          `json:"pengelolaSatkerName"`
	PengelolaPihakTigaID       int             `json:"pengelolaPihakTigaID"`
	PengelolaPihakTigaName     string          `json:"PengelolaPihakTigaName"`
	PengelolaKsoID             int             `json:"pengelolaKsoID"`
	PengelolaKsoName           string          `json:"pengelolaKsoName"`
	PermasalahanSengketaID     int             `json:"permasalahanSengketaID"`
	PermasalahanSengketaUraian string          `json:"permasalahanSengketaUraian"`
	PermasalahanDokumenID      string          `json:"permasalahanDokumenID"`
	PermasalahanDokumenUraian  string          `json:"PermasalahanDokumenUraian"`
	PermasalahanHilangID       int             `json:"permasalahanHilangID"`
	PermasalahanHilangUraian   string          `json:"permasalahanHilangUraian"`
	PermasalahanRusakID        int             `json:"permasalahanRusakID"`
	PermasalahanRusakUraian    string          `json:"PermasalahanRusakUraian"`
	PermasalahanLainnyaID      int             `json:"permasalahanLainnyaID"`
	PermasalahanLainnyaUraian  string          `json:"permasalahanLainnyaUraian"`
	RencanaPemecahan           string          `json:"rencanaPemecahan"`
	RealisasiPemecahan         string          `json:"realisasiPemecahan"`
}

type SpiBmnGetInfoResponse struct {
	Datas          *[]SpiBmnGetResponse
	PaginationInfo *abstraction.PaginationInfo
}

type SpiBpmResponse struct {
	abstraction.ID
	model.SpiBmnEntity
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
}
