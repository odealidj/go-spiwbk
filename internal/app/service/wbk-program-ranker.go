package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
	"net/http"
)

type WbkProgramRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	GetSatkerNilaiByThnAngID(*abstraction.Context, *dto.WbkProgramRankerGetRequest) (*dto.WbkProgramRankerGetSatkerNilaiInfoResponse, error)
	GetByThnAngIDAndSatkerID(*abstraction.Context, *dto.WbkProgramRankerGetRequest) (*dto.WbkProgramRankerGetInfoResponse, error)
}

type wbkProgramRankerService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramRankerRepository repository.WbkProgramRanker
	Db                         *gorm.DB
}

func NewWbkProgramRankerService(f *factory.Factory) *wbkProgramRankerService {
	wbkProgramRankerRepository := f.WbkProgramRankerRepository

	db := f.Db
	return &wbkProgramRankerService{wbkProgramRankerRepository, db}

}

func (s *wbkProgramRankerService) GetSatkerNilaiByThnAngID(ctx *abstraction.Context,
	payload *dto.WbkProgramRankerGetRequest) (*dto.WbkProgramRankerGetSatkerNilaiInfoResponse, error) {

	var result *dto.WbkProgramRankerGetSatkerNilaiInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRankerRepository.FindSatkerNilaiByThnAngID(ctx,
			&model.WbkProgramRankerFilter{WbkProgramRankerEntityFilter: model.WbkProgramRankerEntityFilter{
				ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID},
			}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkProgramRankerGetSatkerNilaiInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *wbkProgramRankerService) GetByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *dto.WbkProgramRankerGetRequest) (*dto.WbkProgramRankerGetInfoResponse, error) {

	var result *dto.WbkProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		wbkProgramRankerGetResponses, info, err := s.WbkProgramRankerRepository.FindByThnAngIDAndSatkerID(ctx,
			&model.WbkProgramRankerFilter{WbkProgramRankerEntityFilter: model.WbkProgramRankerEntityFilter{
				ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID},
			}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range wbkProgramRankerGetResponses {
			wbkProgramRankerGetResponses[i].Row = i + 1
		}

		result = &dto.WbkProgramRankerGetInfoResponse{
			Datas:          &wbkProgramRankerGetResponses,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
