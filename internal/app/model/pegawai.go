package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type PegawaiEntityFilter struct {
	SatkerID *uint16 `json:"satker_id" query:"satker_id"`
	Nik      *string `json:"nik" query:"nik"`
	Name     *string `json:"name" query:"name" filter:"LIKE"`
}

type PegawaiEntity struct {
	SatkerID           uint16 `json:"satker_id"`
	Nik                string `json:"nik"`
	Name               string `json:"name"`
	PendidikanTerakhir string `json:"pendidikan_terakhir"`
}

type Pegawai struct {
	abstraction.EntityInc
	PegawaiEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type PegawaiFilter struct {
	PegawaiEntityFilter
}

func (m *Pegawai) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Pegawai) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
