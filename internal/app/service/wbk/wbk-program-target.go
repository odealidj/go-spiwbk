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

type WbkProgramTargetService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	Upsert(*abstraction.Context, *wbk3.WbkProgramTargetUpsertRequest) (*wbk3.WbkProgramTargetResponse, error)
	Get(*abstraction.Context, *wbk3.WbkProgramTargetGetRequest) (*wbk3.WbkProgramTargetGetInfoResponse, error)
}

type wbkProgramTargetService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramTargetRepository wbk.WbkProgramTarget
	Db                         *gorm.DB
}

func NewWbkProgramTargetService(f *factory.Factory) *wbkProgramTargetService {
	wbkProgramTargetRepository := f.WbkProgramTargetRepository

	db := f.Db
	return &wbkProgramTargetService{wbkProgramTargetRepository, db}

}

func (s *wbkProgramTargetService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkProgramTargetUpsertRequest) (*wbk3.WbkProgramTargetResponse, error) {

	var result *wbk3.WbkProgramTargetResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramTargetRepository.Upsert(ctx, &wbk2.WbkProgramTarget{Context: ctx,
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

		result = &wbk3.WbkProgramTargetResponse{
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
	payload *wbk3.WbkProgramTargetGetRequest) (*wbk3.WbkProgramTargetGetInfoResponse, error) {

	var result *wbk3.WbkProgramTargetGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramTargetRepository.Find(ctx,
			&wbk2.WbkProgramTargetFilter{WbkProgramTargetEntityFilter: payload.WbkProgramTargetEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkProgramTargetGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
