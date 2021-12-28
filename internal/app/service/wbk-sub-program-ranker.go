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

type WbkSubProgramRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *dto.WbkSubProgramRankerUpsertRequest) (*dto.WbkSubProgramRankerResponse, error)
	Get(*abstraction.Context, *dto.WbkSubProgramRankerGetRequest) (*dto.WbkSubProgramRankerGetInfoResponse, error)
}

type wbkSubProgramRankerService struct {
	//SpiAngRepository repository.SpiAng
	WbkSubProgramRankerRepository repository.WbkSubProgramRanker
	Db                            *gorm.DB
}

func NewWbkSubProgramRankerService(f *factory.Factory) *wbkSubProgramRankerService {
	wbkSubProgramRankerRepository := f.WbkSubProgramRankerRepository

	db := f.Db
	return &wbkSubProgramRankerService{wbkSubProgramRankerRepository, db}

}

func (s *wbkSubProgramRankerService) Upsert(ctx *abstraction.Context, payload *dto.WbkSubProgramRankerUpsertRequest) (*dto.WbkSubProgramRankerResponse, error) {

	var result *dto.WbkSubProgramRankerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkSubProgramRankerRepository.Upsert(ctx, &model.WbkSubProgramRanker{Context: ctx,
			WbkSubProgramRankerEntity: payload.WbkSubProgramRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &dto.WbkSubProgramRankerResponse{
			ID:                        abstraction.ID{ID: data.ID},
			WbkSubProgramRankerEntity: data.WbkSubProgramRankerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkSubProgramRankerService) Get(ctx *abstraction.Context,
	payload *dto.WbkSubProgramRankerGetRequest) (*dto.WbkSubProgramRankerGetInfoResponse, error) {

	var result *dto.WbkSubProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkSubProgramRankerRepository.Find(ctx,
			&model.WbkSubProgramRankerFilter{WbkSubProgramRankerEntityFilter: payload.WbkSubProgramRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkSubProgramRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
