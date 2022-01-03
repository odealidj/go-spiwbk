package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSubProgramUraianBulanEntityFilter struct {
	WbkSubProgramUraianID *int `json:"wbk_sub_program_uraian_id" query:"wbk_sub_program_uraian_id"`
	BulanID               *int `json:"bulan_id" query:"bulan_id"`
	ID                    *int `json:"id" query:"id" alias:"wspu"`
}

type WbkSubProgramUraianBulanEntity struct {
	WbkSubProgramUraianID int `json:"wbk_sub_program_uraian_id"`
	BulanID               int `json:"bulan_id"`
}

type WbkSubProgramUraianBulan struct {
	abstraction.ID
	WbkSubProgramUraianBulanEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkSubProgramUraianBulanFilter struct {
	WbkSubProgramUraianBulanEntityFilter
}

func (m *WbkSubProgramUraianBulan) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkSubProgramUraianBulan) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
