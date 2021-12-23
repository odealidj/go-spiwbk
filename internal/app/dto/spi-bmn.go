package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"github.com/shopspring/decimal"
)

type SpiBmnSaveRequest struct {
	ThnAngID int `json:"thnAngID"`
	SatkerID int `json:"satkerID"`
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
	Row                 int             `json:"row"`
	SpiAngID            int             `json:"spiAngID"`
	ThnAngID            int             `json:"thnAngID"`
	SatkerID            int             `json:"satkerID"`
	GroupPackageValueID int             `json:"groupPackageValueID"`
	PaketName           string          `json:"paketName"`
	Barang              decimal.Decimal `json:"barang"`
	Modal               decimal.Decimal `json:"modal"`
	Sosial              decimal.Decimal `json:"sosial"`
	Lainnya             decimal.Decimal `json:"lainnya"`
	MethodPbj           string          `json:"methodPbj"`
	Rencana1            bool            `json:"rencana1"`
	Rencana2            bool            `json:"rencana2"`
	Rencana3            bool            `json:"rencana3"`
	Rencana4            bool            `json:"rencana4"`
	Rencana5            bool            `json:"rencana5"`
	Rencana6            bool            `json:"rencana6"`
	Rencana7            bool            `json:"rencana7"`
	Rencana8            bool            `json:"rencana8"`
	Rencana9            bool            `json:"rencana9"`
	Rencana10           bool            `json:"rencana10"`
	Rencana11           bool            `json:"rencana11"`
	Rencana12           bool            `json:"rencana12"`
	Realisasi1          bool            `json:"realisasi1"`
	Realisasi2          bool            `json:"realisasi2"`
	Realisasi3          bool            `json:"realisasi3"`
	Realisasi4          bool            `json:"realisasi4"`
	Realisasi5          bool            `json:"realisasi5"`
	Realisasi6          bool            `json:"realisasi6"`
	Realisasi7          bool            `json:"realisasi7"`
	Realisasi8          bool            `json:"realisasi8"`
	Realisasi9          bool            `json:"realisasi9"`
	Realisasi10         bool            `json:"realisasi10"`
	Realisasi11         bool            `json:"realisasi11"`
	Realisasi12         bool            `json:"realisasi12"`
	Permasalahan        string          `json:"permasalahan"`
	RencanaPemecahan    string          `json:"rencanaPemecahan"`
}

type SpiBmnGetReformatResponse struct {
	Num                 int              `json:"num"`
	Row                 string           `json:"row"`
	SpiAngID            *int             `json:"spiAngID"`
	ThnAngID            *int             `json:"thnAngID"`
	SatkerID            *int             `json:"satkerID"`
	GroupPackageValueID *int             `json:"groupPackageValueID"`
	PaketName           string           `json:"paketName"`
	Barang              *decimal.Decimal `json:"barang"`
	Modal               *decimal.Decimal `json:"modal"`
	Sosial              *decimal.Decimal `json:"sosial"`
	Lainnya             *decimal.Decimal `json:"lainnya"`
	MethodPbj           string           `json:"methodPbj"`
	Rencana1            *bool            `json:"rencana1"`
	Rencana2            *bool            `json:"rencana2"`
	Rencana3            *bool            `json:"rencana3"`
	Rencana4            *bool            `json:"rencana4"`
	Rencana5            *bool            `json:"rencana5"`
	Rencana6            *bool            `json:"rencana6"`
	Rencana7            *bool            `json:"rencana7"`
	Rencana8            *bool            `json:"rencana8"`
	Rencana9            *bool            `json:"rencana9"`
	Rencana10           *bool            `json:"rencana10"`
	Rencana11           *bool            `json:"rencana11"`
	Rencana12           *bool            `json:"rencana12"`
	Realisasi1          *bool            `json:"realisasi1"`
	Realisasi2          *bool            `json:"realisasi2"`
	Realisasi3          *bool            `json:"realisasi3"`
	Realisasi4          *bool            `json:"realisasi4"`
	Realisasi5          *bool            `json:"realisasi5"`
	Realisasi6          *bool            `json:"realisasi6"`
	Realisasi7          *bool            `json:"realisasi7"`
	Realisasi8          *bool            `json:"realisasi8"`
	Realisasi9          *bool            `json:"realisasi9"`
	Realisasi10         *bool            `json:"realisasi10"`
	Realisasi11         *bool            `json:"realisasi11"`
	Realisasi12         *bool            `json:"realisasi12"`
	Permasalahan        string           `json:"permasalahan"`
	RencanaPemecahan    string           `json:"rencanaPemecahan"`
}

type SpiBmnGetInfoResponse struct {
	Datas          *[]SpiBmnGetReformatResponse
	PaginationInfo *abstraction.PaginationInfo
}

type SpiBpmResponse struct {
	abstraction.ID
	model.SpiBmnEntity
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
}
