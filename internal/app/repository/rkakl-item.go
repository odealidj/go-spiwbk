package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RkaklItem interface {
	Create(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	Update(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	Delete(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	FindByID(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	Upsert(*abstraction.Context, *model.RkaklItem) (*model.RkaklItem, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type rkaklItem struct {
	abstraction.Repository
}

func NewRkaklItem(db *gorm.DB) *rkaklItem {
	return &rkaklItem{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *rkaklItem) Create(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) Update(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) Delete(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) FindByID(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) Upsert(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sub_komponen_akun_id", "name", "volume", "harga", "biaya"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) FirstOrCreate(ctx *abstraction.Context, m *model.RkaklItem) (*model.RkaklItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"sub_komponen_akun_id": m.SubKomponenAkunID,
		"name": m.Name, "volume": m.Volume, "harga": m.Harga, "biaya": m.Biaya}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklItem) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
