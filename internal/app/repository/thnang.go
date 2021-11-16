package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
)

type ThnAng interface {
	Create(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	CreateBatch(*abstraction.Context, []model.ThnAng) ([]model.ThnAng, error)
	Update(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	Delete(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	FindByID(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	Find(*abstraction.Context) ([]model.ThnAng, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type thnang struct {
	abstraction.Repository
}

func NewThnAng(db *gorm.DB) *thnang {
	return &thnang{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *thnang) Create(ctx *abstraction.Context, m *model.ThnAng) (*model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) CreateBatch(ctx *abstraction.Context, m []model.ThnAng) ([]model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) Update(ctx *abstraction.Context, m *model.ThnAng) (*model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) Delete(ctx *abstraction.Context, m *model.ThnAng) (*model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) FindByID(ctx *abstraction.Context, m *model.ThnAng) (*model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.EntityInc.IDInc.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) Find(ctx *abstraction.Context) ([]model.ThnAng, error) {
	var result []model.ThnAng
	conn := r.CheckTrx(ctx)

	err := conn.Order("id desc").Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *thnang) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
