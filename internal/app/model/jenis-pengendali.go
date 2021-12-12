package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type JenisPengendaliEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisPengendaliEntity struct {
	Name string `json:"name"`
}

type JenisPengendali struct {
	abstraction.Entity
	JenisPengendaliEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisPengendaliFilter struct {
	JenisPengendaliEntityFilter
}

func (m *JenisPengendali) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisPengendali) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
