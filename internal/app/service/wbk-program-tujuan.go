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

type WbkProgramTujuanService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	Upsert(*abstraction.Context, *dto.WbkProgramTujuanUpsertRequest) (*dto.WbkProgramTujuanResponse, error)
	Get(*abstraction.Context, *dto.WbkProgramTujuanGetRequest) (*dto.WbkProgramTujuanGetInfoResponse, error)
}

type wbkProgramTujuanService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramTujuanRepository repository.WbkProgramTujuan
	Db                         *gorm.DB
}

func NewWbkProgramTujuanService(f *factory.Factory) *wbkProgramTujuanService {
	wbkProgramTujuanRepository := f.WbkProgramTujuanRepository

	db := f.Db
	return &wbkProgramTujuanService{wbkProgramTujuanRepository, db}

}

func (s *wbkProgramTujuanService) Upsert(ctx *abstraction.Context, payload *dto.WbkProgramTujuanUpsertRequest) (*dto.WbkProgramTujuanResponse, error) {

	var result *dto.WbkProgramTujuanResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramTujuanRepository.Upsert(ctx, &model.WbkProgramTujuan{Context: ctx,
			WbkProgramTujuanEntity: payload.WbkProgramTujuanEntity,
		})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid wbk program tujuan")
		}

		result = &dto.WbkProgramTujuanResponse{
			ID:                     abstraction.ID{ID: data.ID},
			WbkProgramTujuanEntity: data.WbkProgramTujuanEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkProgramTujuanService) Get(ctx *abstraction.Context,
	payload *dto.WbkProgramTujuanGetRequest) (*dto.WbkProgramTujuanGetInfoResponse, error) {

	var result *dto.WbkProgramTujuanGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramTujuanRepository.Find(ctx,
			&model.WbkProgramTujuanFilter{WbkProgramTujuanEntityFilter: payload.WbkProgramTujuanEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkProgramTujuanGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
