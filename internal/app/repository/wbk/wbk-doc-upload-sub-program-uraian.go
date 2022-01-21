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

type WbkDocUploadSubProgramUraian interface {
	Create(*abstraction.Context, *wbk.WbkDocUploadSubProgramUraian) (*wbk.WbkDocUploadSubProgramUraian, error)
	Upsert(*abstraction.Context, *wbk.WbkDocUploadSubProgramUraian) (*wbk.WbkDocUploadSubProgramUraian, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	Find(*abstraction.Context, *wbk.WbkDocUploadSubProgramUraianFilter, *abstraction.Pagination) ([]wbk2.WbkDocUploadSubProgramUraianGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type wbkDocUploadSubProgramUraian struct {
	abstraction.Repository
}

func NewWbkDocUploadSubProgramUraian(db *gorm.DB) *wbkDocUploadSubProgramUraian {
	return &wbkDocUploadSubProgramUraian{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *wbkDocUploadSubProgramUraian) Create(ctx *abstraction.Context, m *wbk.WbkDocUploadSubProgramUraian) (*wbk.WbkDocUploadSubProgramUraian, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"wbk_sub_program_uraian_bulan_id": m.WbkSubProgramUraianBulanID,
		"path": m.Path, "ket": m.Ket}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkDocUploadSubProgramUraian) Upsert(ctx *abstraction.Context, m *wbk.WbkDocUploadSubProgramUraian) (*wbk.WbkDocUploadSubProgramUraian, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"wbk_sub_program_uraian_bulan_id",
			"path", "ket"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *wbkDocUploadSubProgramUraian) Find(ctx *abstraction.Context,
	m *wbk.WbkDocUploadSubProgramUraianFilter, p *abstraction.Pagination) ([]wbk2.WbkDocUploadSubProgramUraianGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []wbk2.WbkDocUploadSubProgramUraianGetResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("tas.thn_ang_id = %d and tas.satker_id = %d and wpr.deleted_at is NULL",
	//*m.ThnAngID, *m.SatkerID)

	query := conn.Table("wbk_komponen wk").
		Select(
			`wspr.id, wspr.code, wspr.name,	
		sum(CASE when wsprb.bulan_id = 1 then 1 else 0 end) 'b1',
        sum(CASE when wsprb.bulan_id = 2 then 1 else 0 end) 'b2',
        sum(CASE when wsprb.bulan_id = 3 then 1 else 0 end) 'b3',
        sum(CASE when wsprb.bulan_id = 4 then 1 else 0 end) 'b4',
        sum(CASE when wsprb.bulan_id = 5 then 1 else 0 end) 'b5',
        sum(CASE when wsprb.bulan_id = 6 then 1 else 0 end) 'b6',
        sum(CASE when wsprb.bulan_id = 7 then 1 else 0 end) 'b7',
        sum(CASE when wsprb.bulan_id = 8 then 1 else 0 end) 'b8',
        sum(CASE when wsprb.bulan_id = 9 then 1 else 0 end) 'b9',
        sum(CASE when wsprb.bulan_id = 10 then 1 else 0 end) 'b10',
        sum(CASE when wsprb.bulan_id = 11 then 1 else 0 end) 'b11',
        sum(CASE when wsprb.bulan_id = 12 then 1 else 0 end) 'b12',
	wspr.wbk_program_ranker_id,
	CONCAT(wk.code,". ",wk.name) as komponen,
	CONCAT(wp.code,". ",wp.name) as program,
	CONCAT(wpr.code,". ",wpr.name) as program_renja, wspr.frekuensi_ranker_id,
        fr.name as frekuensi_waktu`).
		Joins(`inner join wbk_program wp ON wp.wbk_komponen_id = wk.id and wk.deleted_at is NULL`).
		Joins(`inner join wbk_program_ranker wpr on wpr.wbk_program_id = wp.id and wpr.deleted_at is NULL`).
		Joins(`inner join wbk_sub_program_ranker wspr on wpr.id = wspr.wbk_program_ranker_id`).
		Joins(`left outer join wbk_sub_program_ranker_bulan wsprb on wsprb.wbk_sub_program_ranker_id  = wspr.id 
	and wspr.deleted_at is NULL`).
		Joins(`left outer join bulan b on wsprb.bulan_id  = b.id and b.deleted_at is NULL`).
		Joins(`left outer join frekuensi_ranker fr on wspr.frekuensi_ranker_id = fr.id 
        	and fr.deleted_at is NULL`)

	query = r.Filter(ctx, query, *m).Where("wspr.deleted_at is NULL").
		Group(`wspr.id, wspr.code, wspr.name,wspr.wbk_program_ranker_id,
		CONCAT(wk.code,". ",wk.name),
		CONCAT(wp.code,". ",wp.name),
		CONCAT(wpr.code,". ",wpr.name), wspr.frekuensi_ranker_id, fr.name`)
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

func (r *wbkDocUploadSubProgramUraian) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
