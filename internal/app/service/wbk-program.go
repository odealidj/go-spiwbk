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

type WbkProgramService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	Get(*abstraction.Context, *dto.WbkProgramGetRequest) (*dto.WbkProgramGetInfoResponse, error)
	GetNilaiByThnAngIDAndSatkerID(*abstraction.Context, *dto.WbkProgramGetRequest) (*dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse, error)
}

type wbkProgramService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramRepository repository.WbkProgram
	Db                   *gorm.DB
}

func NewWbkProgramService(f *factory.Factory) *wbkProgramService {
	wbkProgramRepository := f.WbkProgramRepository

	db := f.Db
	return &wbkProgramService{wbkProgramRepository, db}

}

func (s *wbkProgramService) Get(ctx *abstraction.Context,
	payload *dto.WbkProgramGetRequest) (*dto.WbkProgramGetInfoResponse, error) {

	var result *dto.WbkProgramGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRepository.Find(ctx,
			&model.WbkProgramFilter{WbkProgramEntityFilter: model.WbkProgramEntityFilter{}}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkProgramGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *wbkProgramService) GetNilaiByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *dto.WbkProgramGetRequest) (*dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse, error) {

	var result *dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRepository.FindByThnAngIDAndSatkerID(ctx,
			&model.WbkProgramFilter{WbkProgramEntityFilter: model.WbkProgramEntityFilter{
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

		result = &dto.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
