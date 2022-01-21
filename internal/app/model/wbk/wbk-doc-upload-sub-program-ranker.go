package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkDocUploadSubProgramRankerEntityFilter struct {
	ThnAngID                   *int `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID                   *int `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkSubProgramRankerBulanID *int `json:"wbk_sub_program_ranker_bulan_id" query:"wbk_sub_program_ranker_bulan_id"`
	ID                         *int `json:"id" query:"id" alias:"wspr"`
}

type WbkDocUploadSubProgramRankerEntity struct {
	WbkSubProgramRankerBulanID int     `json:"wbk_sub_program_ranker_bulan_id"`
	Path                       string  `json:"path"`
	Ket                        *string `json:"ket"`
}

type WbkDocUploadSubProgramRanker struct {
	abstraction.IDInc
	WbkDocUploadSubProgramRankerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkDocUploadSubProgramRankerFilter struct {
	WbkDocUploadSubProgramRankerEntityFilter
}

func (m *WbkDocUploadSubProgramRanker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkDocUploadSubProgramRanker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
