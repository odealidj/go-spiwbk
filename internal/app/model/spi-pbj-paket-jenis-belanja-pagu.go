package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type SpiPbjPaketJenisBelanjaPaguEntityFilter struct {
	ThnAngID            *int `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID            *int `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	GroupPackageValueID *int `json:"group_package_value_id" query:"group_package_value_id" filter:"NOFILTER"`
}

type SpiPbjPaketJenisBelanjaPaguEntity struct {
	SpiPbjPaketID      int `json:"spiPbjPaketID"`
	JenisBelanjaPaguID int `json:"jenisBelanjaPaguID"`
	SubKomponenAkunID  int `json:"subKomponenAkunID"`
}

type SpiPbjPaketJenisBelanjaPagu struct {
	abstraction.ID
	SpiPbjPaketJenisBelanjaPaguEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiPbjPaketJenisBelanjaPaguFilter struct {
	SpiPbjPaketJenisBelanjaPaguEntityFilter
}

func (m *SpiPbjPaketJenisBelanjaPagu) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiPbjPaketJenisBelanjaPagu) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
