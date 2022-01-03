package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSubProgramUraianEntityFilter struct {
	WbkSubProgramRankerID *int    `json:"wbk_sub_program_ranker_id" query:"wbk_sub_program_ranker_id" filter:"NOFILTER"`
	FrekuensiRankerID     *int    `json:"frekuensi_ranker_id" query:"frekuensi_ranker_id"`
	ID                    *int    `json:"id" query:"id" alias:"wspu"`
	Code                  *int    `json:"code" query:"code"`
	Name                  *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkSubProgramUraianEntity struct {
	WbkSubProgramRankerID int    `json:"wbk_sub_program_ranker_id"`
	FrekuensiRankerID     int    `json:"frekuensi_ranker_id"`
	Code                  int    `json:"code"`
	Name                  string `json:"name"`
	Ket                   string `json:"ket"`
}

type WbkSubProgramUraian struct {
	abstraction.IDInc
	WbkSubProgramUraianEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkSubProgramUraianFilter struct {
	WbkSubProgramUraianEntityFilter
}

func (m *WbkSubProgramUraian) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkSubProgramUraian) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
