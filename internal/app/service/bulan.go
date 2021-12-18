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

type BulanService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, request *dto.BulanGetRequest) (*dto.BulanGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type bulanService struct {
	Repository repository.Bulan
	Db         *gorm.DB
}

func NewBulanService(f *factory.Factory) *bulanService {
	bulanRepository := f.BulanRepository
	db := f.Db
	return &bulanService{bulanRepository, db}

}

func (s *bulanService) Get(ctx *abstraction.Context, payload *dto.BulanGetRequest) (*dto.BulanGetResponses, error) {
	var result *dto.BulanGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		bulans, info, err := s.Repository.Find(ctx, &payload.BulanFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		bulanResponse := &[]dto.BulanResponse{}
		for _, bulan := range bulans {

			*bulanResponse = append(*bulanResponse,
				dto.BulanResponse{
					ID:          abstraction.ID{ID: bulan.ID.ID},
					BulanEntity: bulan.BulanEntity,
				})
		}
		result = &dto.BulanGetResponses{
			Datas:          bulanResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
