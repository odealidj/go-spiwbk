package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type ThnAngEntity struct {
	Year string `json:"year" form:"year" gorm:"type:varchar(4);uniqueIndex"`
}

type ThnAng struct {
	abstraction.EntityInc
	ThnAngEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *ThnAng) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *ThnAng) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
