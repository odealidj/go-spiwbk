package wbk

import (
	"codeid-boiler/internal/abstraction"
	wbk_dto "codeid-boiler/internal/app/dto/wbk"
	wbk_model "codeid-boiler/internal/app/model/wbk"
	"codeid-boiler/internal/app/repository/wbk"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"strings"
)

type WbkSubProgramRankerService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk_dto.WbkSubProgramRankerUpsertRequest) (*wbk_dto.WbkSubProgramRankerResponse, error)
	Get(*abstraction.Context, *wbk_dto.WbkSubProgramRankerGetRequest) (*wbk_dto.WbkSubProgramRankerGetInfoResponse, error)
}

type wbkSubProgramRankerService struct {
	//SpiAngRepository repository.SpiAng
	WbkSubProgramRankerRepository      wbk.WbkSubProgramRanker
	WbkSubProgramRankerBulanRepository wbk.WbkSubProgramRankerBulan
	Db                                 *gorm.DB
}

func NewWbkSubProgramRankerService(f *factory.Factory) *wbkSubProgramRankerService {
	wbkSubProgramRankerRepository := f.WbkSubProgramRankerRepository
	wbkSubProgramRankerBulanRepository := f.WbkSubProgramRankerBulanRepository
	db := f.Db
	return &wbkSubProgramRankerService{wbkSubProgramRankerRepository,
		wbkSubProgramRankerBulanRepository, db}

}

func (s *wbkSubProgramRankerService) Upsert(ctx *abstraction.Context, payload *wbk_dto.WbkSubProgramRankerUpsertRequest) (*wbk_dto.WbkSubProgramRankerResponse, error) {

	var result *wbk_dto.WbkSubProgramRankerResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkSubProgramRankerRepository.Upsert(ctx, &wbk_model.WbkSubProgramRanker{Context: ctx,
			WbkSubProgramRankerEntity: payload.WbkSubProgramRankerEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		if payload.BulanID != nil {

			dataBulans := strings.Split(*payload.BulanID, ",")
			if len(dataBulans) > 0 {

				for i, dataBulan := range dataBulans {
					bulanID, err := strconv.Atoi(dataBulan)
					if err != nil {
						return res.CustomErrorBuilderWithData(http.StatusBadRequest,
							"Invalid bulan "+strconv.Itoa(i), err.Error())
					}
					_, err = s.WbkSubProgramRankerBulanRepository.Upsert(ctx, &wbk_model.WbkSubProgramRankerBulan{Context: ctx,
						WbkSubProgramRankerBulanEntity: wbk_model.WbkSubProgramRankerBulanEntity{
							WbkSubProgramRankerID: int(data.ID), BulanID: bulanID,
						}})
					if err != nil {
						return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
							"Invalid bulan "+strconv.Itoa(i), err.Error())
					}

				}
			}
		}

		result = &wbk_dto.WbkSubProgramRankerResponse{
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
	payload *wbk_dto.WbkSubProgramRankerGetRequest) (*wbk_dto.WbkSubProgramRankerGetInfoResponse, error) {

	var result *wbk_dto.WbkSubProgramRankerGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkSubProgramRankerRepository.Find(ctx,
			&wbk_model.WbkSubProgramRankerFilter{WbkSubProgramRankerEntityFilter: payload.WbkSubProgramRankerEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk_dto.WbkSubProgramRankerGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
