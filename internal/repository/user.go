package repository

import (
	"code-boiler/internal/abstractions"
	"code-boiler/internal/model"

	"gorm.io/gorm"
)

type User interface {
	FindByID(id int) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Create(data *model.User) (*model.User, error)
	Update(id int, data *model.User) (*model.User, error)
	Delete(id int) (*model.User, error)

	WithTrx(trx *gorm.DB) *user
}

type user struct {
	abstractions.Repository
}

func NewUser(dbConnection *gorm.DB) *user {
	return &user{
		abstractions.Repository{
			DBConnection: dbConnection,
		},
	}
}

func (r *user) WithTrx(trx *gorm.DB) *user {
	new := &user{
		abstractions.Repository{
			DBConnection: trx,
		},
	}
	return new
}

func (r *user) FindByID(id int) (*model.User, error) {
	var data *model.User
	err := r.DBConnection.Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *user) FindByEmail(email string) (*model.User, error) {
	var data model.User
	err := r.DBConnection.Where("email = ?", email).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *user) Create(payload *model.User) (*model.User, error) {
	err := r.DBConnection.Create(&payload).Error
	if err != nil {
		return payload, err
	}
	err = r.DBConnection.Model(&payload).First(&payload).Error
	if err != nil {
		return payload, err
	}
	return payload, nil
}

func (r *user) Update(id int, data *model.User) (*model.User, error) {
	var user *model.User
	err := r.DBConnection.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}
	err = r.DBConnection.Model(&user).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *user) Delete(id int) (*model.User, error) {
	var data *model.User
	data, err := r.FindByID(id)
	if err != nil {
		return nil, err
	}
	err = r.DBConnection.Where("id = ?", id).Delete(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
