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

type FrekuensiRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *dto.FrekuensiRankerUpsertRequest) (*dto.FrekuensiRankerResponse, error)
	Get(*abstraction.Context, *dto.FrekuensiRankerGetRequest) (*dto.FrekuensiRankerGetInfoResponse, error)
}

type frekuensiRankerService struct {
	//SpiAngRepository repository.SpiAng
	FrekuensiRankerRepository repository.FrekuensiRanker
	Db                        *gorm.DB
}

func NewFrekuensiRankerService(f *factory.Factory) *frekuensiRankerService {
	frekuensiRankerRepository := f.FrekuensiRankerRepository

	db := f.Db
	return &frekuensiRankerService{frekuensiRankerRepository, db}

}

func (s *frekuensiRankerService) Upsert(ctx *abstraction.Context, payload *dto.FrekuensiRankerUpsertRequest) (*dto.FrekuensiRankerResponse, error) {

	var result *dto.FrekuensiRankerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.FrekuensiRankerRepository.Upsert(ctx, &model.FrekuensiRanker{Context: ctx,
			ID:                    payload.ID,
			FrekuensiRankerEntity: payload.FrekuensiRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &dto.FrekuensiRankerResponse{
			ID:                    abstraction.ID{ID: data.ID.ID},
			FrekuensiRankerEntity: data.FrekuensiRankerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *frekuensiRankerService) Get(ctx *abstraction.Context,
	payload *dto.FrekuensiRankerGetRequest) (*dto.FrekuensiRankerGetInfoResponse, error) {

	var result *dto.FrekuensiRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.FrekuensiRankerRepository.Find(ctx,
			&model.FrekuensiRankerFilter{FrekuensiRankerEntityFilter: payload.FrekuensiRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.FrekuensiRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
