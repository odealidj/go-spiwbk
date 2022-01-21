package model

import (
	"codeid-boiler/internal/abstraction"
	"gorm.io/gorm"
)

type ThnAngSatkerEntityFilter struct {
	ThnAngID *int `json:"thn_ang_id" query:"thn_ang_id"`
	SatkerID *int `json:"satker_id" query:"satker_id"`
}

type ThnAngSatkerEntity struct {
	ThnAngID int `json:"thn_ang_id"`
	SatkerID int `json:"satker_id"`
}

type ThnAngSatker struct {
	abstraction.IDInc
	ThnAngSatkerEntity
	abstraction.DeleteAt
	//SpiSdm  []SpiSdm
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type ThnAngSatkerFilter struct {
	ThnAngSatkerEntityFilter
}

func (m *ThnAngSatker) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

func (m *ThnAngSatker) BeforeUpdate(tx *gorm.DB) (err error) {
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
