package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkProgramRankerEntityFilter struct {
	ThnAngID     *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID     *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkProgramID *int    `json:"wbk_program_id" query:"wbk_program_id"`
	Code         *string `json:"code" query:"code"`
	Name         *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkProgramRankerEntity struct {
	WbkProgramID int    `json:"wbk_program_id"`
	Code         int    `json:"code"`
	Name         string `json:"name"`
	Tag          string `json:"tag"`
}

type WbkProgramRanker struct {
	abstraction.IDInc
	WbkProgramRankerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkProgramRankerFilter struct {
	WbkProgramRankerEntityFilter
}

func (m *WbkProgramRanker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkProgramRanker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
