package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiPbjPaketEntityFilter struct {
	ThnAngID *int `json:"thnAngID" query:"thnAngID" filter:"NOFILTER"`
	SatkerID *int `json:"satkerID" query:"satkerID" filter:"NOFILTER"`
}

type SpiPbjPaketEntity struct {
	SpiAngID            int    `json:"spiAngID"`
	GroupPackageValueID int    `json:"groupPackageValueID"`
	KomponenID          int    `json:"komponenID"`
	MethodApbjID        *int   `json:"methodApbjID"`
	Permasalahan        string `json:"permasalahan"`
	RencanaPemecahan    string `json:"rencanaPemecahan"`
}

type SpiPbjPaket struct {
	abstraction.EntityInc
	SpiPbjPaketEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiPbjPaketFilter struct {
	SpiPbjPaketEntityFilter
}

func (m *SpiPbjPaket) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiPbjPaket) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
