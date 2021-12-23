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

type SpiBmn interface {
	Upsert(*abstraction.Context, *model.SpiBmn) (*model.SpiBmn, error)
	//Update(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Delete(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//FindByID(*abstraction.Context, *model.SpiAngItem) (*model.SpiAngItem, error)
	//Find(*abstraction.Context, *model.SpiAngKesesuaianFilter, *abstraction.Pagination) (*[]model.SpiAngItem, *abstraction.PaginationInfo, error)
	FindSpiBmnByThnAngIDAndSatkerID(*abstraction.Context,
		*model.SpiBmnFilter, *abstraction.Pagination) ([]dto.SpiBmnGetResponse, *abstraction.PaginationInfo, error)
	checkTrx(*abstraction.Context) *gorm.DB
}
type spiBmn struct {
	abstraction.Repository
}

func NewSpiBmn(db *gorm.DB) *spiBmn {
	return &spiBmn{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiBmn) Create(ctx *abstraction.Context, m *model.SpiBmn) (*model.SpiBmn, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"spi_ang_id": m.ID,
		"jenis_bmn_id": m.JenisBmnID, "jenis_bmn_uraian": m.JenisBmnUraian,
		"nilai_bmn": m.NilaiBmn, "pengelola_bmn_satker_id": m.PengelolaBmnSatkerID,
		"pengelola_bmn_pihak_tiga_id": m.PengelolaBmnPihakTigaID, "pengelola_bmn_kso_id": m.PengelolaBmnKsoID,
		"permasalahan_bmn_id": m.PermasalahanBmnID, "uraian_permasalahan": m.UraianPermasalahan,
		"rencana_pemecahan": m.RencanaPemecahan, "realisasi_pemecahan": m.RealisasiPemecahan}).
		WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiBmn) Upsert(ctx *abstraction.Context, m *model.SpiBmn) (*model.SpiBmn, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"spi_ang_id", "jenis_bmn_id", "jenis_bmn_uraian",
			"nilai_bmn", "pengelola_bmn_satker_id", "pengelola_bmn_pihak_tiga_id", "pengelola_bmn_kso_id",
			"permasalahan_bmn_id", "uraian_permasalahan", "rencana_pemecahan", "realisasi_pemecahan"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiBmn) FindSpiBmnByThnAngIDAndSatkerID(ctx *abstraction.Context,
	m *model.SpiBmnFilter, p *abstraction.Pagination) ([]dto.SpiBmnGetResponse, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []dto.SpiBmnGetResponse
	var info abstraction.PaginationInfo

	partQuery := fmt.Sprintf("sa.thn_ang_id =%d and sa.satker_id =%d and sb.deleted_at is NULL",
		*m.ThnAngID, *m.ThnAngID)

	query := conn.Table("spi_bmn sb").
		Select(
			`sb.id, sa.thn_ang_id, sa.satker_id , sb.spi_ang_id, sb.jenis_bmn_id, jb.name as jenis_bmn_name, sb.jenis_bmn_uraian, sb.nilai_bmn, 
	case when sb.pengelola_bmn_satker_id is not null then  sb.pengelola_bmn_satker_id else null end 'pengelola_satker_id',
	case when sb.pengelola_bmn_satker_id is not null then  s.name else null end 'pengelola_satker_name',
	case when sb.pengelola_bmn_pihak_tiga_id is not null then sb.pengelola_bmn_pihak_tiga_id else null end 'pengelola_pihak_tiga_id', 
	case when sb.pengelola_bmn_pihak_tiga_id is not null then pt.name else null end 'pengelola_pihak_tiga_name', 
	case when sb.pengelola_bmn_kso_id is not null then sb.pengelola_bmn_kso_id else null end 'pengelola_kso_id',
	case when sb.pengelola_bmn_kso_id is not null then k.name else null end 'pengelola_kso_name',
	case when sb.permasalahan_bmn_id = 1 then sb.permasalahan_bmn_id else null end 'permasalahan_sengketa_id', 
	case when sb.permasalahan_bmn_id = 1 then sb.uraian_permasalahan else '' end 'permasalahan_sengketa_uraian', 
	case when sb.permasalahan_bmn_id = 2 then sb.permasalahan_bmn_id else null end 'permasalahan_dokumen_id', 
	case when sb.permasalahan_bmn_id = 2 then sb.uraian_permasalahan else '' end 'permasalahan_dokumen_uraian', 
	case when sb.permasalahan_bmn_id = 3 then sb.permasalahan_bmn_id else null end 'permasalahan_hilang_id', 
	case when sb.permasalahan_bmn_id = 3 then sb.uraian_permasalahan else '' end 'permasalahan_hilang_uraian', 
	case when sb.permasalahan_bmn_id = 4 then sb.permasalahan_bmn_id else null end 'permasalahan_rusak_id', 
	case when sb.permasalahan_bmn_id = 4 then sb.uraian_permasalahan else '' end 'permasalahan_rusak_uraian', 
	case when sb.permasalahan_bmn_id = 5 then sb.permasalahan_bmn_id else null end 'permasalahan_lainnya_id', 
	case when sb.permasalahan_bmn_id = 5 then sb.uraian_permasalahan else '' end 'permasalahan_lainnya_uraian', 
	sb.rencana_pemecahan, sb.realisasi_pemecahan 
		`).
		Joins(`inner join spi_ang sa ON sb.spi_ang_id = sa.id and sa.deleted_at is NULL`).
		Joins(`inner join permasalahan_bmn pb ON sb.permasalahan_bmn_id = pb.id and pb.deleted_at is NULL`).
		Joins(`inner JOIN jenis_bmn jb on sb.jenis_bmn_id = jb.id and jb.deleted_at is NULL`).
		Joins(`LEFT outer JOIN pengelola_bmn_satker pbs ON sb.pengelola_bmn_satker_id = pbs.id and pbs.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN satker s on pbs.satker_id = s.id and s.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pengelola_bmn pb2 on pbs.pengelola_bmn_id = pb2.id and pb2.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pengelola_bmn_pihak_tiga pbpt on sb.pengelola_bmn_pihak_tiga_id = pbpt.id and pbpt.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pengelola_bmn pb3 on pbpt.pengelola_bmn_id = pb3.id and pb3.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pihak_tiga pt on pbpt.pihak_tiga_id = pt.id and pt.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pengelola_bmn_kso pbk ON sb.pengelola_bmn_kso_id = pbk.id and pbk.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN pengelola_bmn pb4 on pbk.pengelola_bmn_id = pb4.id and pb4.deleted_at is NULL`).
		Joins(`LEFT OUTER JOIN kso k on pbk.kso_id = k.id and k.deleted_at is NULL`)

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
		sortBy := "sb.jenis_bmn_id"
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

func (r *spiBmn) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
