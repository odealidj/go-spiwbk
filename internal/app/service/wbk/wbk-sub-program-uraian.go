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

type WbkSubProgramUraianService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk3.WbkSubProgramUraianUpsertRequest) (*wbk3.WbkSubProgramUraianResponse, error)
	Get(*abstraction.Context, *wbk3.WbkSubProgramUraianGetRequest) (*wbk3.WbkSubProgramUraianGetInfoResponse, error)
}

type wbkSubProgramUraianService struct {
	//SpiAngRepository repository.SpiAng
	WbkSubProgramUraianRepository wbk.WbkSubProgramUraian
	Db                            *gorm.DB
}

func NewWbkSubProgramUraianService(f *factory.Factory) *wbkSubProgramUraianService {
	wbkSubProgramUraianRepository := f.WbkSubProgramUraianRepository

	db := f.Db
	return &wbkSubProgramUraianService{wbkSubProgramUraianRepository, db}

}

func (s *wbkSubProgramUraianService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkSubProgramUraianUpsertRequest) (*wbk3.WbkSubProgramUraianResponse, error) {

	var result *wbk3.WbkSubProgramUraianResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkSubProgramUraianRepository.Upsert(ctx, &wbk2.WbkSubProgramUraian{Context: ctx,
			WbkSubProgramUraianEntity: payload.WbkSubProgramUraianEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &wbk3.WbkSubProgramUraianResponse{
			ID:                        abstraction.ID{ID: data.ID},
			WbkSubProgramUraianEntity: data.WbkSubProgramUraianEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkSubProgramUraianService) Get(ctx *abstraction.Context,
	payload *wbk3.WbkSubProgramUraianGetRequest) (*wbk3.WbkSubProgramUraianGetInfoResponse, error) {

	var result *wbk3.WbkSubProgramUraianGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkSubProgramUraianRepository.Find(ctx,
			&wbk2.WbkSubProgramUraianFilter{WbkSubProgramUraianEntityFilter: payload.WbkSubProgramUraianEntityFilter}, &payload.PaginationArr)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkSubProgramUraianGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
