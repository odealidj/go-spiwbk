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

type MethodApbjService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.MethodApbjGetRequest) (*dto.MethodApbjGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type methodApbjService struct {
	Repository repository.MethodApbj
	Db         *gorm.DB
}

func NewMethodApbjService(f *factory.Factory) *methodApbjService {
	methodApbjRepository := f.MethodApbjRepository
	db := f.Db
	return &methodApbjService{methodApbjRepository, db}

}

func (s *methodApbjService) Get(ctx *abstraction.Context, payload *dto.MethodApbjGetRequest) (*dto.MethodApbjGetResponses, error) {
	var result *dto.MethodApbjGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		methodApbjs, info, err := s.Repository.Find(ctx, &payload.MethodApbjFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		methodApbjResponses := &[]dto.MethodApbjResponse{}
		for _, methodApbj := range methodApbjs {

			*methodApbjResponses = append(*methodApbjResponses,
				dto.MethodApbjResponse{
					ID:               abstraction.ID{ID: methodApbj.ID.ID},
					MethodApbjEntity: methodApbj.MethodApbjEntity,
				})
		}
		result = &dto.MethodApbjGetResponses{
			Datas:          methodApbjResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
