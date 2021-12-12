package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type ProgramEntityFilter struct {
	Code string `json:"code" query:"code"`
	Name string `json:"name" query:"name" filter:"LIKE"`
}

type ProgramEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Program struct {
	abstraction.EntityInc
	ProgramEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type ProgramFilter struct {
	ProgramEntityFilter
}

func (m *Program) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Program) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
