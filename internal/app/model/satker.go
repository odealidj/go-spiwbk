package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SatkerEntityFilter struct {
	UnitID    *uint16 `json:"unit_id" query:"unit_id"`
	KabkotaID *uint16 `json:"kabkota_id" query:"kabkota_id"`
	Code      *string `json:"code" query:"code"`
	Name      *string `json:"name" query:"name" filter:"LIKE"`
}

type SatkerEntity struct {
	UnitID    uint16 `json:"unit_id"`
	KabkotaID uint16 `json:"kabkota_id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
}

type Satker struct {
	abstraction.EntityInc
	SatkerEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SatkerFilter struct {
	SatkerEntityFilter
}

func (m *Satker) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Satker) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
