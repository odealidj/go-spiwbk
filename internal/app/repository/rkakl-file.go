package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
)

type RkaklFile interface {
	Create(*abstraction.Context, *model.RkaklFile) (*model.RkaklFile, error)
	Update(*abstraction.Context, *model.RkaklFile) (*model.RkaklFile, error)
	Delete(*abstraction.Context, *model.RkaklFile) (*model.RkaklFile, error)
	FindByID(*abstraction.Context, *model.RkaklFile) (*model.RkaklFile, error)
	//Find(*abstraction.Context, *model.SpiSdmFilter, *abstraction.Pagination) (*[]model.SpiSdm, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type rkaklFile struct {
	abstraction.Repository
}

func NewRkaklFile(db *gorm.DB) *rkaklFile {
	return &rkaklFile{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *rkaklFile) Create(ctx *abstraction.Context, m *model.RkaklFile) (*model.RkaklFile, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklFile) Update(ctx *abstraction.Context, m *model.RkaklFile) (*model.RkaklFile, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklFile) Delete(ctx *abstraction.Context, m *model.RkaklFile) (*model.RkaklFile, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklFile) FindByID(ctx *abstraction.Context, m *model.RkaklFile) (*model.RkaklFile, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {

		return nil, err
	}
	return m, nil
}

func (r *rkaklFile) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
