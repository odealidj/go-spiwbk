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

type GroupPackageValueService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.GroupPackageValueGetRequest) (*dto.GroupPackageValueGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type groupPackageValueService struct {
	Repository repository.GroupPackageValue
	Db         *gorm.DB
}

func NewGroupPackageValueService(f *factory.Factory) *groupPackageValueService {
	groupPackageValueRepository := f.GroupPackageValueRepository
	db := f.Db
	return &groupPackageValueService{groupPackageValueRepository, db}

}

func (s *groupPackageValueService) Get(ctx *abstraction.Context, payload *dto.GroupPackageValueGetRequest) (*dto.GroupPackageValueGetResponses, error) {
	var result *dto.GroupPackageValueGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		groupPackageValues, info, err := s.Repository.Find(ctx, &payload.GroupPackageValueFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		groupPackageValueResponse := &[]dto.GroupPackageValueResponse{}
		for _, groupPackageValue := range groupPackageValues {

			*groupPackageValueResponse = append(*groupPackageValueResponse,
				dto.GroupPackageValueResponse{
					ID:                      abstraction.ID{ID: groupPackageValue.ID.ID},
					GroupPackageValueEntity: groupPackageValue.GroupPackageValueEntity,
				})
		}
		result = &dto.GroupPackageValueGetResponses{
			Datas:          groupPackageValueResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
