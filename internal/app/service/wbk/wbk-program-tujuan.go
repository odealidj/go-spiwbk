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

type WbkProgramTujuanService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	Upsert(*abstraction.Context, *wbk3.WbkProgramTujuanUpsertRequest) (*wbk3.WbkProgramTujuanResponse, error)
	Get(*abstraction.Context, *wbk3.WbkProgramTujuanGetRequest) (*wbk3.WbkProgramTujuanGetInfoResponse, error)
}

type wbkProgramTujuanService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramTujuanRepository wbk.WbkProgramTujuan
	Db                         *gorm.DB
}

func NewWbkProgramTujuanService(f *factory.Factory) *wbkProgramTujuanService {
	wbkProgramTujuanRepository := f.WbkProgramTujuanRepository

	db := f.Db
	return &wbkProgramTujuanService{wbkProgramTujuanRepository, db}

}

func (s *wbkProgramTujuanService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkProgramTujuanUpsertRequest) (*wbk3.WbkProgramTujuanResponse, error) {

	var result *wbk3.WbkProgramTujuanResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramTujuanRepository.Upsert(ctx, &wbk2.WbkProgramTujuan{Context: ctx,
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

		result = &wbk3.WbkProgramTujuanResponse{
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
	payload *wbk3.WbkProgramTujuanGetRequest) (*wbk3.WbkProgramTujuanGetInfoResponse, error) {

	var result *wbk3.WbkProgramTujuanGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramTujuanRepository.Find(ctx,
			&wbk2.WbkProgramTujuanFilter{WbkProgramTujuanEntityFilter: payload.WbkProgramTujuanEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkProgramTujuanGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
