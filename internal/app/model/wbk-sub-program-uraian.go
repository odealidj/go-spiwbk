package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSubProgramUraianEntityFilter struct {
	ThnAngID              *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID              *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkSubProgramRankerID int     `json:"wbk_sub_program_ranker_id" query:"wbk_sub_program_ranker_id"`
	BulanID               int     `json:"bulan_id" query:"bulan_id"`
	FrekuensiRankerID     int     `json:"frekuensi_ranker_id" query:"frekuensi_ranker_id"`
	Code                  *string `json:"code" query:"code"`
	Name                  *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkSubProgramUraianEntity struct {
	WbkSubProgramRankerID int    `json:"wbk_sub_program_ranker_id"`
	BulanID               int    `json:"bulan_id"`
	FrekuensiRankerID     int    `json:"frekuensi_ranker_id"`
	Code                  string `json:"code"`
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
