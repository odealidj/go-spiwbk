package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiSdmEntityFilter struct {
	ThnAgnID *uint16 `json:"thn_agn_id" query:"thn_agn_id"`
	SatkerID *uint16 `json:"satker_id" query:"satker_id"`
}

type SpiSdmEntity struct {
	ThnAgnID uint16 `json:"thn_agn_id"`
	SatkerID uint16 `json:"satker_id"`
}

type SpiSdm struct {
	abstraction.EntityInc
	SpiSdmEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiSdmFilter struct {
	SpiSdmEntityFilter
}

func (m *SpiSdm) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SpiSdm) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
