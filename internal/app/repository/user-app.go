package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"

	//"fmt"

	"gorm.io/gorm"
)

type UserApp interface {
	Find(*abstraction.Context, *model.UserApp) (*model.UserApp, error)
	//Create(*abstraction.Context, *model.UserApp) (*model.UserApp, error)
	CreateUserApp(*abstraction.Context, *model.UserApp) (*model.UserApp, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type userapp struct {
	abstraction.Repository
}

func NewUserApp(db *gorm.DB) *userapp {
	return &userapp{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *userapp) Find(ctx *abstraction.Context, m *model.UserApp) (*model.UserApp, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *userapp) CreateUserApp(ctx *abstraction.Context, m *model.UserApp) (*model.UserApp, error) {
	conn := r.CheckTrx(ctx)
	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return m, nil

}

func (r *userapp) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
