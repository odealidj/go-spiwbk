package wbk

import (
	"codeid-boiler/internal/abstraction"
	wbk_dto "codeid-boiler/internal/app/dto/wbk"
	"codeid-boiler/internal/app/model"
	wbk_model "codeid-boiler/internal/app/model/wbk"
	"codeid-boiler/internal/app/repository"
	wbk_repo "codeid-boiler/internal/app/repository/wbk"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"fmt"
	"gorm.io/gorm"
	"net/http"
)

type WbkSatkerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk_dto.WbkSatkerUpsertRequest) ([]wbk_dto.WbkSatkerResponse, error)
	GetProgram(*abstraction.Context,
		*wbk_dto.WbkSatkerGetProgramByThnAngAndSatkerAndWbkKomponenRequest) (*wbk_dto.WbkSatkerGetProgramInfoResponse, error)
}

type wbkSatkerService struct {
	WbkSatkerRepository    wbk_repo.WbkSatker
	WbkKomponenRepository  wbk_repo.WbkKomponen
	ThnAngSatkerRepository repository.ThnAngSatker

	Db *gorm.DB
}

func NewWbkSatkerService(f *factory.Factory) *wbkSatkerService {
	wbkSatkerRepository := f.WbkSatkerRepository
	wbkKomponenRepository := f.WbkKomponenRepository
	thnAngSatkerRepository := f.ThnAngSatkerRepository
	db := f.Db
	return &wbkSatkerService{wbkSatkerRepository, wbkKomponenRepository,
		thnAngSatkerRepository, db}

}

func (s *wbkSatkerService) Upsert(ctx *abstraction.Context, payload *wbk_dto.WbkSatkerUpsertRequest) ([]wbk_dto.WbkSatkerResponse, error) {

	var result []wbk_dto.WbkSatkerResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		thnAngSatker, err := s.ThnAngSatkerRepository.FirstOrCreate(ctx, &model.ThnAngSatker{
			ThnAngSatkerEntity: model.ThnAngSatkerEntity{ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID}})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid FirstOrCreate thnangsatker", err.Error())
		}

		wbkKomponens, _, err := s.WbkKomponenRepository.Find(ctx, nil, nil)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid FirstOrCreate thnangsatker", err.Error())
		}

		wbkSatkers := &[]wbk_model.WbkSatker{}

		for _, wbkKomponen := range wbkKomponens {
			*wbkSatkers = append(*wbkSatkers, wbk_model.WbkSatker{
				WbkSatkerEntity: wbk_model.WbkSatkerEntity{ThnAngSatkerID: thnAngSatker.ID,
					WbkKomponenID: wbkKomponen.ID.ID},
			})
		}

		datas, err := s.WbkSatkerRepository.Upsert(ctx, wbkSatkers)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid upsert wbk satker", err.Error())
		}

		for _, wbkSatker := range *datas {
			result = append(result, wbk_dto.WbkSatkerResponse{
				ID:              abstraction.ID{ID: wbkSatker.ID},
				WbkSatkerEntity: wbkSatker.WbkSatkerEntity,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkSatkerService) GetProgram(ctx *abstraction.Context,
	payload *wbk_dto.WbkSatkerGetProgramByThnAngAndSatkerAndWbkKomponenRequest) (*wbk_dto.WbkSatkerGetProgramInfoResponse, error) {

	var result *wbk_dto.WbkSatkerGetProgramInfoResponse

	datas, info, err := s.WbkSatkerRepository.FindProram(ctx,
		&wbk_model.WbkSatkerFilter{WbkSatkerEntityFilter: wbk_model.WbkSatkerEntityFilter{SatkerID: &payload.SatkerID, ThnAngID: &payload.ThnAngID,
			WbkKomponenID: payload.WbkKomponenID}}, &payload.Pagination)
	if err != nil {
		return nil, res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			"Invalid Spi bmn", err.Error())
	}

	//num := 0
	for i, _ := range datas {
		datas[i].Row = i + 1
		datas[i].WbkKomponen = fmt.Sprintf("%s. %s", datas[i].WbkKomponenCode, datas[i].WbkKomponenName)
		datas[i].WbkProgram = fmt.Sprintf("%s. %s", datas[i].WbkProgramCode, datas[i].WbkProgramName)
	}

	result = &wbk_dto.WbkSatkerGetProgramInfoResponse{
		Datas:          &datas,
		PaginationInfo: info,
	}

	return result, nil
}
