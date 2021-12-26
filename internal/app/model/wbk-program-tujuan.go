package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkProgramTujuanEntityFilter struct {
	ThnAngID *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	Code     *string `json:"code" query:"code"`
	Name     *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkProgramTujuanEntity struct {
	WbkProgramID int    `json:"wbk_program_id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
}

type WbkProgramTujuan struct {
	abstraction.ID
	WbkProgramTujuanEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkProgramTujuanFilter struct {
	WbkProgramTujuanEntityFilter
}

func (m *WbkProgramTujuan) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkProgramTujuan) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
