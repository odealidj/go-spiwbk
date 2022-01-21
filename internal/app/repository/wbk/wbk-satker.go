package wbk

import (
	"codeid-boiler/internal/abstraction"
	dto_wbk "codeid-boiler/internal/app/dto/wbk"
	"codeid-boiler/internal/app/model/wbk"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"math"
	"sync"
)

type WbkSatker interface {
	Create(*abstraction.Context, *wbk.WbkSatker) (*wbk.WbkSatker, error)
	Upsert(*abstraction.Context, *[]wbk.WbkSatker) (*[]wbk.WbkSatker, error)
	//Find(*abstraction.Context, *wbk.WbkSatkerFilter, *abstraction.Pagination) ([]dto_wbk.WbkSatkerGetResponse, *abstraction.PaginationInfo, error)
	FindProram(*abstraction.Context, *wbk.WbkSatkerFilter, *abstraction.Pagination) ([]dto_wbk.WbkSatkerGetProgramResponse, *abstraction.PaginationInfo, error)
	//FindByTahunIDAndSatkerID(*abstraction.Context, *wbk.WbkSatkerFilter, *abstraction.Pagination) ([]dto_wbk.WbkSatkerGetResponse, *abstraction.PaginationInfo, error)

	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkSatker struct {
	abstraction.Repository
}

func NewWbkSatker(db *gorm.DB) *wbkSatker {
	return &wbkSatker{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *wbkSatker) Create(ctx *abstraction.Context, m *wbk.WbkSatker) (*wbk.WbkSatker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_komponen_id": m.WbkKomponenID,
		"thn_ang_satker_id": m.ThnAngSatkerID}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSatker) Upsert(ctx *abstraction.Context, m *[]wbk.WbkSatker) (*[]wbk.WbkSatker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_komponen_id", "thn_ang_satker_id"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSatker) FindProram(ctx *abstraction.Context,
	m *wbk.WbkSatkerFilter, p *abstraction.Pagination) ([]dto_wbk.WbkSatkerGetProgramResponse, *abstraction.PaginationInfo, error) {

	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto_wbk.WbkSatkerGetProgramResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
	//	*m.ThnAngID, *m.SatkerID)

	partQuery := fmt.Sprintf("ws.deleted_at is NULL")

	query := conn.Table("wbk_satker ws").
		Select(
			`ws.*,wk.code as wbk_komponen_code, wk.name as wbk_komponen_name, 
			wp.id as wbk_program_id, wp.code as wbk_program_code, wp.name as wbk_program_name,
			tas.thn_ang_id, ta.year,tas.satker_id, s.name as satker_name
	`).
		Joins(`inner join wbk_komponen wk ON ws.wbk_komponen_id = wk.id and wk.deleted_at is NULL`).
		Joins(`inner join wbk_program wp ON wp.wbk_komponen_id = wk.id and wp.deleted_at is NULL`).
		Joins(`inner join thn_ang_satker tas on ws.thn_ang_satker_id = tas.id  and tas.deleted_at is NULL`).
		Joins(`inner join satker s on tas.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`inner join thn_ang ta on tas.thn_ang_id = ta.id and ta.deleted_at is NULL`)
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
		sortBy := "wp.num"
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
	if limit > 0 {
		query = query.Limit(limit).Offset(offset)
	}

	//}

	err = query.Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return result, &info, err
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

	return result, &info, nil
}

func (r *wbkSatker) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
