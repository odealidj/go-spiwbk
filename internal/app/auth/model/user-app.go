package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"

	"gorm.io/gorm"
)

type UserAppEntity struct {
	RoleUserId   int    `json:"role_user_id"`
	SatkerId     int    `json:"satker_id"`
	JabatanId    int    `json:"jabatan_id"`
	Nip          string `json:"nip" validate:"required" gorm:"index:idx_user_app_nip,unique"`
	Name         string `json:"name" validate:"required"`
	Address      string `json:"address" validate:"required"`
	PostalCode   string `json:"postal_code" validate:"required"`
	PhoneNumber  string `json:"phone_number" validate:"required"`
	MobileNumber string `json:"mobile_number" validate:"required"`
	Email        string `json:"email" validate:"required,email"`
	//Relation
	LoginApp []LoginApp `json:"login_app" gorm:"foreignKey:user_app_id"`
}

type UserApp struct {
	// abstraction
	abstraction.EntityInc

	// entity
	UserAppEntity

	// context
	Context *abstraction.Context `json:"-" gorm:"-"`
}

func (m *UserApp) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *UserApp) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
