package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiAngItemEntityFilter struct {
	SpiAngID   int `json:"spiAngID"`
	KomponenID int `json:"komponenID"`
}

type SpiAngItemEntity struct {
	SpiAngID   int `json:"spiAngID"`
	KomponenID int `json:"komponenID"`
}

type SpiAngItem struct {
	abstraction.EntityInc
	SpiAngItemEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiAngItemFilter struct {
	SpiAngItemEntityFilter
}

func (m *SpiAngItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiAngItem) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
