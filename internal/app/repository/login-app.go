package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"

	//"fmt"

	"gorm.io/gorm"
)

type LoginApp interface {
	FindLoginAppByUsername(*abstraction.Context, *model.LoginApp) (*model.LoginApp, error)
	//Create(*abstraction.Context, *model.UserApp) (*model.UserApp, error)
	CreateLoginApp(*abstraction.Context, *model.LoginApp) (*model.LoginApp, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type loginapp struct {
	abstraction.Repository
}

func NewLoginApp(db *gorm.DB) *loginapp {
	return &loginapp{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *loginapp) FindLoginAppByUsername(ctx *abstraction.Context, m *model.LoginApp) (*model.LoginApp, error) {
	conn := r.CheckTrx(ctx)

	//var data *model.LoginApp

	err := conn.Where("username = ?", m.Username).First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *loginapp) CreateLoginApp(ctx *abstraction.Context, m *model.LoginApp) (*model.LoginApp, error) {
	conn := r.CheckTrx(ctx)
	//var userapp model.UserApp
	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return m, nil

}

func (r *loginapp) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
