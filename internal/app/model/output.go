package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type OutputEntityFilter struct {
	Code string `json:"code" query:"code"`
	Name string `json:"name" query:"name" filter:"LIKE"`
}

type OutputEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Output struct {
	abstraction.EntityInc
	OutputEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type OutputFilter struct {
	OutputEntityFilter
}

func (m *Output) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Output) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
