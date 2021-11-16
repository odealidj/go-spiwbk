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

type SpiSdmItem interface {
	Create(*abstraction.Context, *model.SpiSdmItem) (*model.SpiSdmItem, error)
	Update(*abstraction.Context, *model.SpiSdmItem) (*model.SpiSdmItem, error)
	Delete(*abstraction.Context, *model.SpiSdmItem) (*model.SpiSdmItem, error)
	FindByID(*abstraction.Context, *model.SpiSdmItem) (*model.SpiSdmItem, error)
	Find(*abstraction.Context, *model.SpiSdmItemFilter, *abstraction.Pagination) (*[]model.SpiSdmItem, *abstraction.PaginationInfo, error)
	ViewSpiSdmItemBySpiSdmID(*abstraction.Context, *dto.SpiSdmItemViewBySpiSdmIDRequest) ([]dto.SpiSdmItemViewBySpiSdmIDResponse, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type spiSdmItem struct {
	abstraction.Repository
}

func NewSpiSdmItem(db *gorm.DB) *spiSdmItem {
	return &spiSdmItem{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *spiSdmItem) Create(ctx *abstraction.Context, m *model.SpiSdmItem) (*model.SpiSdmItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiSdmItem) Update(ctx *abstraction.Context, m *model.SpiSdmItem) (*model.SpiSdmItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiSdmItem) Delete(ctx *abstraction.Context, m *model.SpiSdmItem) (*model.SpiSdmItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiSdmItem) FindByID(ctx *abstraction.Context, m *model.SpiSdmItem) (*model.SpiSdmItem, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Where("id = ?", m.ID).First(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *spiSdmItem) Find(ctx *abstraction.Context, m *model.SpiSdmItemFilter, p *abstraction.Pagination) (*[]model.SpiSdmItem, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.SpiSdmItem
	var info abstraction.PaginationInfo

	query := conn.Model(&model.SpiSdmItem{})

	//filter
	query = r.Filter(ctx, query, *m)
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
		}
		ChErr <- nil
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
		return &result, &info, err
	}

	// sort
	if p.Sort == nil {
		sort := "desc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "id"
		p.SortBy = &sortBy
	}

	p.Count = count

	sort := fmt.Sprintf("%s %s", *p.SortBy, *p.Sort)
	query = query.Order(sort)

	info = abstraction.PaginationInfo{
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
	p.Count = count

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
		return &result, &info, err
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

	return &result, &info, nil

}

func (r *spiSdmItem) ViewSpiSdmItemBySpiSdmID(ctx *abstraction.Context, m *dto.SpiSdmItemViewBySpiSdmIDRequest) ([]dto.SpiSdmItemViewBySpiSdmIDResponse, error) {
	conn := r.CheckTrx(ctx)

	var result []dto.SpiSdmItemViewBySpiSdmIDResponse

	if err := conn.Raw("select DISTINCT spi_sdm_id, a.uraian, "+
		"MAX(CASE when a.jsid = 1 then a.value else \"\" end) as KPA, "+
		"MAX(CASE when a.jsid = 2 then a.value else \"\" end) as PPK, "+
		"MAX(CASE when a.jsid = 3 then a.value else \"\" end) as PPSPM, "+
		"MAX(CASE when a.jsid = 4 then a.value else \"\" end) as BendaharaPengeluaran, "+
		"MAX(CASE when a.jsid = 5 then a.value else \"\" end) as BendaharaPenerimaan "+
		"from "+
		"(SELECT spi_sdm_id, js.id as jsid, p.nik as value, 'Nik' as uraian  FROM spi_sdm_item ssi "+
		"INNER JOIN jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"where ssi.deleted_at is NULL "+
		"UNION all "+
		"SELECT spi_sdm_id, js.id as jsid, p.name as value, 'Nama Lengkap' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"where ssi.deleted_at is NULL "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, CONCAT (ssi.nosk, \" / \", DATE_FORMAT(ssi.tgl_sk, '%d-%m-%Y')) as value, 'Nomor dan tanggal SK' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"where ssi.deleted_at is NULL "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, p.phone  as value, 'Nomor HP' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, p.email  as value, 'Email' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, p.email  as value, 'Pendidikan Terakhir' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, ''  as value, 'Sertifikasi:' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, jc.id  as value, CONCAT('a. ','Pengadaan barang/jasa') as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id and jc.id= 1 "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, jc.id  as value, CONCAT('b. ','Bendahara') as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id  and jc.id= 2 "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, jc.id  as value, CONCAT('c. ','Standar Akuntansi Pemerintah') as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id  "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id and jc.id= 3 "+
		"where ssi.deleted_at is NULL "+
		"UNION ALL "+
		"SELECT spi_sdm_id, js.id as jsid, ssi.usulan  as value, 'Usulan/Rencana Pengembangan SDM' as uraian  "+
		"FROM spi_sdm_item ssi inner join jenis_sdm js ON ssi.jenis_sdm_id = js.id "+
		"INNER JOIN pegawai p ON ssi.pegawai_id = p.id "+
		"LEFT OUTER join jenis_certificate jc on ssi.jenis_certficate_id = jc.id where ssi.spi_sdm_id = ? "+
		"and ssi.deleted_at is NULL) a GROUP by spi_sdm_id,a.uraian", m.SpiSdmID).Scan(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (r *spiSdmItem) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
