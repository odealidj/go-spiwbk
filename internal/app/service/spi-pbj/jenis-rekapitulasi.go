package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	spi_pbj2 "codeid-boiler/internal/app/dto/spi-pbj"
	"codeid-boiler/internal/app/repository/spi-pbj"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
)

type JenisRekapitulasiService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *spi_pbj2.JenisRekapitulasiGetRequest) (*spi_pbj2.JenisRekapitulasiGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisRekapitulasiService struct {
	Repository spi_pbj.JenisRekapitulasi
	Db         *gorm.DB
}

func NewJenisRekapitulasiService(f *factory.Factory) *jenisRekapitulasiService {
	jenisRekapitulasiRepository := f.JenisRekapitulasiRepository
	db := f.Db
	return &jenisRekapitulasiService{jenisRekapitulasiRepository, db}

}

func (s *jenisRekapitulasiService) Get(ctx *abstraction.Context, payload *spi_pbj2.JenisRekapitulasiGetRequest) (*spi_pbj2.JenisRekapitulasiGetResponses, error) {
	var result *spi_pbj2.JenisRekapitulasiGetResponses

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisRekapitulasies, info, err := s.Repository.Find(ctx, &payload.JenisRekapitulasiFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisRekapitulasiResponse := &[]spi_pbj2.JenisRekapitulasiResponse{}
		for _, jenisRekapitulasi := range jenisRekapitulasies {

			*jenisRekapitulasiResponse = append(*jenisRekapitulasiResponse,
				spi_pbj2.JenisRekapitulasiResponse{
					ID:                      abstraction.ID{ID: jenisRekapitulasi.ID.ID},
					JenisRekapitulasiEntity: jenisRekapitulasi.JenisRekapitulasiEntity,
				})
		}
		result = &spi_pbj2.JenisRekapitulasiGetResponses{
			Datas:          jenisRekapitulasiResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
