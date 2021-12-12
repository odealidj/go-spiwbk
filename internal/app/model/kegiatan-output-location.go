package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type KegiatanOutputLocationEntity struct {
	KegiatanOutputID int    `json:"kegiatan_output_id"`
	Name             string `json:"name"`
}

type KegiatanOutputLocation struct {
	abstraction.EntityInc
	KegiatanOutputLocationEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *KegiatanOutputLocation) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *KegiatanOutputLocation) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
