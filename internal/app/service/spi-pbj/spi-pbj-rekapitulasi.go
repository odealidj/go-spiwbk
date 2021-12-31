package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	spi_pbj3 "codeid-boiler/internal/app/dto/spi-pbj"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/model/spi-pbj"
	"codeid-boiler/internal/app/repository"
	spi_pbj2 "codeid-boiler/internal/app/repository/spi-pbj"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
	"net/http"
)

type SpiPbjRekapitulasiService interface {
	Save(*abstraction.Context, *spi_pbj3.SpiPbjRekapitulasiSaveRequest) ([]spi_pbj3.SpiPbjRekapitulasiResponse, error)
	Upsert(*abstraction.Context, *spi_pbj3.SpiPbjRekapitulasiUpsertRequest) ([]spi_pbj3.SpiPbjRekapitulasiResponse, error)
	GetSpiPbjRekapitulasiByID(*abstraction.Context, *spi_pbj3.SpiPbjRekapitulasiGetRequest) (*spi_pbj3.SpiPbjRekapitulasiGetInfoResponse, error)
}

type spiPbjRekapitulasiService struct {
	SpiAngRepository             repository.SpiAng
	SpiPbjRekapitulasiRepository spi_pbj2.SpiPbjRekapitulasi
	JenisRekapitulasiRepository  spi_pbj2.JenisRekapitulasi
	BulanRepository              repository.Bulan
	Db                           *gorm.DB
}

func NewSpiPbjRekapitulasiService(f *factory.Factory) *spiPbjRekapitulasiService {

	spiPbjRekapitulasiRepository := f.SpiPbjRekapitulasiRepository
	spiAngRepository := f.SpiAngRepository
	jenisRekapitulasiRepository := f.JenisRekapitulasiRepository
	bulanRepository := f.BulanRepository
	db := f.Db
	return &spiPbjRekapitulasiService{spiAngRepository,
		spiPbjRekapitulasiRepository, jenisRekapitulasiRepository,
		bulanRepository, db}

}

func (s *spiPbjRekapitulasiService) Save(ctx *abstraction.Context,
	payload *spi_pbj3.SpiPbjRekapitulasiSaveRequest) ([]spi_pbj3.SpiPbjRekapitulasiResponse, error) {

	var result []spi_pbj3.SpiPbjRekapitulasiResponse
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

		jenisRekapitulasies, _, err := s.JenisRekapitulasiRepository.Find(ctx, &spi_pbj.JenisRekapitulasiFilter{},
			&abstraction.Pagination{})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Jenis rekapitulasi", err.Error())
		}

		bulans, _, err := s.BulanRepository.Find(ctx, &model.BulanFilter{}, &abstraction.Pagination{})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid Bulan", err.Error())
		}

		for _, jenisRekapitulasi := range jenisRekapitulasies {

			for _, bulan := range bulans {

				spiPbjRekapitulasi, err := s.SpiPbjRekapitulasiRepository.Create(ctx, &spi_pbj.SpiPbjRekapitulasi{Context: ctx,
					SpiPbjRekapitulasiEntity: spi_pbj.SpiPbjRekapitulasiEntity{
						SpiAngID: int(spiAng.ID), JenisRekapitulasiID: int(jenisRekapitulasi.ID.ID),
						BulanID: int(bulan.ID.ID), Target: 0.0,
					}})
				if err != nil {
					return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
						"Invalid spi pbj rekapitulasi", err.Error())
				}

				result = append(result, spi_pbj3.SpiPbjRekapitulasiResponse{
					ID:                       abstraction.ID{ID: spiPbjRekapitulasi.ID},
					SpiPbjRekapitulasiEntity: spiPbjRekapitulasi.SpiPbjRekapitulasiEntity,
					SatkerID:                 payload.SatkerID,
					ThnAngID:                 payload.ThnAngID,
				})

			} //end for bulan

		} //end for jenis rekapitulasi

		/*
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
		*/

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiPbjRekapitulasiService) Upsert(ctx *abstraction.Context, payload *spi_pbj3.SpiPbjRekapitulasiUpsertRequest) ([]spi_pbj3.SpiPbjRekapitulasiResponse, error) {

	var result []spi_pbj3.SpiPbjRekapitulasiResponse
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
		spiPbjRekapitulasi, err := s.SpiPbjRekapitulasiRepository.Upsert(ctx, &spi_pbj.SpiPbjRekapitulasi{Context: ctx,
			IDInc:                    abstraction.IDInc{ID: payload.ID.ID},
			SpiPbjRekapitulasiEntity: payload.SpiPbjRekapitulasiEntity})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid upsert spi pbj rekapitulasi", err.Error())
		}

		result = append(result, spi_pbj3.SpiPbjRekapitulasiResponse{
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

func (s *spiPbjRekapitulasiService) GetSpiPbjRekapitulasiByID(ctx *abstraction.Context, payload *spi_pbj3.SpiPbjRekapitulasiGetRequest) (
	*spi_pbj3.SpiPbjRekapitulasiGetInfoResponse, error) {

	var result *spi_pbj3.SpiPbjRekapitulasiGetInfoResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
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
		result = &spi_pbj3.SpiPbjRekapitulasiGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
