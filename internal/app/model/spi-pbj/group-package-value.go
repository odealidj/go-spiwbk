package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type GroupPackageValueEntityFilter struct {
	ThnAngID *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	Name     *string `json:"name" query:"name" filter:"LIKE"`
}

type GroupPackageValueEntity struct {
	Name     string          `json:"name"`
	MinValue decimal.Decimal `json:"minValue"`
	MaxValue decimal.Decimal `json:"maxValue"`
}

type GroupPackageValue struct {
	abstraction.ID
	abstraction.DeleteAt
	GroupPackageValueEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type GroupPackageValueFilter struct {
	GroupPackageValueEntityFilter
}

func (m *GroupPackageValue) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *GroupPackageValue) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
