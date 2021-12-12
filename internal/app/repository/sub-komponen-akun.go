package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubKomponenAkun interface {
	Create(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Update(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Delete(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	FindByID(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Upsert(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type subKomponenAkun struct {
	abstraction.Repository
}

func NewSubKomponenAkun(db *gorm.DB) *subKomponenAkun {
	return &subKomponenAkun{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *subKomponenAkun) Create(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Update(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Delete(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) FindByID(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Upsert(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sub_komponen_id", "akun_id", "name", "biaya", "sdcp"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) FirstOrCreate(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"sub_komponen_id": m.SubKomponenID, "akun_id": m.AkunID,
		"name": m.Name, "biaya": m.Biaya, "sdcp": m.Sdcp}).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
