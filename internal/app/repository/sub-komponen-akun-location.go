package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubKomponenAkunLocation interface {
	Create(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	Update(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	Delete(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	FindByID(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	Upsert(*abstraction.Context, *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type subKomponenAkunLocation struct {
	abstraction.Repository
}

func NewSubKomponenAkunLocation(db *gorm.DB) *subKomponenAkunLocation {
	return &subKomponenAkunLocation{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *subKomponenAkunLocation) Create(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) Update(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) Delete(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) FindByID(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) Upsert(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sub_komponen_akun_id", "name"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) FirstOrCreate(ctx *abstraction.Context, m *model.SubKomponenAkunLocation) (*model.SubKomponenAkunLocation, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"sub_komponen_akun_id": m.SubKomponenAkunID, "name": m.Name}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkunLocation) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
