package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"math"
	"sync"
)

type Satker interface {
	Create(*abstraction.Context, *model.Satker) (*model.Satker, error)
	Update(*abstraction.Context, *model.Satker) (*model.Satker, error)
	Delete(*abstraction.Context, *model.Satker) (*model.Satker, error)
	FindByID(*abstraction.Context, *model.Satker) (*model.Satker, error)
	Find(*abstraction.Context, model.SatkerFilter, *abstraction.Pagination) (*[]model.Satker, *abstraction.PaginationInfo, error)
	Find2(*abstraction.Context, model.SatkerFilter, *abstraction.PaginationArr) (*[]model.Satker, *abstraction.PaginationInfoArr, error)
	Count(*abstraction.Context) (*int64, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type satker struct {
	abstraction.Repository
}

func NewSatker(db *gorm.DB) *satker {
	return &satker{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *satker) Create(ctx *abstraction.Context, m *model.Satker) (*model.Satker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *satker) Update(ctx *abstraction.Context, m *model.Satker) (*model.Satker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *satker) Delete(ctx *abstraction.Context, m *model.Satker) (*model.Satker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *satker) FindByID(ctx *abstraction.Context, m *model.Satker) (*model.Satker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", m.ID).First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *satker) Find(ctx *abstraction.Context, m model.SatkerFilter, p *abstraction.Pagination) (*[]model.Satker, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var count int64
	var result []model.Satker
	var info abstraction.PaginationInfo

	query := conn.Model(&model.Satker{})

	fmt.Println(5)

	//filter
	query = r.Filter(ctx, query, m)
	fmt.Println(6)

	queryCount := query
	err := queryCount.Count(&count).WithContext(ctx.Request().Context()).Error
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

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	info = abstraction.PaginationInfo{
		Pagination: p,
	}
	//limit := *p.PageSize + 1

	fmt.Println(7)

	if p.Page != nil {
		fmt.Println(8)
		if *p.Page < 1 {
			fmt.Println("01")
			*p.Page = 1
		}
		if *p.Page > int(count) {
			fmt.Println("02")
			*p.Page = int(count)
		}

		limit := *p.PageSize
		offset := (*p.Page - 1) * limit
		//offset := (*p.Page * *p.PageSize) - *p.PageSize
		if limit > 0 {
			query = query.Limit(limit).Offset(offset)
		}

	}

	err = query.Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return &result, &info, err
	}

	if p.Page != nil {

		//info.Count = len(result)
		info.Pages = int(count) / *p.PageSize
		info.MoreRecords = false
		//if len(result) > *p.PageSize {
		if *p.Page <= *p.PageSize {
			info.MoreRecords = true
			//info.Count -= 1
			//result = result[:len(result)-1]
		}
	}

	return &result, &info, nil

}

func (r *satker) Find2(ctx *abstraction.Context, m model.SatkerFilter, p *abstraction.PaginationArr) (*[]model.Satker, *abstraction.PaginationInfoArr, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.Satker
	var info abstraction.PaginationInfoArr

	query := conn.Model(&model.Satker{})

	//filter
	query = r.Filter(ctx, query, m)

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
		return &result, &info, err
	}

	//sort
	sort := ""
	pSort := ""
	for i, sortBy := range p.SortBy {
		if len(p.Sort) == 0 {
			p.Sort = []string{"desc"}
		} else {
			if i+1 > len(p.Sort) {
				p.Sort = append(p.Sort, "desc")
			}
		}
		pSort = p.Sort[i]

		if sort == "" {
			sort = fmt.Sprintf("%s %s", sortBy, pSort)
		} else {
			sort = sort + fmt.Sprintf(", %s %s", sortBy, pSort)
		}
	}
	query = query.Order(sort)

	info = abstraction.PaginationInfoArr{
		PaginationArr: p,
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

	if err = query.Find(&result).WithContext(ctx.Request().Context()).Error; err != nil {
		return &result, &info, err
	}

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

	return &result, &info, nil

}

func (r *satker) Count(ctx *abstraction.Context) (*int64, error) {
	conn := r.CheckTrx(ctx)

	var count int64
	err := conn.Model(&model.Satker{}).WithContext(ctx.Request().Context()).Count(&count).Error
	if err != nil {
		return nil, err
	}
	return &count, nil
}

func (r *satker) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
