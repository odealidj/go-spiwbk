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

type JenisPengendaliService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisPengendaliGetRequest) (*dto.JenisPengendaliGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisPengendaliService struct {
	Repository repository.JenisPengendali
	Db         *gorm.DB
}

func NewJenisPengendaliService(f *factory.Factory) *jenisPengendaliService {
	repository := f.JenisPengendaliRepository
	db := f.Db
	return &jenisPengendaliService{repository, db}

}

func (s *jenisPengendaliService) Get(ctx *abstraction.Context, payload *dto.JenisPengendaliGetRequest) (*dto.JenisPengendaliGetResponses, error) {
	var result *dto.JenisPengendaliGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		datas, info, err := s.Repository.Find(ctx, &payload.JenisPengendaliFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisPengendaliResponse := &[]dto.JenisPengendaliResponse{}
		for _, data := range datas {

			*jenisPengendaliResponse = append(*jenisPengendaliResponse,
				dto.JenisPengendaliResponse{
					ID:                    abstraction.ID{ID: data.ID.ID},
					JenisPengendaliEntity: data.JenisPengendaliEntity,
				})
		}
		result = &dto.JenisPengendaliGetResponses{
			Datas:          jenisPengendaliResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
