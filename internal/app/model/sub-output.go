package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SubOutputEntity struct {
	KegiatanOutputLocationID int             `json:"kegiatan_output_location_id"`
	Code                     string          `json:"code"`
	Name                     string          `json:"name"`
	Volume                   string          `json:"volume"`
	Biaya                    decimal.Decimal `json:"biaya"`
}

type SubOutput struct {
	abstraction.EntityInc
	SubOutputEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *SubOutput) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *SubOutput) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
