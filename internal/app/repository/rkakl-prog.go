package repository

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type RkaklProg interface {
	Create(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	Update(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	Delete(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	FindByID(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	//Find(*abstraction.Context, *model.ProgramFilter, *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error)
	FirstOrCreate(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	Upsert(*abstraction.Context, *model.RkaklProg) (*model.RkaklProg, error)
	checkTrx(*abstraction.Context) *gorm.DB
}

type rkaklProg struct {
	abstraction.Repository
}

func NewRkaklProg(db *gorm.DB) *rkaklProg {
	return &rkaklProg{
		abstraction.Repository{
			Db: db,
		},
	}
}

func (r *rkaklProg) Create(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Create(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklProg) Update(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Save(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklProg) Delete(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Delete(&m).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklProg) FindByID(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.First(&m, m.ID).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

/*
func (r *rkaklProg) Find(ctx *abstraction.Context, m *model.ProgramFilter, p *abstraction.Pagination) ([]model.Program, *abstraction.PaginationInfo, error) {
	conn := r.CheckTrx(ctx)

	var err error
	var count int64
	var result []model.Program
	var info abstraction.PaginationInfo

	query := conn.
		Select("*").Find(&result)

	//filter

	query = r.Filter(ctx, query, *m).Where("deleted_at is NULL")
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
		sort := "desc"
		p.Sort = &sort
	}

	if p.SortBy == nil {
		sortBy := "rkakl.id"
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
*/
func (r *rkaklProg) Upsert(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"rkakl_id", "program_id", "biaya"}),
		//UpdateAll: true,
	}).Create(&m).WithContext(ctx.Request().Context()).Error

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklProg) FirstOrCreate(ctx *abstraction.Context, m *model.RkaklProg) (*model.RkaklProg, error) {
	conn := r.CheckTrx(ctx)

	err := conn.FirstOrCreate(&m, map[string]interface{}{"rkakl_id": m.RkaklID, "program_id": m.ProgramID,
		"biaya": m.Biaya}).WithContext(ctx.Request().Context()).Error
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *rkaklProg) checkTrx(ctx *abstraction.Context) *gorm.DB {
	if ctx.Trx != nil {
		return ctx.Trx.Db
	}
	return r.Db
}
