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

type JenisBelanjaPaguService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisBelanjaPaguGetRequest) (*dto.JenisBelanjaPaguGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisBelanjaPaguService struct {
	Repository repository.JenisBelanjaPagu
	Db         *gorm.DB
}

func NewJenisBelanjaPaguService(f *factory.Factory) *jenisBelanjaPaguService {
	jenisBelanjaPaguRepository := f.JenisBelanjaPaguRepository
	db := f.Db
	return &jenisBelanjaPaguService{jenisBelanjaPaguRepository, db}

}

func (s *jenisBelanjaPaguService) Get(ctx *abstraction.Context, payload *dto.JenisBelanjaPaguGetRequest) (*dto.JenisBelanjaPaguGetResponses, error) {
	var result *dto.JenisBelanjaPaguGetResponses

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisBelanjaPagues, info, err := s.Repository.Find(ctx, &payload.JenisBelanjaPaguFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisBelanjaPaguResponses := &[]dto.JenisBelanjaPaguResponse{}
		for _, jenisBelanjaPagu := range jenisBelanjaPagues {

			*jenisBelanjaPaguResponses = append(*jenisBelanjaPaguResponses,
				dto.JenisBelanjaPaguResponse{
					ID:                     abstraction.ID{ID: jenisBelanjaPagu.ID.ID},
					JenisBelanjaPaguEntity: jenisBelanjaPagu.JenisBelanjaPaguEntity,
				})
		}
		result = &dto.JenisBelanjaPaguGetResponses{
			Datas:          jenisBelanjaPaguResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
