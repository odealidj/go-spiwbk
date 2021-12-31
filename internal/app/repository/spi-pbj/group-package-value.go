package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
	"fmt"
	"gorm.io/gorm"
	"math"
	"sync"
)

type GroupPackageValue interface {
	//Create(*abstraction.Context, *model.JenisKesesuaian) (*model.JenisKesesuaian, error)
	//Update(*abstraction.Context, *model.Akun) (*model.Akun, error)
	//Delete(*abstraction.Context, *model.Akun) (*model.Akun, error)
	//FindByID(*abstraction.Context, *model.Akun) (*model.Akun, error)
	Find(*abstraction.Context, *spi_pbj.GroupPackageValueFilter, *abstraction.Pagination) ([]spi_pbj.GroupPackageValue, *abstraction.PaginationInfo, error)
	FindExistSpiPbjPaket(*abstraction.Context, *spi_pbj.GroupPackageValueFilter, *abstraction.Pagination) ([]spi_pbj.GroupPackageValue, *abstraction.PaginationInfo, error)
	//FirstOrCreate(*abstraction.Context, *model.Akun) (*model.Akun, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type groupPackageValue struct {
	abstraction.Repository
}

func NewGroupPackageValue(db *gorm.DB) *groupPackageValue {
	return &groupPackageValue{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *groupPackageValue) Find(ctx *abstraction.Context, m *spi_pbj.GroupPackageValueFilter, p *abstraction.Pagination) ([]spi_pbj.GroupPackageValue, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []spi_pbj.GroupPackageValue
	var info abstraction.PaginationInfo

	query := conn.Model(&result)
	//Table("jenis_rekapitulasi")

	//filter

	query = r.Filter(ctx, query, *m)
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
		sort := "asc"
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

func (r *groupPackageValue) FindExistSpiPbjPaket(ctx *abstraction.Context, m *spi_pbj.GroupPackageValueFilter, p *abstraction.Pagination) ([]spi_pbj.GroupPackageValue, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []spi_pbj.GroupPackageValue
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("sa.thn_ang_id =%d and sa.satker_id =%d and ta.deleted_at is NULL",
		*m.ThnAngID, *m.SatkerID)

	query := conn.Table("thn_ang ta").
		Select(
			`DISTINCT group_package_value.* 
		`).
		Joins(`inner join spi_ang sa ON sa.thn_ang_id = ta.id and sa.deleted_at is NULL`).
		Joins(`inner join satker s on sa.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`inner join spi_pbj_paket spp on spp.spi_ang_id = sa.id and spp.deleted_at is NULL`).
		Joins(`inner join group_package_value on spp.group_package_value_id = group_package_value.id `)

	//Find(&result)

	query = r.Filter(ctx, query, *m).Where(partQuery)

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
		sort := "asc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "sa.id, sa.thn_ang_id ,sa.satker_id , spp.group_package_value_id"
		p.SortBy = &sortBy
	}

	//p.Count = count

	//sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	//query = query.Order(sort)

	info = abstraction.PaginationInfo{
		Count:      int(count),
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
	//p.Count = count

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

func (r *groupPackageValue) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
