package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SubKomponenEntity struct {
	KomponenID int             `json:"komponen_id"`
	Code       string          `json:"code"`
	Name       string          `json:"name"`
	Biaya      decimal.Decimal `json:"biaya"`
}

type SubKomponen struct {
	abstraction.EntityInc
	SubKomponenEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *SubKomponen) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SubKomponen) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
