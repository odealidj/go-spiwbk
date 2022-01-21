package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkDocUploadSubProgramRankerNilaiEntityFilter struct {
	ThnAngID *int `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID *int `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	ID       *int `json:"id" query:"id" alias:"wspr"`
}

type WbkDocUploadSubProgramRankerNilaiEntity struct {
	Ket   string `json:"ket"`
	Nilai int    `json:"nilai"`
}

type WbkDocUploadSubProgramRankerNilai struct {
	abstraction.ID
	WbkDocUploadSubProgramRankerNilaiEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkDocUploadSubProgramRankerNilaiFilter struct {
	WbkDocUploadSubProgramRankerNilaiEntityFilter
}

func (m *WbkDocUploadSubProgramRankerNilai) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkDocUploadSubProgramRankerNilai) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
