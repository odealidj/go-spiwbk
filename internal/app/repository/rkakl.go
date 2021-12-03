package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"math"
	"sync"
)

type Rkakl interface {
	Create(*abstraction.Context, *model.Rkakl) (*model.Rkakl, error)
	Update(*abstraction.Context, *model.Rkakl) (*model.Rkakl, error)
	Delete(*abstraction.Context, *model.Rkakl) (*model.Rkakl, error)
	FindByID(*abstraction.Context, *model.Rkakl) (*model.Rkakl, error)
	FindThnAngSatkerByID(*abstraction.Context, *model.Rkakl) (*model.Rkakl, error)
	Find(*abstraction.Context, *model.RkaklFilter, *abstraction.Pagination) ([]dto.RkaklResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type rkakl struct {
	abstraction.Repository
}

func NewRkakl(db *gorm.DB) *rkakl {
	return &rkakl{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *rkakl) Create(ctx *abstraction.Context, m *model.Rkakl) (*model.Rkakl, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkakl) Update(ctx *abstraction.Context, m *model.Rkakl) (*model.Rkakl, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkakl) Delete(ctx *abstraction.Context, m *model.Rkakl) (*model.Rkakl, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkakl) FindByID(ctx *abstraction.Context, m *model.Rkakl) (*model.Rkakl, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkakl) FindThnAngSatkerByID(ctx *abstraction.Context, m *model.Rkakl) (*model.Rkakl, error) {
	conn := r.CheckTrx(ctx)

	//var result *dto.RkaklResponse
	//auto find by ID & delete_at is null
	err := conn.Joins("Satker").Joins("ThnAng").First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkakl) Find(ctx *abstraction.Context, m *model.RkaklFilter, p *abstraction.Pagination) ([]dto.RkaklResponse, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.RkaklResponse
	var info abstraction.PaginationInfo

	query := conn.Table("rkakl").
		Select(`
				rkakl.*, ta.year as thn_ang_year, s.name as satker_name, rf.filepath
			`).
		Joins("inner join thn_ang ta ON rkakl.thn_ang_id = ta.id").
		Joins("INNER JOIN satker s on rkakl.satker_id = s.id").
		Joins("INNER JOIN rkakl_file rf on rf.id = rkakl.id").
		Find(&result)

	//filter

	query = r.Filter(ctx, query, *m).Where("rkakl.deleted_at is NULL")
	queryCount := query

	ChErr := make(chan error)
	defer close(ChErr)

	group := &sync.WaitGroup{}
	group.Add(2)
	go func(group *sync.WaitGroup) {
		//var rowCount int64
		defer group.Done()

		if err := queryCount.Count(&count).WithContext(ctx.Request().Context()).Error; err != nil {
			ChErr <- err
		} else {
			ChErr <- nil
		}
	}(group)
	go func(group *sync.WaitGroup) {
		defer group.Done()

		counter := 0
		for {
			select {
			case err = <-ChErr:
				counter++
			}
			if counter == 1 {
				break
			}
		}
	}(group)
	group.Wait()

	if err != nil {
		return result, &info, err
	}

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "rkakl.id"
		p.SortBy = &sortBy
	}

	p.Count = count

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	//limit := *p.PageSize + 1

	if p.Page == nil {
		page := 0
		p.Page = &page
	}

	if p.PageSize == nil {
		pageSize := 0
		p.PageSize = &pageSize
	}
	p.Count = count

	//if p.Page != nil && p.PageSize != nil {
	if *p.Page < 0 {
		*p.Page = 1
	}
	if *p.Page > int(count) {
		*p.Page = int(count)
	}

	limit := *p.PageSize
	offset := (*p.Page - 1) * limit
	//offset := (*p.Page * *p.PageSize) - *p.PageSize
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	//}

	err = query.Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return result, &info, err
	}

	//if p.Page != nil && p.PageSize != nil {

	//info.Count = len(result)
	if *p.PageSize == 0 {
		info.Pages = 0
	} else {
		info.Pages = int(math.Ceil(float64(count) / float64(*p.PageSize)))
	}
	info.MoreRecords = false
	//if len(result) > *p.PageSize {
	if *p.Page+1 <= info.Pages {
		info.MoreRecords = true
		//info.Count -= 1
		//result = result[:len(result)-1]
	}
	//}

	return result, &info, nil

}

func (r *rkakl) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
