package repository

import (
	"code-boiler/internal/abstractions"
	"code-boiler/internal/dto"
	"code-boiler/internal/model"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Sample interface {
	Find(payload *dto.SampleGetRequest) ([]*model.Sample, *abstractions.PaginationInfo, error)
	FindByID(id int) (*model.Sample, error)
	FindByKey(key string) (*model.Sample, error)
	Create(data *model.Sample) (*model.Sample, error)
	Update(id int, data *model.Sample) (*model.Sample, error)
	Delete(id int) (*model.Sample, error)

	WithTrx(trx *gorm.DB) *sample
}

type sample struct {
	abstractions.Repository
}

func NewSample(dbConnection *gorm.DB) *sample {
	return &sample{
		abstractions.Repository{
			DBConnection: dbConnection,
		},
	}
}

func (r *sample) WithTrx(trx *gorm.DB) *sample {
	new := &sample{
		abstractions.Repository{
			DBConnection: trx,
		},
	}
	return new
}

func (r *sample) Find(payload *dto.SampleGetRequest) ([]*model.Sample, *abstractions.PaginationInfo, error) {
	var data []*model.Sample
	var info *abstractions.PaginationInfo
	query := r.DBConnection.Model(&model.Sample{})
	var sort string = "created_at desc"
	logrus.Info(sort)

	if payload.Pagination.Page != 0 {
		var limit int = 10
		if payload.Pagination.PageSize != 0 {
			limit = payload.Pagination.PageSize
		}

		var total int64
		var count int64
		query.Count(&total)

		offset := (payload.Pagination.Page - 1) * limit
		query = query.Limit(limit).Offset(offset)
		query.Count(&count)

		info = &abstractions.PaginationInfo{
			Pagination: &abstractions.Pagination{
				Page:     payload.Pagination.Page,
				PageSize: payload.Pagination.PageSize,
				Sort:     payload.Pagination.Sort,
			},
			Total: total,
			Count: count,
		}
	}
	if payload.Pagination.Sort != "" {
		sort = payload.Pagination.Sort
	}
	query = filter(&payload.Filter, query)
	query = query.Order(sort)

	err := query.Find(&data).Error
	if err != nil {
		return data, info, err
	}
	return data, info, nil
}

func (r *sample) FindByID(id int) (*model.Sample, error) {
	var data *model.Sample
	err := r.DBConnection.Where("id = ?", id).First(&data).Error
	if err != nil {
		return data, err
	}
	return data, nil
}

func (r *sample) FindByKey(key string) (*model.Sample, error) {
	var data model.Sample
	err := r.DBConnection.Where("key = ?", key).First(&data).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) Create(payload *model.Sample) (*model.Sample, error) {
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

func (r *sample) Update(id int, data *model.Sample) (*model.Sample, error) {
	var sample *model.Sample
	err := r.DBConnection.Where("id = ?", id).First(&sample).Error
	if err != nil {
		return nil, err
	}
	err = r.DBConnection.Model(&sample).Updates(data).Error
	if err != nil {
		return nil, err
	}
	return sample, nil
}

func (r *sample) Delete(id int) (*model.Sample, error) {
	var data *model.Sample
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

func filter(payload *dto.SampleGetFilterRequest, query *gorm.DB) *gorm.DB {
	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(payload)
	json.Unmarshal(inrec, &inInterface)
	for key, value := range inInterface {
		query = query.Where(fmt.Sprintf("%s = ?", key), value)
	}
	return query
}
