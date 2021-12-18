package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type GroupPackageValueEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type GroupPackageValueEntity struct {
	Name string `json:"name"`
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
