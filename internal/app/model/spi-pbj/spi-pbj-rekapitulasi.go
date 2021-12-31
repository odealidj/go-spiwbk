package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type SpiPbjRekapitulasiEntityFilter struct {
	ID       *int `json:"ID" query:"ID" alias:"spr"`
	SpiAngID *int `json:"spiAngID" query:"spi_ang_id" alias:"spr"`
	ThnAngID *int `json:"thnAngID" query:"thn_ang_id" alias:"sa"`
	SatkerID *int `json:"satkerID" query:"satker_id" alias:"sa"`
}

type SpiPbjRekapitulasiEntity struct {
	SpiAngID            int     `json:"spi_ang_id"`
	JenisRekapitulasiID int     `json:"jenis_rekapitulasi_id"`
	BulanID             int     `json:"bulan_id"`
	Target              float64 `json:"target"`
}

type SpiPbjRekapitulasi struct {
	abstraction.IDInc
	SpiPbjRekapitulasiEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiPbjRekapitulasiFilter struct {
	SpiPbjRekapitulasiEntityFilter
}

func (m *SpiPbjRekapitulasi) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiPbjRekapitulasi) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
