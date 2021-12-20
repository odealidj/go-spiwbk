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

type SpiPbjRekapitulasiService interface {
	Save(*abstraction.Context, *dto.SpiPbjRekapitulasiSaveRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	Upsert(*abstraction.Context, *dto.SpiPbjRekapitulasiUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error)
	GetSpiPbjRekapitulasiByID(*abstraction.Context, *dto.SpiPbjRekapitulasiGetRequest) (*dto.SpiPbjRekapitulasiGetInfoResponse, error)
}

type spiPbjRekapitulasiService struct {
	SpiAngRepository             repository.SpiAng
	SpiPbjRekapitulasiRepository repository.SpiPbjRekapitulasi
	Db                           *gorm.DB
}

func NewSpiPbjRekapitulasiService(f *factory.Factory) *spiPbjRekapitulasiService {

	spiPbjRekapitulasiRepository := f.SpiPbjRekapitulasiRepository
	spiAngRepository := f.SpiAngRepository
	db := f.Db
	return &spiPbjRekapitulasiService{spiAngRepository, spiPbjRekapitulasiRepository, db}

}

func (s *spiPbjRekapitulasiService) Save(ctx *abstraction.Context, payload *dto.SpiPbjRekapitulasiSaveRequest) ([]dto.SpiPbjRekapitulasiResponse, error) {

	var result []dto.SpiPbjRekapitulasiResponse
	//var data *model.ThnAng

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

		payload.SpiPbjRekapitulasiEntity.SpiAngID = int(spiAng.ID)
		spiPbjRekapitulasi, err := s.SpiPbjRekapitulasiRepository.Create(ctx, &model.SpiPbjRekapitulasi{Context: ctx,
			//IDInc:                    abstraction.IDInc{ID: payload.ID.ID},
			SpiPbjRekapitulasiEntity: payload.SpiPbjRekapitulasiEntity})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid upsert spi pbj rekapitulasi", err.Error())
		}

		result = append(result, dto.SpiPbjRekapitulasiResponse{
			ID:                       abstraction.ID{ID: spiPbjRekapitulasi.ID},
			SpiPbjRekapitulasiEntity: spiPbjRekapitulasi.SpiPbjRekapitulasiEntity,
			SatkerID:                 payload.SatkerID,
			ThnAngID:                 payload.ThnAngID,
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiPbjRekapitulasiService) Upsert(ctx *abstraction.Context, payload *dto.SpiPbjRekapitulasiUpsertRequest) ([]dto.SpiPbjRekapitulasiResponse, error) {

	var result []dto.SpiPbjRekapitulasiResponse
	//var data *model.ThnAng

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

		payload.SpiPbjRekapitulasiEntity.SpiAngID = int(spiAng.ID)
		spiPbjRekapitulasi, err := s.SpiPbjRekapitulasiRepository.Upsert(ctx, &model.SpiPbjRekapitulasi{Context: ctx,
			IDInc:                    abstraction.IDInc{ID: payload.ID.ID},
			SpiPbjRekapitulasiEntity: payload.SpiPbjRekapitulasiEntity})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid upsert spi pbj rekapitulasi", err.Error())
		}

		result = append(result, dto.SpiPbjRekapitulasiResponse{
			ID:                       abstraction.ID{ID: spiPbjRekapitulasi.ID},
			SpiPbjRekapitulasiEntity: spiPbjRekapitulasi.SpiPbjRekapitulasiEntity,
			SatkerID:                 payload.SatkerID,
			ThnAngID:                 payload.ThnAngID,
		})

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiPbjRekapitulasiService) GetSpiPbjRekapitulasiByID(ctx *abstraction.Context, payload *dto.SpiPbjRekapitulasiGetRequest) (
	*dto.SpiPbjRekapitulasiGetInfoResponse, error) {

	var result *dto.SpiPbjRekapitulasiGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		datas, info, err := s.SpiPbjRekapitulasiRepository.
			FindSpiPbjRekapitulasiByID(ctx, &payload.SpiPbjRekapitulasiFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		for i, _ := range datas {
			datas[i].Row = i + 1
		}

		/*
			if len(dealers) <= 0 {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data not found"))
			}
		*/
		result = &dto.SpiPbjRekapitulasiGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
