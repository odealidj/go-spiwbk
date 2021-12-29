package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type BulanRankerEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type BulanRankerEntity struct {
	Name string `json:"name"`
}

type BulanRanker struct {
	abstraction.ID
	BulanRankerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type BulanRankerFilter struct {
	BulanRankerEntityFilter
}

func (m *BulanRanker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *BulanRanker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
