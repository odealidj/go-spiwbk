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

type WbkKomponenService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	Get(*abstraction.Context, *dto.WbkKomponenGetRequest) (*dto.WbkKomponenGetInfoResponse, error)
}

type wbkKomponenService struct {
	//SpiAngRepository repository.SpiAng
	WbkKomponenRepository repository.WbkKomponen
	Db                    *gorm.DB
}

func NewWbkKomponenService(f *factory.Factory) *wbkKomponenService {
	wbkKomponenRepository := f.WbkKomponenRepository

	db := f.Db
	return &wbkKomponenService{wbkKomponenRepository, db}

}

func (s *wbkKomponenService) Get(ctx *abstraction.Context,
	payload *dto.WbkKomponenGetRequest) (*dto.WbkKomponenGetInfoResponse, error) {

	var result *dto.WbkKomponenGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkKomponenRepository.Find(ctx,
			&model.WbkKomponenFilter{WbkKomponenEntityFilter: model.WbkKomponenEntityFilter{}}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &dto.WbkKomponenGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
