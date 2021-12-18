package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
)

type JenisRekapitulasiService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisRekapitulasiGetRequest) (*dto.JenisRekapitulasiGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisRekapitulasiService struct {
	Repository repository.JenisRekapitulasi
	Db         *gorm.DB
}

func NewJenisRekapitulasiService(f *factory.Factory) *jenisRekapitulasiService {
	jenisRekapitulasiRepository := f.JenisRekapitulasiRepository
	db := f.Db
	return &jenisRekapitulasiService{jenisRekapitulasiRepository, db}

}

func (s *jenisRekapitulasiService) Get(ctx *abstraction.Context, payload *dto.JenisRekapitulasiGetRequest) (*dto.JenisRekapitulasiGetResponses, error) {
	var result *dto.JenisRekapitulasiGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisRekapitulasies, info, err := s.Repository.Find(ctx, &payload.JenisRekapitulasiFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisRekapitulasiResponse := &[]dto.JenisRekapitulasiResponse{}
		for _, jenisRekapitulasi := range jenisRekapitulasies {

			*jenisRekapitulasiResponse = append(*jenisRekapitulasiResponse,
				dto.JenisRekapitulasiResponse{
					ID:                      abstraction.ID{ID: jenisRekapitulasi.ID.ID},
					JenisRekapitulasiEntity: jenisRekapitulasi.JenisRekapitulasiEntity,
				})
		}
		result = &dto.JenisRekapitulasiGetResponses{
			Datas:          jenisRekapitulasiResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
