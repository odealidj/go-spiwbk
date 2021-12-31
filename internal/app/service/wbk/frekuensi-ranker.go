package wbk

import (
	"codeid-boiler/internal/abstraction"
	wbk3 "codeid-boiler/internal/app/dto/wbk"
	wbk2 "codeid-boiler/internal/app/model/wbk"
	"codeid-boiler/internal/app/repository/wbk"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
	"net/http"
)

type FrekuensiRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk3.FrekuensiRankerUpsertRequest) (*wbk3.FrekuensiRankerResponse, error)
	Get(*abstraction.Context, *wbk3.FrekuensiRankerGetRequest) (*wbk3.FrekuensiRankerGetInfoResponse, error)
}

type frekuensiRankerService struct {
	//SpiAngRepository repository.SpiAng
	FrekuensiRankerRepository wbk.FrekuensiRanker
	Db                        *gorm.DB
}

func NewFrekuensiRankerService(f *factory.Factory) *frekuensiRankerService {
	frekuensiRankerRepository := f.FrekuensiRankerRepository

	db := f.Db
	return &frekuensiRankerService{frekuensiRankerRepository, db}

}

func (s *frekuensiRankerService) Upsert(ctx *abstraction.Context, payload *wbk3.FrekuensiRankerUpsertRequest) (*wbk3.FrekuensiRankerResponse, error) {

	var result *wbk3.FrekuensiRankerResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.FrekuensiRankerRepository.Upsert(ctx, &wbk2.FrekuensiRanker{Context: ctx,
			FrekuensiRankerEntity: payload.FrekuensiRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &wbk3.FrekuensiRankerResponse{
			ID:                    abstraction.ID{ID: data.ID},
			FrekuensiRankerEntity: data.FrekuensiRankerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *frekuensiRankerService) Get(ctx *abstraction.Context,
	payload *wbk3.FrekuensiRankerGetRequest) (*wbk3.FrekuensiRankerGetInfoResponse, error) {

	var result *wbk3.FrekuensiRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.FrekuensiRankerRepository.Find(ctx,
			&wbk2.FrekuensiRankerFilter{FrekuensiRankerEntityFilter: payload.FrekuensiRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.FrekuensiRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
