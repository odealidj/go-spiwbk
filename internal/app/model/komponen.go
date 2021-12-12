package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type KomponenEntity struct {
	SubOutputID int             `json:"sub_output_id"`
	Code        string          `json:"code"`
	Name        string          `json:"name"`
	Volume      string          `json:"volume"`
	Biaya       decimal.Decimal `json:"biaya"`
}

type Komponen struct {
	abstraction.EntityInc
	KomponenEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *Komponen) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Komponen) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
