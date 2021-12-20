package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SubKomponenAkun interface {
	Create(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Update(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Delete(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	FindByID(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	Upsert(*abstraction.Context, *model.SubKomponenAkun) (*model.SubKomponenAkun, error)
	FindByThnAngIDAndSatkerID(*abstraction.Context, *dto.SubKomponenAkunFindByThnAngIDAndSatkerIDRequest) (*[]dto.SubKomponenAkunFindByThnAngIDAndSatkerIDResponse, error)
	FindByKomponenID(*abstraction.Context, *dto.SubKomponenAkunFindByKomponenIDRequest) (*[]dto.SubKomponenAkunFindByKomponenIDResponse, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type subKomponenAkun struct {
	abstraction.Repository
}

func NewSubKomponenAkun(db *gorm.DB) *subKomponenAkun {
	return &subKomponenAkun{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *subKomponenAkun) Create(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Update(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Delete(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) FindByID(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) Upsert(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"sub_komponen_id", "akun_id", "name", "biaya", "sdcp"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) FirstOrCreate(ctx *abstraction.Context, m *model.SubKomponenAkun) (*model.SubKomponenAkun, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"sub_komponen_id": m.SubKomponenID, "akun_id": m.AkunID,
		"name": m.Name, "biaya": m.Biaya, "sdcp": m.Sdcp}).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *subKomponenAkun) FindByThnAngIDAndSatkerID(ctx *abstraction.Context, d *dto.SubKomponenAkunFindByThnAngIDAndSatkerIDRequest) (*[]dto.SubKomponenAkunFindByThnAngIDAndSatkerIDResponse, error) {
	conn := r.CheckTrx(ctx)

	var result *[]dto.SubKomponenAkunFindByThnAngIDAndSatkerIDResponse

	sql := fmt.Sprintf(`SELECT ska.*,a.code as akun_code, a.name as akun_name FROM rkakl r 
		inner join rkakl_prog rp on rp.rkakl_id = r.id and rp.deleted_at is NULL 
		inner join program p on rp.program_id = p.id and p.deleted_at is NULL 
		INNER join prog_kegiatan pk on pk.rkakl_prog_id = rp.id and pk.deleted_at is NULL 
		inner join kegiatan k on pk.kegiatan_id = k.id and k.deleted_at is NULL 
		inner join kegiatan_output ko on ko.prog_kegiatan_id = pk.id and ko.deleted_at is NULL 
		inner join output o on ko.output_id = o.id and o.deleted_at is NULL 
		INNER JOIN kegiatan_output_location kol on kol.kegiatan_output_id = ko.id and kol.deleted_at is NULL
		inner JOIN sub_output so on so.kegiatan_output_location_id = kol.id and so.deleted_at is NULL 
		inner JOIN komponen k2 on k2.sub_output_id = so.id and k2.deleted_at is NULL 
		inner JOIN sub_komponen sk on sk.komponen_id = k2.id and sk.deleted_at is NULL 
		INNER JOIN sub_komponen_akun ska on ska.sub_komponen_id = sk.id and ska.deleted_at is NULL 
		INNER JOIN akun a on ska.akun_id = a.id and a.deleted_at is NULL 
		WHERE r.thn_ang_id = %d and r.satker_id = %d
		and r.deleted_at is NULL`, d.ThnAngID, d.ThnAngID)
	err := conn.Raw(sql).Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *subKomponenAkun) FindByKomponenID(ctx *abstraction.Context, d *dto.SubKomponenAkunFindByKomponenIDRequest) (*[]dto.SubKomponenAkunFindByKomponenIDResponse, error) {
	conn := r.CheckTrx(ctx)

	var result *[]dto.SubKomponenAkunFindByKomponenIDResponse

	sql := fmt.Sprintf(`SELECT ska.*,a.code as akun_code, a.name as akun_name FROM rkakl r 
		inner join rkakl_prog rp on rp.rkakl_id = r.id and rp.deleted_at is NULL 
		inner join program p on rp.program_id = p.id and p.deleted_at is NULL 
		INNER join prog_kegiatan pk on pk.rkakl_prog_id = rp.id and pk.deleted_at is NULL 
		inner join kegiatan k on pk.kegiatan_id = k.id and k.deleted_at is NULL 
		inner join kegiatan_output ko on ko.prog_kegiatan_id = pk.id and ko.deleted_at is NULL 
		inner join output o on ko.output_id = o.id and o.deleted_at is NULL 
		INNER JOIN kegiatan_output_location kol on kol.kegiatan_output_id = ko.id and kol.deleted_at is NULL
		inner JOIN sub_output so on so.kegiatan_output_location_id = kol.id and so.deleted_at is NULL 
		inner JOIN komponen k2 on k2.sub_output_id = so.id and k2.deleted_at is NULL 
		inner JOIN sub_komponen sk on sk.komponen_id = k2.id and sk.deleted_at is NULL 
		INNER JOIN sub_komponen_akun ska on ska.sub_komponen_id = sk.id and ska.deleted_at is NULL 
		INNER JOIN akun a on ska.akun_id = a.id and a.deleted_at is NULL 
		WHERE k2.id = %d
		and r.deleted_at is NULL`, d.KomponenID)
	err := conn.Raw(sql).Find(&result).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (r *subKomponenAkun) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
