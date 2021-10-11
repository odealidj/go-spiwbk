package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/auth/model"


	//"fmt"

	"gorm.io/gorm"
)

type Auth interface {
	FindByUsername(*abstraction.Context, *string) (*model.LoginApp, error)
	//Create(*abstraction.Context, *model.UserApp) (*model.UserApp, error)
	Create(*abstraction.Context, *model.UserApp, *model.LoginApp) (*model.UserApp, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type auth struct {
	abstraction.Repository
}

func NewAuth(db *gorm.DB) *auth {
	return &auth{
		abstraction.Repository{
			Db: db,
		},
	}
}

/*
func (r *auth) Create1(ctx *abstraction.Context, u *model.UserAppEntity, l *model.LoginAppEntity) (*model.UserApp, error) {
	conn := r.checkTrx(ctx)

	var dataUser model.UserApp
	var dataLogin model.LoginApp

	dataUser.UserAppEntity = *u
	err := conn.Create(&dataUser).Error
	if err != nil {
		return nil, err
	}

	dataLogin.ID.ID = dataUser.ID
	dataLogin.LoginAppEntity = *l
	err = conn.Create(&dataLogin).Error
	if err != nil {
		return nil, err
	}

	return &dataUser, nil
}
*/

/*
func (r *auth) Create(ctx *abstraction.Context, u *model.UserAppEntity , l *model.LoginAppEntity) (*model.UserApp, error) {
	conn := r.CheckTrx(ctx)

	var dataUser model.UserApp
	var dataLogin model.LoginApp

	dataUser.UserAppEntity = *u
	err := conn.Create(&dataUser).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	dataLogin.ID.ID = dataUser.ID
	dataLogin.LoginAppEntity = *l
	err = conn.Create(&dataLogin).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return  &dataUser, nil
}

*/


func (r *auth) FindByUsername(ctx *abstraction.Context, username *string) (*model.LoginApp, error) {
	conn := r.CheckTrx(ctx)

	var data *model.LoginApp

	//var loginapp *model.LoginApp
	//var ID uint16
	//var RoleUserId int
	//var UserName string
	//var Passwordhash string
	/*
	   "role_user_id": 1,
	   "satker_id": 1,
	   "jabatan_id": 1,
	   "nip": "nip93",
	   "name": "name",
	   "address": "address",
	   "postal_code": "postal_code",
	   "phone_number": "phone_number",
	   "mobile_number": "mobile_number",
	   "email": "odealidj@gmail.com" */


	   /*
	type Test struct {
		ID int
		Nip string
		Name string
	}
	*/

	//var ts Test

	//err := conn.Raw("SELECT id, nip, name FROM user_app WHERE id = ?", 7).Find(&ts).Error
	//err := row.Scan(&Test).Error

	err := conn.Where("username = ?", username).Preload("UserApp").Limit(1).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (r *auth) Create(ctx *abstraction.Context, u *model.UserApp, l *model.LoginApp) (*model.UserApp, error) {
	conn := r.CheckTrx(ctx)
	//var userapp *model.UserApp

	err := conn.Create(&u).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	login := *l
	login.UserAppId = u.ID
	err = conn.Create(&login).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return u, nil

}

func (r *auth) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
