package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkDocUploadSubProgramUraianEntityFilter struct {
	ThnAngID                   *int `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID                   *int `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	WbkSubProgramUraianBulanID *int `json:"wbk_sub_program_uraian_bulan_id" query:"wbk_sub_program_uraian_bulan_id"`
	ID                         *int `json:"id" query:"id" alias:"wspr"`
}

type WbkDocUploadSubProgramUraianEntity struct {
	WbkSubProgramUraianBulanID int     `json:"wbk_sub_program_uraian_bulan_id"`
	Path                       string  `json:"path"`
	Ket                        *string `json:"ket"`
}

type WbkDocUploadSubProgramUraian struct {
	abstraction.IDInc
	WbkDocUploadSubProgramUraianEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkDocUploadSubProgramUraianFilter struct {
	WbkDocUploadSubProgramUraianEntityFilter
}

func (m *WbkDocUploadSubProgramUraian) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkDocUploadSubProgramUraian) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
