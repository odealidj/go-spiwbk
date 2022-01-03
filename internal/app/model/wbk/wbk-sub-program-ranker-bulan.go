package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSubProgramRankerBulanEntityFilter struct {
	WbkSubProgramRankerID *int `json:"wbk_sub_program_ranker_id" query:"wbk_sub_program_ranker_id"`
	BulanID               *int `json:"bulan_id" query:"bulan_id"`
	ID                    *int `json:"id" query:"id" alias:"wspu"`
}

type WbkSubProgramRankerBulanEntity struct {
	WbkSubProgramRankerID int `json:"wbk_sub_program_ranker_id"`
	BulanID               int `json:"bulan_id"`
}

type WbkSubProgramRankerBulan struct {
	abstraction.ID
	WbkSubProgramRankerBulanEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkSubProgramRankerBulanFilter struct {
	WbkSubProgramRankerBulanEntityFilter
}

func (m *WbkSubProgramRankerBulan) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkSubProgramRankerBulan) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
