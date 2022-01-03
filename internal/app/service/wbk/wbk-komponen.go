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

type WbkKomponenService interface {
	//Save(*abstraction.Context, *dto.WbkProgramRankerSaveRequest) (*dto.WbkProgramRankerResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	Get(*abstraction.Context, *wbk3.WbkKomponenGetRequest) (*wbk3.WbkKomponenGetInfoResponse, error)
}

type wbkKomponenService struct {
	//SpiAngRepository repository.SpiAng
	WbkKomponenRepository wbk.WbkKomponen
	Db                    *gorm.DB
}

func NewWbkKomponenService(f *factory.Factory) *wbkKomponenService {
	wbkKomponenRepository := f.WbkKomponenRepository

	db := f.Db
	return &wbkKomponenService{wbkKomponenRepository, db}

}

func (s *wbkKomponenService) Get(ctx *abstraction.Context,
	payload *wbk3.WbkKomponenGetRequest) (*wbk3.WbkKomponenGetInfoResponse, error) {

	var result *wbk3.WbkKomponenGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.WbkKomponenRepository.Find(ctx,
			&wbk2.WbkKomponenFilter{WbkKomponenEntityFilter: payload.WbkKomponenEntityFilter}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		result = &wbk3.WbkKomponenGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
