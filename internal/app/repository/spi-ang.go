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

type SpiAng interface {
	Create(*abstraction.Context, *model.SpiAng) (*model.SpiAng, error)
	Update(*abstraction.Context, *model.SpiAng) (*model.SpiAng, error)
	Delete(*abstraction.Context, *model.SpiAng) (*model.SpiAng, error)
	FindByID(*abstraction.Context, *model.SpiAng) (*model.SpiAng, error)
	Find(*abstraction.Context, *model.SpiAngFilter, *abstraction.Pagination) (*[]model.SpiAng, *abstraction.PaginationInfo, error)
	FindByThnAngIDAndSatkerIDExistSpiAngItem(*abstraction.Context,
		*model.SpiAngFilter, *abstraction.Pagination) ([]dto.SpiAngResponse, *abstraction.PaginationInfo, error)

	checkTrx(*abstraction.Context) *gorm.DB
}
type spiang struct {
	abstraction.Repository
}

func NewSpiAng(db *gorm.DB) *spiang {
	return &spiang{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiang) Create(ctx *abstraction.Context, m *model.SpiAng) (*model.SpiAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"thn_ang_id": m.ThnAngID, "satker_id": m.SatkerID}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil

}

func (r *spiang) Update(ctx *abstraction.Context, m *model.SpiAng) (*model.SpiAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiang) Delete(ctx *abstraction.Context, m *model.SpiAng) (*model.SpiAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiang) FindByID(ctx *abstraction.Context, m *model.SpiAng) (*model.SpiAng, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", m.ID).First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiang) Find(ctx *abstraction.Context, m *model.SpiAngFilter, p *abstraction.Pagination) (*[]model.SpiAng, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.SpiAng
	var info abstraction.PaginationInfo

	/*
		sql := `SELECT  ss.*, ta.year as year, s.name FROM spi_sdm ss
				INNER JOIN thn_ang ta ON ss.thn_agn_id = ta.id
				INNER JOIN satker s on ss.satker_id = s.id
				WHERE ss.deleted_at IS NULL`
	*/
	//query := conn.Model(&model.SpiSdm{})
	//query := conn.Raw(sql)
	//query := conn.Preload("ThnAng").Preload("Satker").Find(&model.SpiSdm{})
	query := conn.Joins("ThnAng").Joins("Satker").Find(&model.SpiAng{})

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
		return &result, &info, err
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

	return &result, &info, nil

}

func (r *spiang) FindByThnAngIDAndSatkerIDExistSpiAngItem(ctx *abstraction.Context,
	m *model.SpiAngFilter, p *abstraction.Pagination) ([]dto.SpiAngResponse, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.SpiAngResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("sa.deleted_at is NULL " +
		"and EXISTS (SELECT sai.spi_ang_id from spi_ang_item sai WHERE sai.deleted_at is NULL)")

	query := conn.Table("spi_ang sa").
		Select(
			`sa.*, s.name as satker_name, ta.year
		`).
		Joins(`inner join satker s on sa.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`INNER join thn_ang ta on sa.thn_ang_id = ta.id and ta.deleted_at is NULL`).
		Find(&result)

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
		sort := "desc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "sa.id"
		p.SortBy = &sortBy
	}

	//p.Count = count

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

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

func (r *spiang) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
