package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"math"
	"sync"
)

type Kegiatan interface {
	Create(*abstraction.Context, *model.Kegiatan) (*model.Kegiatan, error)
	Update(*abstraction.Context, *model.Kegiatan) (*model.Kegiatan, error)
	Delete(*abstraction.Context, *model.Kegiatan) (*model.Kegiatan, error)
	FindByID(*abstraction.Context, *model.Kegiatan) (*model.Kegiatan, error)
	Find(*abstraction.Context, *model.KegiatanFilter, *abstraction.Pagination) ([]model.Kegiatan, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.Kegiatan) (*model.Kegiatan, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type kegiatan struct {
	abstraction.Repository
}

func NewKegiatan(db *gorm.DB) *kegiatan {
	return &kegiatan{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *kegiatan) Create(ctx *abstraction.Context, m *model.Kegiatan) (*model.Kegiatan, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *kegiatan) Update(ctx *abstraction.Context, m *model.Kegiatan) (*model.Kegiatan, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *kegiatan) Delete(ctx *abstraction.Context, m *model.Kegiatan) (*model.Kegiatan, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *kegiatan) FindByID(ctx *abstraction.Context, m *model.Kegiatan) (*model.Kegiatan, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *kegiatan) Find(ctx *abstraction.Context, m *model.KegiatanFilter, p *abstraction.Pagination) ([]model.Kegiatan, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.Kegiatan
	var info abstraction.PaginationInfo

	query := conn.
		Select("*").Find(&result)

	//filter

	query = r.Filter(ctx, query, *m).Where("deleted_at is NULL")
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
		sortBy := "id"
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

func (r *kegiatan) FirstOrCreate(ctx *abstraction.Context, m *model.Kegiatan) (*model.Kegiatan, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"code": m.Code, "name": m.Name}).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *kegiatan) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
