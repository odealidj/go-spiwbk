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

type MethodApbjService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *spi_pbj2.MethodApbjGetRequest) (*spi_pbj2.MethodApbjGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type methodApbjService struct {
	Repository spi_pbj.MethodApbj
	Db         *gorm.DB
}

func NewMethodApbjService(f *factory.Factory) *methodApbjService {
	methodApbjRepository := f.MethodApbjRepository
	db := f.Db
	return &methodApbjService{methodApbjRepository, db}

}

func (s *methodApbjService) Get(ctx *abstraction.Context, payload *spi_pbj2.MethodApbjGetRequest) (*spi_pbj2.MethodApbjGetResponses, error) {
	var result *spi_pbj2.MethodApbjGetResponses

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		methodApbjs, info, err := s.Repository.Find(ctx, &payload.MethodApbjFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		methodApbjResponses := &[]spi_pbj2.MethodApbjResponse{}
		for _, methodApbj := range methodApbjs {

			*methodApbjResponses = append(*methodApbjResponses,
				spi_pbj2.MethodApbjResponse{
					ID:               abstraction.ID{ID: methodApbj.ID.ID},
					MethodApbjEntity: methodApbj.MethodApbjEntity,
				})
		}
		result = &spi_pbj2.MethodApbjGetResponses{
			Datas:          methodApbjResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
