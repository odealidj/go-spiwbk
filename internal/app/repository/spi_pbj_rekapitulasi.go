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

type SpiPbjRekapitulasi interface {
	Create(*abstraction.Context, *model.SpiPbjRekapitulasi) (*model.SpiPbjRekapitulasi, error)
	Upsert(*abstraction.Context, *model.SpiPbjRekapitulasi) (*model.SpiPbjRekapitulasi, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Find(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) (*[]model.SpiAngItem, *abstraction.PaginationInfo, error)
	FindSpiPbjRekapitulasiByID(*abstraction.Context, *model.SpiPbjRekapitulasiFilter, *abstraction.Pagination) ([]dto.SpiPbjRekapitulasiGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}
type spiPbjRekapitulasi struct {
	abstraction.Repository
}

func NewSpiPbjRekapitulasi(db *gorm.DB) *spiPbjRekapitulasi {
	return &spiPbjRekapitulasi{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiPbjRekapitulasi) Create(ctx *abstraction.Context, m *model.SpiPbjRekapitulasi) (*model.SpiPbjRekapitulasi, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"spi_ang_id": m.SpiAngID,
		"jenis_rekapitulasi_id": m.JenisRekapitulasiID, "bulan_id": m.BulanID,
		"target": m.Target}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiPbjRekapitulasi) Upsert(ctx *abstraction.Context, m *model.SpiPbjRekapitulasi) (*model.SpiPbjRekapitulasi, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"spi_ang_id", "jenis_rekapitulasi_id", "bulan_id", "target"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiPbjRekapitulasi) FindSpiPbjRekapitulasiByID(ctx *abstraction.Context,
	m *model.SpiPbjRekapitulasiFilter, p *abstraction.Pagination) ([]dto.SpiPbjRekapitulasiGetResponse,
	*abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.SpiPbjRekapitulasiGetResponse
	var info abstraction.PaginationInfo

	//partQuery := fmt.Sprintf("sak.deleted_at is NULL and r.thn_ang_id = %d and r.satker_id = %d",
	//	*m.ThnAngID, *m.SatkerID)

	query := conn.Table("spi_pbj_rekapitulasi spr").
		Select(
			`spr.id, spr.spi_ang_id, spr.jenis_rekapitulasi_id
		,sa.thn_ang_id,ta.year
		,sa.satker_id, s.name as satker_name
		,jr.name as pelaksanaan_kegiatan,
		CASE when spr.bulan_id =1 then spr.target else 0 end+
		CASE when spr.bulan_id =2 then spr.target else 0 end+
		CASE when spr.bulan_id =3 then spr.target else 0 end+
		CASE when spr.bulan_id =4 then spr.target else 0 end+
		CASE when spr.bulan_id =5 then spr.target else 0 end+
		CASE when spr.bulan_id =6 then spr.target else 0 end+
		CASE when spr.bulan_id =7 then spr.target else 0 end+
		CASE when spr.bulan_id =8 then spr.target else 0 end+
		CASE when spr.bulan_id =9 then spr.target else 0 end+
		CASE when spr.bulan_id =10 then spr.target else 0 end+
		CASE when spr.bulan_id =11 then spr.target else 0 end+
		CASE when spr.bulan_id =12 then spr.target else 0 end as Total,
		CASE when spr.bulan_id =1 then spr.target else 0 end B01,
		CASE when spr.bulan_id =2 then spr.target else 0 end B02,
		CASE when spr.bulan_id =3 then spr.target else 0 end B03,
		CASE when spr.bulan_id =4 then spr.target else 0 end B04,
		CASE when spr.bulan_id =5 then spr.target else 0 end B05,
		CASE when spr.bulan_id =6 then spr.target else 0 end B06,
		CASE when spr.bulan_id =7 then spr.target else 0 end B07,
		CASE when spr.bulan_id =8 then spr.target else 0 end B08,
		CASE when spr.bulan_id =9 then spr.target else 0 end B09,
		CASE when spr.bulan_id =10 then spr.target else 0 end B10,
		CASE when spr.bulan_id =11 then spr.target else 0 end B11,
		CASE when spr.bulan_id =12 then spr.target else 0 end B12
	`).
		Joins(`inner join spi_ang sa on spr.spi_ang_id = sa.id and sa.deleted_at IS NULL`).
		Joins(`inner JOIN thn_ang ta ON  sa.thn_ang_id = ta.id and ta.deleted_at IS NULL`).
		Joins(`inner JOIN satker s ON sa.satker_id = s.id and s.deleted_at IS NULL`).
		Joins(`inner JOIN jenis_rekapitulasi jr on spr.jenis_rekapitulasi_id = jr.id and jr.deleted_at IS NULL`).
		Joins(`inner JOIN bulan b on b.id = spr.bulan_id`)

	query = r.Filter(ctx, query, *m).Where("sa.deleted_at IS NULL")
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
		sortBy := "spr.jenis_rekapitulasi_id"
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

func (r *spiPbjRekapitulasi) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
