package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SubKomponenAkunEntity struct {
	SubKomponenID int             `json:"sub_komponen_id"`
	AkunID        int             `json:"akun_id"`
	Name          string          `json:"name"`
	Biaya         decimal.Decimal `json:"biaya"`
	Sdcp          string          `json:"sdcp"`
}

type SubKomponenAkun struct {
	abstraction.EntityInc
	SubKomponenAkunEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *SubKomponenAkun) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SubKomponenAkun) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
