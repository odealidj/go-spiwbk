package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkSatkerEntityFilter struct {
	ThnAngID      *int `json:"thn_ang_id" query:"thn_ang_id"`
	SatkerID      *int `json:"satker_id" query:"satker_id"`
	WbkKomponenID *int `json:"wbk_komponen_id" query:"wbk_komponen_id" alias:"ws"`
}

type WbkSatkerEntity struct {
	ThnAngSatkerID uint32 `json:"thn_ang_satker_id"`
	WbkKomponenID  uint32 `json:"wbk_komponen_id"`
}

type WbkSatker struct {
	abstraction.IDInc
	WbkSatkerEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkSatkerFilter struct {
	WbkSatkerEntityFilter
}

func (m *WbkSatker) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkSatker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
