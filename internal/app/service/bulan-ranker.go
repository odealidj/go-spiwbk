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

type BulanRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *dto.BulanRankerUpsertRequest) (*dto.BulanRankerResponse, error)
	Get(*abstraction.Context, *dto.BulanRankerGetRequest) (*dto.BulanRankerGetInfoResponse, error)
}

type bulanRankerService struct {
	//SpiAngRepository repository.SpiAng
	BulanRankerRepository repository.BulanRanker
	Db                    *gorm.DB
}

func NewBulanRankerService(f *factory.Factory) *bulanRankerService {
	bulanRankerRepository := f.BulanRankerRepository

	db := f.Db
	return &bulanRankerService{bulanRankerRepository, db}

}

func (s *bulanRankerService) Upsert(ctx *abstraction.Context, payload *dto.BulanRankerUpsertRequest) (*dto.BulanRankerResponse, error) {

	var result *dto.BulanRankerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.BulanRankerRepository.Upsert(ctx, &model.BulanRanker{Context: ctx,
			ID:                payload.ID,
			BulanRankerEntity: payload.BulanRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &dto.BulanRankerResponse{
			ID:                abstraction.ID{ID: data.ID.ID},
			BulanRankerEntity: data.BulanRankerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *bulanRankerService) Get(ctx *abstraction.Context,
	payload *dto.BulanRankerGetRequest) (*dto.BulanRankerGetInfoResponse, error) {

	var result *dto.BulanRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.BulanRankerRepository.Find(ctx,
			&model.BulanRankerFilter{BulanRankerEntityFilter: payload.BulanRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.BulanRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
