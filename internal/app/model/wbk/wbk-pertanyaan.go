package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkPertanyaanEntityFilter struct {
	ThnAngID           *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID           *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	ID                 *int    `json:"id" query:"id" alias:"wpt"`
	WbkProgramRankerID int     `json:"wbk_program_ranker_id"`
	Code               *string `json:"code" query:"code"`
	Name               *string `json:"name" query:"name" filter:"LIKE"`
	Target             *string `json:"target" query:"target" filter:"LIKE"`
}

type WbkPertanyaanEntity struct {
	WbkProgramRankerID int    `json:"wbk_program_ranker_id"`
	Code               string `json:"code,omitempty"`
	Name               string `json:"name"`
	Target             string `json:"target"`
}

type WbkPertanyaan struct {
	abstraction.IDInc
	WbkPertanyaanEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkPertanyaanFilter struct {
	WbkPertanyaanEntityFilter
}

func (m *WbkPertanyaan) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkPertanyaan) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
