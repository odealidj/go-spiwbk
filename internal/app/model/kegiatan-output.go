package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type KegiatanOutputEntity struct {
	ProgKegiatanID int             `json:"prog_kegiatan_id"`
	OutputID       int             `json:"output_id"`
	Volume         string          `json:"volume"`
	Biaya          decimal.Decimal `json:"biaya"`
}

type KegiatanOutput struct {
	abstraction.EntityInc
	KegiatanOutputEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *KegiatanOutput) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *KegiatanOutput) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
