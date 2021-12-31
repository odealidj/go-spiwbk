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

type WbkProgramService interface {
	Save(*abstraction.Context, *wbk3.WbkProgramUpsertRequest) (*wbk3.WbkProgramResponse, error)
	Upsert(*abstraction.Context, *wbk3.WbkProgramUpsertRequest) (*wbk3.WbkProgramResponse, error)
	Get(*abstraction.Context, *wbk3.WbkProgramGetRequest) (*wbk3.WbkProgramGetInfoResponse, error)
	GetNilaiByThnAngIDAndSatkerID(*abstraction.Context, *wbk3.WbkProgramGetRequest) (*wbk3.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse, error)
}

type wbkProgramService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramRepository wbk.WbkProgram
	Db                   *gorm.DB
}

func NewWbkProgramService(f *factory.Factory) *wbkProgramService {
	wbkProgramRepository := f.WbkProgramRepository

	db := f.Db
	return &wbkProgramService{wbkProgramRepository, db}

}

func (s *wbkProgramService) Save(ctx *abstraction.Context, payload *wbk3.WbkProgramUpsertRequest) (*wbk3.WbkProgramResponse, error) {

	var result *wbk3.WbkProgramResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramRepository.Create(ctx, &wbk2.WbkProgram{Context: ctx,
			WbkProgramEntity: payload.WbkProgramEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program", err.Error())
		}

		result = &wbk3.WbkProgramResponse{
			ID:               abstraction.ID{ID: data.ID},
			WbkProgramEntity: data.WbkProgramEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkProgramService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkProgramUpsertRequest) (*wbk3.WbkProgramResponse, error) {

	var result *wbk3.WbkProgramResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramRepository.Upsert(ctx, &wbk2.WbkProgram{Context: ctx,
			WbkProgramEntity: payload.WbkProgramEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program", err.Error())
		}

		result = &wbk3.WbkProgramResponse{
			ID:               abstraction.ID{ID: data.ID},
			WbkProgramEntity: data.WbkProgramEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkProgramService) Get(ctx *abstraction.Context,
	payload *wbk3.WbkProgramGetRequest) (*wbk3.WbkProgramGetInfoResponse, error) {

	var result *wbk3.WbkProgramGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRepository.Find(ctx,
			&wbk2.WbkProgramFilter{WbkProgramEntityFilter: payload.WbkProgramEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkProgramGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *wbkProgramService) GetNilaiByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *wbk3.WbkProgramGetRequest) (*wbk3.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse, error) {

	var result *wbk3.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRepository.FindByThnAngIDAndSatkerID(ctx,
			&wbk2.WbkProgramFilter{WbkProgramEntityFilter: wbk2.WbkProgramEntityFilter{
				ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID},
			}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkProgramNilaiGetByThnAngIDAndSatkerIDInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
