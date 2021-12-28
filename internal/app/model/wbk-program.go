package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkProgramEntityFilter struct {
	ThnAngID      *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID      *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkKomponenID *int    `json:"wbk_komponen_id" query:"wbk_komponen_id"`
	Code          *string `json:"code" query:"code"`
	Name          *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkProgramEntity struct {
	WbkKomponenID int    `json:"wbk_komponen_id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	Tag           string `json:"tag"`
}

type WbkProgram struct {
	abstraction.IDInc
	WbkProgramEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkProgramFilter struct {
	WbkProgramEntityFilter
}

func (m *WbkProgram) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkProgram) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
