package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type FrekuensiRankerEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type FrekuensiRankerEntity struct {
	Name string `json:"name"`
}

type FrekuensiRanker struct {
	abstraction.ID
	FrekuensiRankerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type FrekuensiRankerFilter struct {
	FrekuensiRankerEntityFilter
}

func (m *FrekuensiRanker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *FrekuensiRanker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
