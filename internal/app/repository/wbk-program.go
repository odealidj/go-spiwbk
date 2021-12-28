package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"sync"
)

type WbkProgram interface {
	//Create(*abstraction.Context, *model.WbkProgramRanker) (*model.WbkProgramRanker, error)
	//Upsert(*abstraction.Context, *model.WbkProgramRanker) (*model.WbkProgramRanker, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Find(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) (*[]model.SpiAngItem, *abstraction.PaginationInfo, error)
	Upsert(*abstraction.Context, *model.WbkProgram) (*model.WbkProgram, error)
	Find(*abstraction.Context, *model.WbkProgramFilter, *abstraction.Pagination) ([]dto.WbkProgramGetResponse, *abstraction.PaginationInfo, error)
	FindByThnAngIDAndSatkerID(*abstraction.Context, *model.WbkProgramFilter, *abstraction.Pagination) ([]dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse, *abstraction.PaginationInfo, error)

	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkProgram struct {
	abstraction.Repository
}

func NewWbkProgram(db *gorm.DB) *wbkProgram {
	return &wbkProgram{
		abstraction.Repository{
			Db: db,
		},
	}
}

/*
func (r *wbkProgramRanker) Create(ctx *abstraction.Context, m *model.WbkProgramRanker) (*model.WbkProgramRanker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_program_id": m.WbkProgramID,
		"code": m.Code, "name": m.Name, "tag": m.Tag}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

*/

func (r *wbkProgram) Upsert(ctx *abstraction.Context, m *model.WbkProgram) (*model.WbkProgram, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_komponen_id", "code", "name", "tag"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkProgram) Find(ctx *abstraction.Context,
	m *model.WbkProgramFilter, p *abstraction.Pagination) ([]dto.WbkProgramGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.WbkProgramGetResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
	//	*m.ThnAngID, *m.SatkerID)

	partQuery := fmt.Sprintf("wp.deleted_at is NULL")

	query := conn.Table("wbk_program wp").
		Select(
			`wp.*, wk.name as wbk_komponen_name
	`).
		Joins(`inner join wbk_komponen wk ON wp.wbk_komponen_id = wk.id and wk.deleted_at is NULL`)

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
		sortBy := "id"
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

func (r *wbkProgram) FindByThnAngIDAndSatkerID(ctx *abstraction.Context,
	m *model.WbkProgramFilter, p *abstraction.Pagination) ([]dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
		*m.ThnAngID, *m.SatkerID)

	query := conn.Table("wbk_program wp").
		Select(
			`wp.id, wp.wbk_komponen_id, wp.code, wp.name, wp.tag, sum(wprn.nilai)/Count(wprn.nilai) as nilai
	`).
		Joins(`left outer join wbk_program_ranker wpr on wpr.wbk_program_id = wp.id and wp.deleted_at is NULL`).
		Joins(`left outer join  wbk_program_ranker_nilai wprn on wprn.wbk_program_ranker_id = wpr.id and wprn.deleted_at is NULL`).
		Joins(`left outer join thn_ang_satker tas on wprn.thn_ang_satker_id = tas.id and tas.deleted_at is NULL`).
		Joins(`left outer join satker s on tas.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`left outer join thn_ang ta on tas.thn_ang_id = ta.id and ta.deleted_at is NULL`)

	query = r.Filter(ctx, query, *m).Where(partQuery).
		Group(`wp.id, wp.wbk_komponen_id, wp.code, wp.name, wp.tag`)
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

func (r *wbkProgram) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
