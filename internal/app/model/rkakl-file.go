package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type RkaklFileEntity struct {
	Filepath string `json:"filepath" form:"file"`
}

type RkaklFile struct {
	abstraction.Entity
	RkaklFileEntity

	//Rkakl Rkakl `gorm:"foreignKey:id"`

	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *RkaklFile) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *RkaklFile) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
