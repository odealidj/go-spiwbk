package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type JenisKesesuaianEntityFilter struct {
	Code *string `json:"code" query:"code"`
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisKesesuaianEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type JenisKesesuaian struct {
	abstraction.Entity
	JenisKesesuaianEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisKesesuaianFilter struct {
	JenisKesesuaianEntityFilter
}

func (m *JenisKesesuaian) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisKesesuaian) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
