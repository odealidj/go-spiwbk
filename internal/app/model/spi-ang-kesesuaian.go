package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiAngKesesuaianEntityFilter struct {
	ThnAngID *int `json:"thnAngID" query:"thnAngID" filter:"NOFILTER"`
	SatkerID *int `json:"satkerID" query:"satkerID" filter:"NOFILTER"`
}

type SpiAngKesesuaianEntity struct {
	SpiAngItemID      int  `json:"spiAngItemID"`
	JenisKesesuaianID int  `json:"jenisKesesuaianID"`
	JenisPengendaliID int  `json:"jenisPengendaliID"`
	IsCheck           bool `json:"isCheck"`
}

type SpiAngKesesuaian struct {
	abstraction.EntityInc
	SpiAngKesesuaianEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiAngKesesuaianFilter struct {
	SpiAngKesesuaianEntityFilter
}

func (m *SpiAngKesesuaian) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiAngKesesuaian) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
