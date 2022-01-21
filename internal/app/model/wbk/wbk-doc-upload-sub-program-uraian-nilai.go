package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkDocUploadSubProgramUraianNilaiEntityFilter struct {
	ThnAngID *int `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID *int `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	ID       *int `json:"id" query:"id" alias:"wspr"`
}

type WbkDocUploadSubProgramUraianNilaiEntity struct {
	Ket   string `json:"ket"`
	Nilai int    `json:"nilai"`
}

type WbkDocUploadSubProgramUraianNilai struct {
	abstraction.ID
	WbkDocUploadSubProgramUraianNilaiEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkDocUploadSubProgramUraianNilaiFilter struct {
	WbkDocUploadSubProgramUraianNilaiEntityFilter
}

func (m *WbkDocUploadSubProgramUraianNilai) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkDocUploadSubProgramUraianNilai) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
