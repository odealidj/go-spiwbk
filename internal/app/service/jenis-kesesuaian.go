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

type JenisKesesuaianService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisKesesuaianGetRequest) (*dto.JenisKesesuaianGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisKesesuaianService struct {
	Repository repository.JenisKesesuaian
	Db         *gorm.DB
}

func NewJenisKesesuaianService(f *factory.Factory) *jenisKesesuaianService {
	repository := f.JenisKesesuaianRepository
	db := f.Db
	return &jenisKesesuaianService{repository, db}

}

func (s *jenisKesesuaianService) Get(ctx *abstraction.Context, payload *dto.JenisKesesuaianGetRequest) (*dto.JenisKesesuaianGetResponses, error) {
	var result *dto.JenisKesesuaianGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisKesesuaians, info, err := s.Repository.Find(ctx, &payload.JenisKesesuaianFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisKesesuaianResponse := &[]dto.JenisKesesuaianResponse{}
		for _, jenisKesesuaian := range jenisKesesuaians {

			*jenisKesesuaianResponse = append(*jenisKesesuaianResponse,
				dto.JenisKesesuaianResponse{
					ID:                    abstraction.ID{ID: jenisKesesuaian.ID.ID},
					JenisKesesuaianEntity: jenisKesesuaian.JenisKesesuaianEntity,
				})
		}
		result = &dto.JenisKesesuaianGetResponses{
			Datas:          jenisKesesuaianResponse,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
