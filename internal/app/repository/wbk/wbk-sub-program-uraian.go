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

type WbkSubProgramUraian interface {
	Create(*abstraction.Context, *wbk.WbkSubProgramUraian) (*wbk.WbkSubProgramUraian, error)
	Upsert(*abstraction.Context, *wbk.WbkSubProgramUraian) (*wbk.WbkSubProgramUraian, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	Find(*abstraction.Context, *wbk.WbkSubProgramUraianFilter, *abstraction.PaginationArr) ([]wbk2.WbkSubProgramUraianGetResponse, *abstraction.PaginationInfoArr, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkSubProgramUraian struct {
	abstraction.Repository
}

func NewWbkSubProgramUraian(db *gorm.DB) *wbkSubProgramUraian {
	return &wbkSubProgramUraian{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *wbkSubProgramUraian) Create(ctx *abstraction.Context, m *wbk.WbkSubProgramUraian) (*wbk.WbkSubProgramUraian, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_sub_program_ranker_id": m.WbkSubProgramRankerID,
		"frekuensi_ranker_id": m.FrekuensiRankerID, "code": m.Code, "name": m.Name, "ket": m.Ket}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSubProgramUraian) Upsert(ctx *abstraction.Context, m *wbk.WbkSubProgramUraian) (*wbk.WbkSubProgramUraian, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_sub_program_ranker_id",
			"frekuensi_ranker_id", "code", "name", "ket"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkSubProgramUraian) Find(ctx *abstraction.Context,
	m *wbk.WbkSubProgramUraianFilter, p *abstraction.PaginationArr) ([]wbk2.WbkSubProgramUraianGetResponse,
	*abstraction.PaginationInfoArr, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []wbk2.WbkSubProgramUraianGetResponse
	var info abstraction.PaginationInfoArr

	partJoin := ""
	if m.WbkSubProgramRankerID != nil {
		partJoin = fmt.Sprintf(`inner join wbk_sub_program_uraian wspu on wpr.id = wspu.wbk_sub_program_ranker_id 
	and  wspu.wbk_sub_program_ranker_id =%d and wspu.deleted_at is NULL`, *m.WbkSubProgramRankerID)
	} else {
		partJoin = fmt.Sprintf(`inner join wbk_sub_program_uraian wspu on wpr.id = wspu.wbk_sub_program_ranker_id 
	and wspu.deleted_at is NULL`)
	}

	query := conn.Table("wbk_sub_program_ranker wspr").
		Select(
			`wspu.id, wspu.wbk_sub_program_ranker_id, wspu.frekuensi_ranker_id,
		   fr.name as frekuensi_waktu, 
		   wspu.code, wspu.name, 
		   Sum(case when wspub.bulan_id = 1 then 1 else 0 end) 'b1',
		   Sum(case when wspub.bulan_id = 2 then 1 else 0 end) 'b2',
		   Sum(case when wspub.bulan_id = 3 then 1 else 0 end) 'b3',
		   Sum(case when wspub.bulan_id = 4 then 1 else 0 end) 'b4',
		   Sum(case when wspub.bulan_id = 5 then 1 else 0 end) 'b5',
		   Sum(case when wspub.bulan_id = 6 then 1 else 0 end) 'b6',
		   Sum(case when wspub.bulan_id = 7 then 1 else 0 end) 'b7',
		   Sum(case when wspub.bulan_id = 8 then 1 else 0 end) 'b8',
		   Sum(case when wspub.bulan_id = 9 then 1 else 0 end) 'b9',
		   Sum(case when wspub.bulan_id = 10 then 1 else 0 end) 'b10',
		   Sum(case when wspub.bulan_id = 11 then 1 else 0 end) 'b11',
		   Sum(case when wspub.bulan_id = 12 then 1 else 0 end) 'b12',
		   wspu.ket,
		   CONCAT(wk.code, ". ", wk.name) as komponen, 
		   CONCAT(wp.code,". ", wp.name) as program,
		   CONCAT(wpr.code, ". ",wpr.name) as program_renja,
		   CONCAT(wspr.code, ". ", wspr.name) as sub_program_renja,
		   wspr.id as  wbk_sub_program_ranker_id`).
		Joins(`inner join wbk_program_ranker wpr on wspr.wbk_program_ranker_id = wpr.id and wpr.deleted_at is NULL`).
		Joins(`inner join wbk_program wp on wpr.wbk_program_id = wp.id and wp.deleted_at is NULL`).
		Joins(`inner join wbk_komponen wk on wp.wbk_komponen_id = wk.id and wk.deleted_at is NULL`).
		Joins(partJoin).
		Joins(`inner join frekuensi_ranker fr on wspu.frekuensi_ranker_id = fr.id and fr.deleted_at is NULL`).
		Joins(`inner join wbk_sub_program_uraian_bulan wspub on wspub.wbk_sub_program_uraian_id = wspu.id 
	and wspub.deleted_at is NULL`).
		Joins(`inner JOIN bulan b on wspub.bulan_id = b.id and b.deleted_at is NULL`)

	query = r.Filter(ctx, query, *m).Where("wspu.deleted_at is NULL").
		Group(`wspu.id, wspu.wbk_sub_program_ranker_id, wspu.frekuensi_ranker_id,
				fr.name, wspu.code, wspu.name, wspu.ket,  CONCAT(wk.code, ". ", wk.name),
				CONCAT(wp.code,". ", wp.name), CONCAT(wpr.code, ". ",wpr.name), CONCAT(wspr.code, ". ", wspr.name),
				wspr.id`)
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

	//sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	info = abstraction.PaginationInfoArr{
		Count:         int(count),
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

func (r *wbkSubProgramUraian) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
