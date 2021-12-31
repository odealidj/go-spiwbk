package wbk

import (
	"codeid-boiler/internal/abstraction"
	wbk2 "codeid-boiler/internal/app/dto/wbk"
	"codeid-boiler/internal/app/model/wbk"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"sync"
)

type WbkProgramRanker interface {
	Create(*abstraction.Context, *wbk.WbkProgramRanker) (*wbk.WbkProgramRanker, error)
	Upsert(*abstraction.Context, *wbk.WbkProgramRanker) (*wbk.WbkProgramRanker, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	Find(*abstraction.Context, *wbk.WbkProgramRankerFilter, *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetResponse, *abstraction.PaginationInfo, error)
	FindSatkerNilaiByThnAngID(*abstraction.Context, *wbk.WbkProgramRankerFilter, *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetSatkerNilaiResponse, *abstraction.PaginationInfo, error)
	FindByThnAngIDAndSatkerID(*abstraction.Context, *wbk.WbkProgramRankerFilter, *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetResponse, *abstraction.PaginationInfo, error)

	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkProgramRanker struct {
	abstraction.Repository
}

func NewWbkProgramRanker(db *gorm.DB) *wbkProgramRanker {
	return &wbkProgramRanker{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *wbkProgramRanker) Create(ctx *abstraction.Context, m *wbk.WbkProgramRanker) (*wbk.WbkProgramRanker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_program_id": m.WbkProgramID,
		"code": m.Code, "name": m.Name, "tag": m.Tag}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkProgramRanker) Upsert(ctx *abstraction.Context, m *wbk.WbkProgramRanker) (*wbk.WbkProgramRanker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_program_id", "code", "name", "tag"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkProgramRanker) Find(ctx *abstraction.Context,
	m *wbk.WbkProgramRankerFilter, p *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []wbk2.WbkProgramRankerGetResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
	//*m.ThnAngID, *m.SatkerID)

	query := conn.Table("wbk_program wp").
		Select(
			`wp.id as wbk_program_id, wp.code as wbk_program_code,  wp.name as wbk_program_name,
					wpr.*
	`).
		Joins(`inner join wbk_program_ranker wpr ON wpr.wbk_program_id = wp.id and wpr.deleted_at is NULL`)

	query = r.Filter(ctx, query, *m).Where("wp.deleted_at is NULL")
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
		sortBy := "wp.id, wpr.id"
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

func (r *wbkProgramRanker) FindSatkerNilaiByThnAngID(ctx *abstraction.Context,
	m *wbk.WbkProgramRankerFilter, p *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetSatkerNilaiResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []wbk2.WbkProgramRankerGetSatkerNilaiResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("s.deleted_at is NULL")

	partQueryTas := fmt.Sprintf("left outer join thn_ang_satker tas on tas.satker_id = s.id and tas.thn_ang_id = %d and s.deleted_at is NULL", *m.ThnAngID)

	query := conn.Table("satker s").
		Select(
			`s.id, s.unit_id, s.kabkota_id, s.code, s.name, 
					IFNULL(sum(wprn.nilai)/COUNT(*),0) as nilai
		`).
		Joins(partQueryTas).
		Joins(`left outer join thn_ang ta on tas.thn_ang_id = ta.id and ta.deleted_at is NULL`).
		Joins(`LEFT outer join wbk_program_ranker_nilai wprn on wprn.thn_ang_satker_id = tas.id and wprn.deleted_at is NULL`).
		Joins(`left outer join wbk_program_ranker wpr on wprn.wbk_program_ranker_id = wpr.id and wpr.deleted_at is NULL`)

	query = r.Filter(ctx, query, *m).Where(partQuery).
		Group("s.unit_id, s.kabkota_id, s.code, s.name")
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

func (r *wbkProgramRanker) FindByThnAngIDAndSatkerID(ctx *abstraction.Context,
	m *wbk.WbkProgramRankerFilter, p *abstraction.Pagination) ([]wbk2.WbkProgramRankerGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []wbk2.WbkProgramRankerGetResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
		*m.ThnAngID, *m.SatkerID)

	query := conn.Table("wbk_program wp").
		Select(
			`wp.id as wbk_program_id, wp.code as wbk_program_code, wp.name as wbk_program_name,
					wpr.*, wprn.nilai
	`).
		Joins(`inner join wbk_program_ranker wpr on wpr.wbk_program_id = wp.id and wp.deleted_at is NULL`).
		Joins(`left outer join  wbk_program_ranker_nilai wprn on wprn.wbk_program_ranker_id = wpr.id and wprn.deleted_at is NULL`).
		Joins(`left outer join thn_ang_satker tas on wprn.thn_ang_satker_id = tas.id and tas.deleted_at is NULL`).
		Joins(`left outer join satker s on tas.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`left outer join thn_ang ta on tas.thn_ang_id = ta.id and ta.deleted_at is NULL`)

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

func (r *wbkProgramRanker) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
