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

type WbkSubProgramUraianService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)

	Upsert(*abstraction.Context, *wbk_dto.WbkSubProgramUraianUpsertRequest) (*wbk_dto.WbkSubProgramUraianResponse, error)
	Get(*abstraction.Context, *wbk_dto.WbkSubProgramUraianGetRequest) (*wbk_dto.WbkSubProgramUraianGetInfoResponse, error)
}

type wbkSubProgramUraianService struct {
	//SpiAngRepository repository.SpiAng
	WbkSubProgramUraianRepository      wbk.WbkSubProgramUraian
	WbkSubProgramUraianBulanRepository wbk.WbkSubProgramUraianBulan
	Db                                 *gorm.DB
}

func NewWbkSubProgramUraianService(f *factory.Factory) *wbkSubProgramUraianService {
	wbkSubProgramUraianRepository := f.WbkSubProgramUraianRepository
	wbkSubProgramUraianBulanRepository := f.WbkSubProgramUraianBulanRepository
	db := f.Db
	return &wbkSubProgramUraianService{wbkSubProgramUraianRepository,
		wbkSubProgramUraianBulanRepository, db}

}

func (s *wbkSubProgramUraianService) Upsert(ctx *abstraction.Context, payload *wbk_dto.WbkSubProgramUraianUpsertRequest) (*wbk_dto.WbkSubProgramUraianResponse, error) {

	var result *wbk_dto.WbkSubProgramUraianResponse
	//var data *model.ThnAng

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.WbkSubProgramUraianRepository.Upsert(ctx, &wbk_model.WbkSubProgramUraian{Context: ctx,
			WbkSubProgramUraianEntity: payload.WbkSubProgramUraianEntity,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid wbk program ranker", err.Error())
		}

		dataBulans := strings.Split(payload.BulanID, ",")
		if len(dataBulans) > 0 {

			for i, dataBulan := range dataBulans {
				bulanID, err := strconv.Atoi(dataBulan)
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusBadRequest,
						"Invalid bulan "+strconv.Itoa(i), err.Error())
				}
				_, err = s.WbkSubProgramUraianBulanRepository.Upsert(ctx, &wbk_model.WbkSubProgramUraianBulan{Context: ctx,
					WbkSubProgramUraianBulanEntity: wbk_model.WbkSubProgramUraianBulanEntity{
						WbkSubProgramUraianID: int(data.ID), BulanID: bulanID,
					}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Invalid bulan "+strconv.Itoa(i), err.Error())
				}

			}
		}

		result = &wbk_dto.WbkSubProgramUraianResponse{
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
	payload *wbk_dto.WbkSubProgramUraianGetRequest) (*wbk_dto.WbkSubProgramUraianGetInfoResponse, error) {

	var result *wbk_dto.WbkSubProgramUraianGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkSubProgramUraianRepository.Find(ctx,
			&wbk_model.WbkSubProgramUraianFilter{WbkSubProgramUraianEntityFilter: payload.WbkSubProgramUraianEntityFilter}, &payload.PaginationArr)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk_dto.WbkSubProgramUraianGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
