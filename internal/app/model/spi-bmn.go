package model

import (
	"codeid-boiler/internal/abstraction"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SpiBmnEntityFilter struct {
	ThnAngID *int `json:"thnAngID" query:"thnAngID" filter:"NOFILTER"`
	SatkerID *int `json:"satkerID" query:"satkerID" filter:"NOFILTER"`
}

type SpiBmnEntity struct {
	SpiAngID                int             `json:"spiAngID"`
	JenisBmnID              int             `json:"jenisBmnID"`
	JenisBmnUraian          string          `json:"jenisBmnUraian"`
	NilaiBmn                decimal.Decimal `json:"nilaiBmn"`
	PengelolaBmnSatkerID    string          `json:"pengelolaBmnSatkerID"`
	PengelolaBmnPihakTigaID int             `json:"pengelolaBmnPihakTigaID"`
	PengelolaBmnKsoID       int             `json:"pengelolaBmnKsoID"`
	PermasalahanBmnID       int             `json:"permasalahanBmnID"`
	UraianPemasalahan       string          `json:"uraianPemasalahan"`
	RencanaPemecahan        string          `json:"rencanaPemecahan"`
	RealisasiPemecahan      string          `json:"realisasiPemecahan"`
}

type SpiBmn struct {
	abstraction.IDInc
	SpiBmnEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiBmnFilter struct {
	SpiBmnEntityFilter
}

func (m *SpiBmn) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiBmn) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
