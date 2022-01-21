package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSubProgramRankerEntityFilter struct {
	ThnAngID           *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID           *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkProgramRankerID *int    `json:"wbk_program_ranker_id" query:"wbk_program_ranker_id"`
	ID                 *int    `json:"id" query:"id" alias:"wspr"`
	Code               *string `json:"code" query:"code"`
	Name               *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkSubProgramRankerEntity struct {
	WbkProgramRankerID int    `json:"wbk_program_ranker_id"`
	Code               string `json:"code"`
	Name               string `json:"name"`
	FrekuensiRankerID  *int   `json:"frekuensi_Ranker_Id"`
}

type WbkSubProgramRanker struct {
	abstraction.IDInc
	WbkSubProgramRankerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkSubProgramRankerFilter struct {
	WbkSubProgramRankerEntityFilter
}

func (m *WbkSubProgramRanker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkSubProgramRanker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
