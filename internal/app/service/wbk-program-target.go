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

type WbkProgramTargetService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	Upsert(*abstraction.Context, *dto.WbkProgramTargetUpsertRequest) (*dto.WbkProgramTargetResponse, error)
	Get(*abstraction.Context, *dto.WbkProgramTargetGetRequest) (*dto.WbkProgramTargetGetInfoResponse, error)
}

type wbkProgramTargetService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramTargetRepository repository.WbkProgramTarget
	Db                         *gorm.DB
}

func NewWbkProgramTargetService(f *factory.Factory) *wbkProgramTargetService {
	wbkProgramTargetRepository := f.WbkProgramTargetRepository

	db := f.Db
	return &wbkProgramTargetService{wbkProgramTargetRepository, db}

}

func (s *wbkProgramTargetService) Upsert(ctx *abstraction.Context, payload *dto.WbkProgramTargetUpsertRequest) (*dto.WbkProgramTargetResponse, error) {

	var result *dto.WbkProgramTargetResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramTargetRepository.Upsert(ctx, &model.WbkProgramTarget{Context: ctx,
			WbkProgramTargetEntity: payload.WbkProgramTargetEntity,
		})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid wbk program tujuan")
		}

		result = &dto.WbkProgramTargetResponse{
			ID:                     int(data.ID),
			WbkProgramTargetEntity: data.WbkProgramTargetEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkProgramTargetService) Get(ctx *abstraction.Context,
	payload *dto.WbkProgramTargetGetRequest) (*dto.WbkProgramTargetGetInfoResponse, error) {

	var result *dto.WbkProgramTargetGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramTargetRepository.Find(ctx,
			&model.WbkProgramTargetFilter{WbkProgramTargetEntityFilter: payload.WbkProgramTargetEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkProgramTargetGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
