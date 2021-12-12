package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubKomponen interface {
	Create(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	Update(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	Delete(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	FindByID(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	Upsert(*abstraction.Context, *model.SubKomponen) (*model.SubKomponen, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type subKomponen struct {
	abstraction.Repository
}

func NewSubKomponen(db *gorm.DB) *subKomponen {
	return &subKomponen{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *subKomponen) Create(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) Update(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) Delete(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) FindByID(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) Upsert(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"komponen_id", "code", "name", "biaya"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) FirstOrCreate(ctx *abstraction.Context, m *model.SubKomponen) (*model.SubKomponen, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"komponen_id": m.KomponenID, "code": m.Code, "name": m.Name, "biaya": m.Biaya}).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponen) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
