package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type MethodApbjEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type MethodApbjEntity struct {
	Name string `json:"name"`
}

type MethodApbj struct {
	abstraction.ID
	abstraction.DeleteAt
	MethodApbjEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type MethodApbjFilter struct {
	MethodApbjEntityFilter
}

func (m *MethodApbj) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *MethodApbj) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
