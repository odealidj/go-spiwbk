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

type WbkSubProgramRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk3.WbkSubProgramRankerUpsertRequest) (*wbk3.WbkSubProgramRankerResponse, error)
	Get(*abstraction.Context, *wbk3.WbkSubProgramRankerGetRequest) (*wbk3.WbkSubProgramRankerGetInfoResponse, error)
}

type wbkSubProgramRankerService struct {
	//SpiAngRepository repository.SpiAng
	WbkSubProgramRankerRepository wbk.WbkSubProgramRanker
	Db                            *gorm.DB
}

func NewWbkSubProgramRankerService(f *factory.Factory) *wbkSubProgramRankerService {
	wbkSubProgramRankerRepository := f.WbkSubProgramRankerRepository

	db := f.Db
	return &wbkSubProgramRankerService{wbkSubProgramRankerRepository, db}

}

func (s *wbkSubProgramRankerService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkSubProgramRankerUpsertRequest) (*wbk3.WbkSubProgramRankerResponse, error) {

	var result *wbk3.WbkSubProgramRankerResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkSubProgramRankerRepository.Upsert(ctx, &wbk2.WbkSubProgramRanker{Context: ctx,
			WbkSubProgramRankerEntity: payload.WbkSubProgramRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &wbk3.WbkSubProgramRankerResponse{
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
	payload *wbk3.WbkSubProgramRankerGetRequest) (*wbk3.WbkSubProgramRankerGetInfoResponse, error) {

	var result *wbk3.WbkSubProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkSubProgramRankerRepository.Find(ctx,
			&wbk2.WbkSubProgramRankerFilter{WbkSubProgramRankerEntityFilter: payload.WbkSubProgramRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkSubProgramRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
