package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiAngEntityFilter struct {
	ThnAngID *uint16 `json:"thn_ang_id" query:"thn_ang_id"`
	SatkerID *uint16 `json:"satker_id" query:"satker_id"`
}

type SpiAngEntity struct {
	ThnAngID uint16 `json:"thn_ang_id" form:"thn_ang_id"`
	SatkerID uint16 `json:"satker_id" form:"satker_id"`
}

type SpiAng struct {
	abstraction.EntityInc
	SpiAngEntity
	//Year    string               `json:"year"`
	ThnAng  ThnAng //`gorm:"foreignKey:thn_ang_id"`
	Satker  Satker
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiAngFilter struct {
	SpiAngEntityFilter
}

func (m *SpiAng) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SpiAng) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
