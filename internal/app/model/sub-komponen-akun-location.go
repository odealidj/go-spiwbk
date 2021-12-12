package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SubKomponenAkunLocationEntity struct {
	SubKomponenAkunID int    `json:"sub_komponen_akun_id"`
	Name              string `json:"name"`
}

type SubKomponenAkunLocation struct {
	abstraction.EntityInc
	SubKomponenAkunLocationEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *SubKomponenAkunLocation) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SubKomponenAkunLocation) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
