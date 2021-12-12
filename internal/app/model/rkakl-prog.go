package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type RkaklProgEntity struct {
	RkaklID   int             `json:"rkakl_id"`
	ProgramID int             `json:"program_id"`
	Biaya     decimal.Decimal `json:"biaya"`
}

type RkaklProg struct {
	abstraction.EntityInc
	RkaklProgEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *RkaklProg) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *RkaklProg) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
