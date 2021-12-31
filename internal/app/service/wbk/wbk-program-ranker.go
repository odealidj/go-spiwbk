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

type WbkProgramRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk3.WbkProgramRankerUpsertRequest) (*wbk3.WbkProgramRankerResponse, error)
	GetSatkerNilaiByThnAngID(*abstraction.Context, *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetSatkerNilaiInfoResponse, error)
	GetByThnAngIDAndSatkerID(*abstraction.Context, *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetInfoResponse, error)
	Get(*abstraction.Context, *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetInfoResponse, error)
}

type wbkProgramRankerService struct {
	//SpiAngRepository repository.SpiAng
	WbkProgramRankerRepository wbk.WbkProgramRanker
	Db                         *gorm.DB
}

func NewWbkProgramRankerService(f *factory.Factory) *wbkProgramRankerService {
	wbkProgramRankerRepository := f.WbkProgramRankerRepository

	db := f.Db
	return &wbkProgramRankerService{wbkProgramRankerRepository, db}

}

func (s *wbkProgramRankerService) Upsert(ctx *abstraction.Context, payload *wbk3.WbkProgramRankerUpsertRequest) (*wbk3.WbkProgramRankerResponse, error) {

	var result *wbk3.WbkProgramRankerResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkProgramRankerRepository.Upsert(ctx, &wbk2.WbkProgramRanker{Context: ctx,
			WbkProgramRankerEntity: payload.WbkProgramRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		result = &wbk3.WbkProgramRankerResponse{
			ID:                     abstraction.ID{ID: data.ID},
			WbkProgramRankerEntity: data.WbkProgramRankerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *wbkProgramRankerService) GetSatkerNilaiByThnAngID(ctx *abstraction.Context,
	payload *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetSatkerNilaiInfoResponse, error) {

	var result *wbk3.WbkProgramRankerGetSatkerNilaiInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkProgramRankerRepository.FindSatkerNilaiByThnAngID(ctx,
			&wbk2.WbkProgramRankerFilter{WbkProgramRankerEntityFilter: wbk2.WbkProgramRankerEntityFilter{
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

		result = &wbk3.WbkProgramRankerGetSatkerNilaiInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *wbkProgramRankerService) GetByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetInfoResponse, error) {

	var result *wbk3.WbkProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		wbkProgramRankerGetResponses, info, err := s.WbkProgramRankerRepository.FindByThnAngIDAndSatkerID(ctx,
			&wbk2.WbkProgramRankerFilter{WbkProgramRankerEntityFilter: wbk2.WbkProgramRankerEntityFilter{
				ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID},
			}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range wbkProgramRankerGetResponses {
			wbkProgramRankerGetResponses[i].Row = i + 1
		}

		result = &wbk3.WbkProgramRankerGetInfoResponse{
			Datas:          &wbkProgramRankerGetResponses,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *wbkProgramRankerService) Get(ctx *abstraction.Context,
	payload *wbk3.WbkProgramRankerGetRequest) (*wbk3.WbkProgramRankerGetInfoResponse, error) {

	var result *wbk3.WbkProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		wbkProgramRankerGetResponses, info, err := s.WbkProgramRankerRepository.Find(ctx,
			&wbk2.WbkProgramRankerFilter{WbkProgramRankerEntityFilter: payload.WbkProgramRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range wbkProgramRankerGetResponses {
			wbkProgramRankerGetResponses[i].Row = i + 1
		}

		result = &wbk3.WbkProgramRankerGetInfoResponse{
			Datas:          &wbkProgramRankerGetResponses,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
