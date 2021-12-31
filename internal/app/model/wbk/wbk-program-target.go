package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkProgramTargetEntityFilter struct {
	ThnAngID           *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID           *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	ID                 *int    `json:"id" query:"id" alias:"wpt"`
	WbkProgramTujuanID *int    `json:"wbk_program_tujuan_id" alias:"wpt"`
	Code               *string `json:"code" query:"code"`
	Name               *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkProgramTargetEntity struct {
	WbkProgramTujuanID int    `json:"wbk_program_tujuan_id"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name"`
}

type WbkProgramTarget struct {
	abstraction.IDInc
	WbkProgramTargetEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkProgramTargetFilter struct {
	WbkProgramTargetEntityFilter
}

func (m *WbkProgramTarget) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkProgramTarget) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
