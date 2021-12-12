package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RkaklItemEntity struct {
	SubKomponenAkunID int             `json:"sub_komponen_akun_id"`
	Name              string          `json:"name"`
	Volume            string          `json:"volume"`
	Harga             decimal.Decimal `json:"harga"`
	Biaya             decimal.Decimal `json:"biaya"`
}

type RkaklItem struct {
	abstraction.EntityInc
	RkaklItemEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *RkaklItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *RkaklItem) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
