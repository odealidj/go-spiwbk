package repository

import (
	"code-boiler/internal/abstraction"
	"code-boiler/internal/model"
	"fmt"
	"reflect"

	"gorm.io/gorm"
)

type Sample interface {
	Find(ctx *abstraction.Context, m *model.SampleFilterModel, p *abstraction.Pagination) (*[]model.SampleEntityModel, *abstraction.PaginationInfo, error)
	FindByID(ctx *abstraction.Context, id *int) (*model.SampleEntityModel, error)
	Create(ctx *abstraction.Context, e *model.SampleEntity) (*model.SampleEntityModel, error)
	Update(ctx *abstraction.Context, id *int, e *model.SampleEntity) (*model.SampleEntityModel, error)
	Delete(ctx *abstraction.Context, id *int) (*model.SampleEntityModel, error)

	checkTrx(ctx *abstraction.Context) *gorm.DB
	filter(ctx *abstraction.Context, query *gorm.DB, payload *model.SampleFilterModel) *gorm.DB
}

type sample struct {
	abstraction.Repository
}

func NewSample(db *gorm.DB) *sample {
	return &sample{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *sample) Find(ctx *abstraction.Context, m *model.SampleFilterModel, p *abstraction.Pagination) (*[]model.SampleEntityModel, *abstraction.PaginationInfo, error) {
	conn := r.checkTrx(ctx)

	var datas []model.SampleEntityModel
	var info abstraction.PaginationInfo

	query := conn.Model(&model.SampleEntityModel{})

	// filter
	query = r.filter(ctx, query, m)

	// sort
	if p.Sort != "" && p.SortBy != "" {
		sort := fmt.Sprintf("%s %s", p.SortBy, p.Sort)
		query = query.Order(sort)
	}

	// pagination
	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	limit := p.PageSize + 1
	offset := (p.Page - 1) * limit
	query = query.Limit(limit).Offset(offset)

	err := query.Find(&datas).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &datas, &info, err
	}

	info.Count = len(datas)
	info.MoreRecords = false
	if len(datas) > p.PageSize {
		info.MoreRecords = true
		info.Count -= 1
		datas = datas[:len(datas)-1]
	}

	return &datas, &info, nil
}

func (r *sample) FindByID(ctx *abstraction.Context, id *int) (*model.SampleEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.SampleEntityModel
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) Create(ctx *abstraction.Context, e *model.SampleEntity) (*model.SampleEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.SampleEntityModel
	data.Context = ctx
	data.SampleEntity = *e
	err := conn.Create(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) Update(ctx *abstraction.Context, id *int, e *model.SampleEntity) (*model.SampleEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.SampleEntityModel
	data.Context = ctx
	data.SampleEntity = *e
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Model(&data).Updates(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *sample) Delete(ctx *abstraction.Context, id *int) (*model.SampleEntityModel, error) {
	conn := r.checkTrx(ctx)

	var data model.SampleEntityModel
	data.Context = ctx
	err := conn.Where("id = ?", id).First(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	err = conn.Where("id = ?", id).Delete(&data).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (r *sample) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}

func (r *sample) filter(ctx *abstraction.Context, query *gorm.DB, payload *model.SampleFilterModel) *gorm.DB {
	mVal := reflect.ValueOf(*payload)
	mType := reflect.TypeOf(*payload)

	for i := 0; i < mVal.NumField(); i++ {
		mValChild := mVal.Field(i)
		mTypeChild := mType.Field(i)

		for j := 0; j < mValChild.NumField(); j++ {
			val := mValChild.Field(j)

			if !val.IsNil() {
				if val.Kind() == reflect.Ptr {
					val = mValChild.Field(j).Elem()
				}

				key := mTypeChild.Type.Field(j).Tag.Get("query")
				filter := mTypeChild.Type.Field(j).Tag.Get("filter")

				switch filter {
				case "LIKE":
					query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+val.String()+"%")
				case "ILIKE":
					query = query.Where(fmt.Sprintf("%s ILIKE ?", key), "%"+val.String()+"%")
				default:
					query = query.Where(fmt.Sprintf("%s = ?", key), val.String())
				}
			}
		}
	}

	return query
}
