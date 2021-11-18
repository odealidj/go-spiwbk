package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"

	"gorm.io/gorm"
)

type UserAppEntity struct {
	RoleUserId   uint16 `json:"role_user_id"`
	SatkerId     uint16 `json:"satker_id"`
	JabatanId    uint16 `json:"jabatan_id"`
	Nip          string `json:"nip" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	MobileNumber string `json:"mobile_number" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
}

type UserApp struct {
	// abstraction
	abstraction.Entity

	// entity
	UserAppEntity

	//Relation
	//LoginAppID uint16 `json:"id" validate:"required" gorm:"index:idx_user_app_login_app_id,unique;not null"`
	//LoginApp LoginApp

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *UserAppEntity) BeforeCreate(tx *gorm.DB) (err error) {

	return
}

func (m *UserApp) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *UserApp) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	//m.ModifiedBy = &m.Context.Auth.Name
	return
}
