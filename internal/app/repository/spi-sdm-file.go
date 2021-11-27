package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
)

type SpiSdmFile interface {
	Create(*abstraction.Context, *model.SpiSdmFile) (*model.SpiSdmFile, error)
	//Update(*abstraction.Context, *model.SpiSdm) (*model.SpiSdm, error)
	//Delete(*abstraction.Context, *model.SpiSdm) (*model.SpiSdm, error)
	//FindByID(*abstraction.Context, *model.SpiSdm) (*model.SpiSdm, error)
	//Find(*abstraction.Context, *model.SpiSdmFilter, *abstraction.Pagination) (*[]model.SpiSdm, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type spiSdmFile struct {
	abstraction.Repository
}

func NewSpiSdmFile(db *gorm.DB) *spiSdmFile {
	return &spiSdmFile{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiSdmFile) Create(ctx *abstraction.Context, m *model.SpiSdmFile) (*model.SpiSdmFile, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiSdmFile) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
