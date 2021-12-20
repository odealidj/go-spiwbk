package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProgKegiatanEntity struct {
	RkaklProgID int             `json:"rkakl_prog_id"`
	KegiatanID  int             `json:"kegiatan_id"`
	Biaya       decimal.Decimal `json:"biaya"`
}

type ProgKegiatan struct {
	abstraction.EntityInc
	ProgKegiatanEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *ProgKegiatan) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *ProgKegiatan) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
