package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type KegiatanEntityFilter struct {
	Code string `json:"code" query:"code"`
	Name string `json:"name" query:"name" filter:"LIKE"`
}

type KegiatanEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Kegiatan struct {
	abstraction.EntityInc
	KegiatanEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type KegiatanFilter struct {
	KegiatanEntityFilter
}

func (m *Kegiatan) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Kegiatan) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
