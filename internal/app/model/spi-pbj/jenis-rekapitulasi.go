package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type JenisRekapitulasiEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisRekapitulasiEntity struct {
	Name string `json:"name"`
}

type JenisRekapitulasi struct {
	abstraction.ID
	abstraction.DeleteAt
	JenisRekapitulasiEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisRekapitulasiFilter struct {
	JenisRekapitulasiEntityFilter
}

func (m *JenisRekapitulasi) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisRekapitulasi) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
