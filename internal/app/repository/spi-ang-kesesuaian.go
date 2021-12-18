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

type SpiAngKesesuaian interface {
	Create(*abstraction.Context, *model.SpiAngKesesuaian) (*model.SpiAngKesesuaian, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Find(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) (*[]model.SpiAngItem, *abstraction.PaginationInfo, error)
	FindSpiKesesuaianByThnAngIDAndSatkerID(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) ([]dto.SpiAngKesesuaianGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}
type spiAngKesesuaian struct {
	abstraction.Repository
}

func NewSpiAngKesesuaian(db *gorm.DB) *spiAngKesesuaian {
	return &spiAngKesesuaian{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiAngKesesuaian) Create(ctx *abstraction.Context, m *model.SpiAngKesesuaian) (*model.SpiAngKesesuaian, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"spi_ang_item_id": m.SpiAngItemID,
		"jenis_kesesuaian_id": m.JenisKesesuaianID, "jenis_pengendali_id": m.JenisPengendaliID,
		"is_check": m.IsCheck}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiAngKesesuaian) FindSpiKesesuaianByThnAngIDAndSatkerID(ctx *abstraction.Context,
	m *model.SpiAngKesesuaianFilter, p *abstraction.Pagination) ([]dto.SpiAngKesesuaianGetResponse, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.SpiAngKesesuaianGetResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("sak.deleted_at is NULL and r.thn_ang_id = %d and r.satker_id = %d",
		*m.ThnAngID, *m.SatkerID)

	query := conn.Table("spi_ang_kesesuaian sak").
		Select(
			`sak.id as spi_ang_kesesuaian_id,sak.spi_ang_item_id, sak.jenis_kesesuaian_id,sak.jenis_pengendali_id,  
					sai.spi_ang_id, sai.komponen_id, sa.thn_ang_id, sa.satker_id, 
					CONCAT(p.code,"/",k2.code,"/",o.code,"/",k.code) as program_kegiatan_output_komponen,
					CONCAT(jk.code, ". ",jk.name) as jenis_kendali_name,
					IFNULL(case when sak.jenis_pengendali_id  = 1 then 
						case when sak.is_check = true then true end
						end, FALSE) pengusul_ya,
					IFNULL(case when sak.jenis_pengendali_id  = 1 then 
						case when sak.is_check = false then false end
						end, FALSE) pengusul_tidak,
					IFNULL(case when sak.jenis_pengendali_id  = 2 then 
						case when sak.is_check = true then true end
						end, FALSE )keu_satker_ya,
					IFNULL(case when sak.jenis_pengendali_id  = 2 then 
						case when sak.is_check = false then false end
						end, FALSE) keu_satker_tidak,	
					IFNULL(case when sak.jenis_pengendali_id  = 3 then 
						case when sak.is_check = true then true end
						end, FALSE )keu_eselon1_ya,
					IFNULL(case when sak.jenis_pengendali_id  = 3 then 
						case when sak.is_check = false then false end
						end, FALSE) keu_eselon1_tidak	
		`).
		Joins(`inner join jenis_pengendali jp ON sak.jenis_pengendali_id = jp.id and jp.deleted_at is NULL`).
		Joins(`inner join jenis_kesesuaian jk on sak.jenis_kesesuaian_id = jk.id  and jk.deleted_at is NULL`).
		Joins(`inner JOIN spi_ang_item sai ON sak.spi_ang_item_id = sai.id and sai.deleted_at is NULL`).
		Joins(`INNER JOIN spi_ang sa on sai.spi_ang_id = sa.id and sa.deleted_at is NULL`).
		Joins(`INNER JOIN komponen k on sai.komponen_id = k.id AND k.deleted_at is NULL`).
		Joins(`INNER JOIN sub_output so on k.sub_output_id = so.id and so.deleted_at is NULL`).
		Joins(`INNER JOIN kegiatan_output_location kol on so.kegiatan_output_location_id = kol.id and kol.deleted_at is NULL`).
		Joins(`INNER JOIN kegiatan_output ko on kol.kegiatan_output_id = ko.id and ko.deleted_at is NULL`).
		Joins(`INNER JOIN output o on ko.output_id = o.id and o.deleted_at is NULL`).
		Joins(`INNER JOIN prog_kegiatan pk on ko.prog_kegiatan_id = pk.id and pk.deleted_at is NULL`).
		Joins(`INNER JOIN kegiatan k2 on pk.kegiatan_id = k2.id and k2.deleted_at is NULL`).
		Joins(`INNER JOIN rkakl_prog rp on pk.rkakl_prog_id = rp.id and rp.deleted_at is NULL`).
		Joins(`INNER JOIN program p on rp.program_id = p.id  and p.deleted_at is NULL`).
		Joins(`INNER JOIN rkakl r on rp.rkakl_id = r.id and r.deleted_at is NULL`)
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
		sortBy := "sai.komponen_id, sak.jenis_kesesuaian_id, sak.jenis_pengendali_id"
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

func (r *spiAngKesesuaian) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
