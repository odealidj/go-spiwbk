package wbk

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type WbkKomponenEntityFilter struct {
	ThnAngID *int    `json:"thn_ang_id" query:"thn_ang_id" filter:"NOFILTER"`
	SatkerID *int    `json:"satker_id" query:"satker_id" filter:"NOFILTER"`
	Code     *string `json:"code" query:"code"`
	Name     *string `json:"name" query:"name" filter:"LIKE"`
}

type WbkKomponenEntity struct {
	Code string `json:"code"`
	Name string `json:"name"`
	Tag  string `json:"tag"`
}

type WbkKomponen struct {
	abstraction.ID
	WbkKomponenEntity
	abstraction.DeleteAt
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type WbkKomponenFilter struct {
	WbkKomponenEntityFilter
}

func (m *WbkKomponen) BeforeCreate(tx *gorm.DB) (err error) {
	//m.CreatedAt = *date.DateTodayLocal()
	//m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *WbkKomponen) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
