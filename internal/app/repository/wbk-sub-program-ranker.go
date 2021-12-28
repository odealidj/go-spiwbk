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

type WbkSubProgramRanker interface {
	Create(*abstraction.Context, *model.WbkSubProgramRanker) (*model.WbkSubProgramRanker, error)
	Upsert(*abstraction.Context, *model.WbkSubProgramRanker) (*model.WbkSubProgramRanker, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	Find(*abstraction.Context, *model.WbkSubProgramRankerFilter, *abstraction.Pagination) ([]dto.WbkSubProgramRankerGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkSubProgramRanker struct {
	abstraction.Repository
}

func NewWbkSubProgramRanker(db *gorm.DB) *wbkSubProgramRanker {
	return &wbkSubProgramRanker{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *wbkSubProgramRanker) Create(ctx *abstraction.Context, m *model.WbkSubProgramRanker) (*model.WbkSubProgramRanker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_program_ranker_id": m.WbkSubProgramRankerEntity,
		"code": m.Code, "name": m.Name}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSubProgramRanker) Upsert(ctx *abstraction.Context, m *model.WbkSubProgramRanker) (*model.WbkSubProgramRanker, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_program_ranker_id", "code", "name"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSubProgramRanker) Find(ctx *abstraction.Context,
	m *model.WbkSubProgramRankerFilter, p *abstraction.Pagination) ([]dto.WbkSubProgramRankerGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.WbkSubProgramRankerGetResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
	//*m.ThnAngID, *m.SatkerID)

	query := conn.Table("wbk_sub_program_ranker wspr").
		Select(
			`wspr.id, wspr.wbk_program_ranker_id, wspr.code , wspr.name
	`)

	query = r.Filter(ctx, query, *m).Where("wspr.deleted_at is NULL")
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
		sortBy := "wspr.id"
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

func (r *wbkSubProgramRanker) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
