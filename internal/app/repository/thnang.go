package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"math"
	"sync"
)

type ThnAng interface {
	Create(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	CreateBatch(*abstraction.Context, []model.ThnAng) ([]model.ThnAng, error)
	Update(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	Delete(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	FindByID(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	FindByYear(*abstraction.Context, *model.ThnAng) (*model.ThnAng, error)
	Find(*abstraction.Context, *model.ThnAngFilter, *abstraction.Pagination) (*[]model.ThnAng, *abstraction.PaginationInfo, error)
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

func (r *thnang) FindByYear(ctx *abstraction.Context, m *model.ThnAng) (*model.ThnAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.WithContext(ctx.Request().Context()).Where("year = ?", m.Year).Find(&m).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *thnang) Find(ctx *abstraction.Context, m *model.ThnAngFilter, p *abstraction.Pagination) (*[]model.ThnAng, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.ThnAng
	var info abstraction.PaginationInfo

	query := conn.Model(&model.ThnAng{})

	//filter
	query = r.Filter(ctx, query, *m)
	queryCount := query

	ChErr := make(chan error)
	defer close(ChErr)

	group := &sync.WaitGroup{}
	group.Add(2)
	go func(group *sync.WaitGroup) {
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
		return &result, &info, err
	}

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "id"
		p.SortBy = &sortBy
	}

	p.Count = count

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	info = abstraction.PaginationInfo{
		Pagination: p,
	}

	if p.Page == nil {
		page := 0
		p.Page = &page
	}

	if p.PageSize == nil {
		pageSize := 0
		p.PageSize = &pageSize
	}
	p.Count = count

	if *p.Page < 0 {
		*p.Page = 1
	}
	if *p.Page > int(count) {
		*p.Page = int(count)
	}

	limit := *p.PageSize
	offset := (*p.Page - 1) * limit
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	err = query.Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &result, &info, err
	}

	if *p.PageSize == 0 {
		info.Pages = 0
	} else {
		info.Pages = int(math.Ceil(float64(count) / float64(*p.PageSize)))
	}
	info.MoreRecords = false

	if *p.Page+1 <= info.Pages {
		info.MoreRecords = true
	}

	return &result, &info, nil
}

func (r *thnang) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
