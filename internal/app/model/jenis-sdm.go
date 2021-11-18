package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type JenisSdmEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisSdmEntity struct {
	Name string `json:"name"`
}

type JenisSdm struct {
	abstraction.EntityInc
	JenisSdmEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisSdmFilter struct {
	JenisSdmEntityFilter
}

func (m *JenisSdm) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisSdm) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
