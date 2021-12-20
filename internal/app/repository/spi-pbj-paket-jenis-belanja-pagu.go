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

type SpiPbjPaketJenisBelanjaPagu interface {
	Upsert(*abstraction.Context, *model.SpiPbjPaketJenisBelanjaPagu) (*model.SpiPbjPaketJenisBelanjaPagu, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Find(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) (*[]model.SpiAngItem, *abstraction.PaginationInfo, error)
	FindspiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(*abstraction.Context,
		*model.SpiPbjPaketJenisBelanjaPaguFilter, *abstraction.Pagination) ([]dto.SpiPbjPaketJenisBelanjaPaguGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}
type spiPbjPaketJenisBelanjaPagu struct {
	abstraction.Repository
}

func NewSpiPbjPaketJenisBelanjaPagu(db *gorm.DB) *spiPbjPaketJenisBelanjaPagu {
	return &spiPbjPaketJenisBelanjaPagu{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiPbjPaketJenisBelanjaPagu) Create(ctx *abstraction.Context, m *model.SpiPbjPaketJenisBelanjaPagu) (*model.SpiPbjPaketJenisBelanjaPagu, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"spi_pbj_paket_id": m.SpiPbjPaketID,
		"jenis_belanja_pagu_id": m.JenisBelanjaPaguID, "sub_komponen_akun_id": m.SubKomponenAkunID}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiPbjPaketJenisBelanjaPagu) Upsert(ctx *abstraction.Context, m *model.SpiPbjPaketJenisBelanjaPagu) (*model.SpiPbjPaketJenisBelanjaPagu, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"spi_pbj_paket_id", "jenis_belanja_pagu_id", "sub_komponen_akun_id"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiPbjPaketJenisBelanjaPagu) FindspiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(ctx *abstraction.Context,
	m *model.SpiPbjPaketJenisBelanjaPaguFilter, p *abstraction.Pagination) ([]dto.SpiPbjPaketJenisBelanjaPaguGetResponse, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.SpiPbjPaketJenisBelanjaPaguGetResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("sa.thn_ang_id =%d and sa.satker_id =%d and spp.group_package_value_id =%d and ta.deleted_at is NULL",
		*m.ThnAngID, *m.SatkerID, *m.GroupPackageValueID)

	query := conn.Table("thn_ang ta").
		Select(
			`sa.id as spi_ang_id,
			sa.thn_ang_id ,sa.satker_id , spp.group_package_value_id,
			k.name as paket_name, 
			sum(case when sppjbp.jenis_belanja_pagu_id = 1 then ska.biaya else 0 end) barang,
			sum(case when sppjbp.jenis_belanja_pagu_id = 2 then ska.biaya else 0 end) modal,
			sum(case when sppjbp.jenis_belanja_pagu_id = 3 then ska.biaya else 0 end) sosial,
			sum(case when sppjbp.jenis_belanja_pagu_id = 4 then ska.biaya else 0 end) lainnya,
			IFNULL(ma.name,'') as method_pbj,
			case when spr.bulan_id = 1 then true else false end 'rencana1',
			case when spr.bulan_id = 2 then true else false end 'rencana2',
			case when spr.bulan_id = 3 then true else false end 'rencana3',
			case when spr.bulan_id = 4 then true else false end 'rencana4',
			case when spr.bulan_id = 5 then true else false end 'rencana5',
			case when spr.bulan_id = 6 then true else false end 'rencana6',
			case when spr.bulan_id = 7 then true else false end 'rencana7',
			case when spr.bulan_id = 8 then true else false end 'rencana8',
			case when spr.bulan_id = 9 then true else false end 'rencana9',
			case when spr.bulan_id = 10 then true else false end 'rencana10',
			case when spr.bulan_id = 11 then true else false end 'rencana11',
			case when spr.bulan_id = 12 then true else false end 'rencana12',
			case when spr2.bulan_id = 1 then true else false end 'realisasi1',
			case when spr2.bulan_id = 2 then true else false end 'realisasi2',
			case when spr2.bulan_id = 3 then true else false end 'realisasi3',
			case when spr2.bulan_id = 4 then true else false end 'realisasi4',
			case when spr2.bulan_id = 5 then true else false end 'realisasi5',
			case when spr2.bulan_id = 6 then true else false end 'realisasi6',
			case when spr2.bulan_id = 7 then true else false end 'realisasi7',
			case when spr2.bulan_id = 8 then true else false end 'realisasi8',
			case when spr2.bulan_id = 9 then true else false end 'realisasi9',
			case when spr2.bulan_id = 10 then true else false end 'realisasi10',
			case when spr2.bulan_id = 11 then true else false end 'realisasi11',
			case when spr2.bulan_id = 12 then true else false end 'realisasi12',
			spp.permasalahan, spp.rencana_pemecahan 
		`).
		Joins(`inner join spi_ang sa ON sa.thn_ang_id = ta.id and sa.deleted_at is NULL`).
		Joins(`inner join satker s on sa.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`inner join spi_pbj_paket spp on spp.spi_ang_id = sa.id and spp.deleted_at is NULL`).
		Joins(`inner join group_package_value gpv on spp.group_package_value_id = gpv.id  and gpv.deleted_at is NULL `).
		Joins(`inner join komponen k on spp.komponen_id = k.id and k.deleted_at is NULL`).
		Joins(`inner join sub_komponen sk on sk.komponen_id = k.id and sk.deleted_at is NULL`).
		Joins(`inner join sub_komponen_akun ska on ska.sub_komponen_id = sk.id  and ska.deleted_at is NULL`).
		Joins(`inner join akun a on ska.akun_id = a.id and a.deleted_at is NULL`).
		Joins(`inner JOIN spi_pbj_paket_jenis_belanja_pagu sppjbp on sppjbp.spi_pbj_paket_id = spp.id and sppjbp.deleted_at is NULL`).
		Joins(`INNER join jenis_belanja_pagu jbp on sppjbp.jenis_belanja_pagu_id = jbp.id and jbp.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN method_apbj ma on spp.method_apbj_id = ma.id and ma.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN spi_pbj_rencana spr ON spr.spi_pbj_paket_id = spp.id and spr.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN spi_pbj_realisasi spr2 on spr2.spi_pbj_paket_id = spp.id and spr2.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN bulan b on spr.bulan_id = b.id`).
		Joins(`LEFT OUTER JOIN bulan b2 on spr2.bulan_id = b2.id`)
	//Find(&result)

	query = r.Filter(ctx, query, *m).Where(partQuery).
		Group(`sa.id,sa.thn_ang_id ,sa.satker_id, spp.group_package_value_id,
			k.name,IFNULL(ma.name,''),spr.bulan_id,spr2.bulan_id, spp.permasalahan, spp.rencana_pemecahan
		`)
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

func (r *spiPbjPaketJenisBelanjaPagu) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
