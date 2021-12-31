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

type GroupPackageValueService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *spi_pbj2.GroupPackageValueGetRequest) (*spi_pbj2.GroupPackageValueGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type groupPackageValueService struct {
	Repository spi_pbj.GroupPackageValue
	Db         *gorm.DB
}

func NewGroupPackageValueService(f *factory.Factory) *groupPackageValueService {
	groupPackageValueRepository := f.GroupPackageValueRepository
	db := f.Db
	return &groupPackageValueService{groupPackageValueRepository, db}

}

func (s *groupPackageValueService) Get(ctx *abstraction.Context, payload *spi_pbj2.GroupPackageValueGetRequest) (*spi_pbj2.GroupPackageValueGetResponses, error) {
	var result *spi_pbj2.GroupPackageValueGetResponses

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		groupPackageValues, info, err := s.Repository.Find(ctx, &payload.GroupPackageValueFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		groupPackageValueResponse := &[]spi_pbj2.GroupPackageValueResponse{}
		for _, groupPackageValue := range groupPackageValues {

			*groupPackageValueResponse = append(*groupPackageValueResponse,
				spi_pbj2.GroupPackageValueResponse{
					ID:                      abstraction.ID{ID: groupPackageValue.ID.ID},
					GroupPackageValueEntity: groupPackageValue.GroupPackageValueEntity,
				})
		}
		result = &spi_pbj2.GroupPackageValueGetResponses{
			Datas:          groupPackageValueResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
