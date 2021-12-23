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

type SpiBmnService interface {
	Save(*abstraction.Context, *dto.SpiBmnSaveRequest) (*dto.SpiBmnResponse, error)
	//Upsert(*abstraction.Context, *dto.SpiPbjPaketJenisBelanjaPaguUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	GetSpiBmnByThnAngIDAndSatkerID(*abstraction.Context, *dto.SpiBmnGetRequest) (*dto.SpiBmnGetInfoResponse, error)
}

type spiBmnService struct {
	SpiAngRepository repository.SpiAng
	SpiBmnRepository repository.SpiBmn
	Db               *gorm.DB
}

func NewSpiBmnService(f *factory.Factory) *spiBmnService {
	spiAngRepository := f.SpiAngRepository
	spiBmnRepository := f.SpiBmnRepository
	db := f.Db
	return &spiBmnService{spiAngRepository, spiBmnRepository, db}

}

func (s *spiBmnService) Save(ctx *abstraction.Context, payload *dto.SpiBmnSaveRequest) (
	*dto.SpiBmnResponse, error) {

	var result *dto.SpiBmnResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiAng, err := s.SpiAngRepository.Create(ctx, &model.SpiAng{Context: ctx, SpiAngEntity: model.SpiAngEntity{
			ThnAngID: uint16(payload.ThnAngID), SatkerID: uint16(payload.SatkerID),
		}})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid spi ang")
		}

		payload.SpiAngID = int(spiAng.ID)
		spiBmn, err := s.SpiBmnRepository.Upsert(ctx, &model.SpiBmn{Context: ctx,
			SpiBmnEntity: payload.SpiBmnEntity})
		if err != nil {

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi bmn", err.Error())
		}

		result = &dto.SpiBmnResponse{
			ID:           abstraction.ID{ID: spiAng.ID},
			SpiBmnEntity: spiBmn.SpiBmnEntity,
		}

		//time.Sleep(time.Minute * 5)

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiBmnService) GetSpiBmnByThnAngIDAndSatkerID(ctx *abstraction.Context,
	payload *dto.SpiBmnGetRequest) (*dto.SpiBmnGetInfoResponse, error) {

	var result *dto.SpiBmnGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiBmnGetResponses, info, err := s.SpiBmnRepository.FindSpiBmnByThnAngIDAndSatkerID(ctx,
			&model.SpiBmnFilter{SpiBmnEntityFilter: model.SpiBmnEntityFilter{
				ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID},
			}, &payload.Pagination)
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Spi bmn", err.Error())
		}

		//num := 0
		for i, _ := range spiBmnGetResponses {
			spiBmnGetResponses[i].Row = i + 1
		}

		result = &dto.SpiBmnGetInfoResponse{
			Datas:          &spiBmnGetResponses,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
