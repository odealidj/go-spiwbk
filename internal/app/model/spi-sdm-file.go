package model

import "codeid-boiler/internal/abstraction"

type SpiSdmFileEntity struct {
	FilePath string `json:"file_path"`
	Title    string `json:"title"`
}

type SpiSdmFile struct {
	abstraction.Entity
	SpiSdmFileEntity

	SpiSdm SpiSdm `gorm:"foreignKey:id"`

	Context *abstraction.Context `json:"-" gorm:"-"`
}
