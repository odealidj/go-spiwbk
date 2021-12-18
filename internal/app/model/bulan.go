package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type BulanEntityFilter struct {
	Code *string `json:"code" query:"code"`
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type BulanEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type Bulan struct {
	abstraction.ID
	abstraction.DeleteAt
	BulanEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type BulanFilter struct {
	BulanEntityFilter
}

func (m *Bulan) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Bulan) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
