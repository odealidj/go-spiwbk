package model

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/pkg/constant"
	"codeid-boiler/pkg/util/date"
	"gorm.io/gorm"
)

type SpiSdmItemEntityFilter struct {
	Name              *string `json:"name" query:"name" filter:"LIKE"`
	SpiSdmID          *uint16 `json:"spi_sdm_id" query:"spi_sdm_id"`
	JenisSdmID        *uint16 `json:"jenis_sdm_id" query:"jenis_sdm_id"`
	PegawaiID         *uint16 `json:"pegawai_id" query:"pegawai_id"`
	JenisCertficateID *uint16 `json:"jenis_certficate_id" query:"jenis_certficate_id"`
	Nosk              *string `json:"nosk" query:"nosk"`
}

type SpiSdmItemEntity struct {
	SpiSdmID          uint16 `json:"spi_sdm_id"`
	JenisSdmID        uint16 `json:"jenis_sdm_id"`
	PegawaiID         uint16 `json:"pegawai_id"`
	JenisCertficateID uint16 `json:"jenis_certficate_id"`
	Nosk              string `json:"nosk"`
	TglSk             string `json:"tgl_sk"`
}

type SpiSdmItem struct {
	abstraction.EntityInc
	SpiSdmItemEntity
	Context *abstraction.Context `json:"-" gorm:"-"`
}

type SpiSdmItemFilter struct {
	SpiSdmItemEntityFilter
}

func (m *SpiSdmItem) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	m.CreatedBy = constant.DB_DEFAULT_CREATED_BY
	return
}

func (m *SpiSdmItem) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	m.ModifiedBy = &m.Context.Auth.Name
	return
}
