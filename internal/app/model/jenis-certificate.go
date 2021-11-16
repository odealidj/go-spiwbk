package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type JenisCertificateEntityFilter struct {
	Name *string `json:"name" query:"name" filter:"LIKE"`
}

type JenisCertificateEntity struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

type JenisCertificate struct {
	abstraction.EntityInc
	JenisCertificateEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type JenisCertificateFilter struct {
	JenisCertificateEntityFilter
}

func (m *JenisCertificate) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY

	return
}

func (m *JenisCertificate) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
