package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type JenisBelanjaPaguEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisBelanjaPaguEntity struct {
	Name string `json:"name"`
}

type JenisBelanjaPagu struct {
	abstraction.ID
	abstraction.DeleteAt
	JenisBelanjaPaguEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisBelanjaPaguFilter struct {
	JenisBelanjaPaguEntityFilter
}

func (m *JenisBelanjaPagu) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisBelanjaPagu) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
