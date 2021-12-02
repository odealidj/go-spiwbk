package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type RkaklEntityFilter struct {
	ThnAngID   *int    `json:"thn_ang_id" query:"thn_ang_id"`
	ThnAngYear *string `json:"thn_ang_year" query:"thn_ang_year"`
	SatkerID   *int    `json:"satker_id" query:"satker_id"`
	SatkerName *string `json:"satker_name" query:"satker_name"`
	Title      *string `json:"title" query:"title" filter:"LIKE"`
}

type RkaklEntity struct {
	ThnAngID    int    `json:"thn_ang_id" form:"thn_ang_id"`
	SatkerID    int    `json:"satker_id" form:"satker_id"`
	LoginAppID  int    `json:"login_app_id" form:"login_app_id"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	IsApproved  *bool  `json:"is_approved,omitempty"`
}

type Rkakl struct {
	abstraction.EntityInc
	RkaklEntity
	ThnAng   ThnAng //`gorm:"foreignKey:thn_ang_id"`
	Satker   Satker
	LoginApp LoginApp
	Context  *abstraction.Context `json:"-" gorm:"-"`
}

type RkaklFilter struct {
	RkaklEntityFilter
}

func (m *Rkakl) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *Rkakl) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
