package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type AppEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type AppEntity struct {
	Name string `json:"name"`
}

type App struct {
	abstraction.ID
	AppEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type AppFilter struct {
	AppEntityFilter
}

func (m *App) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *App) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
