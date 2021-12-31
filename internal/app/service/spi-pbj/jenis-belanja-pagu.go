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

type JenisBelanjaPaguService interface {
	//Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *spi_pbj2.JenisBelanjaPaguGetRequest) (*spi_pbj2.JenisBelanjaPaguGetResponses, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type jenisBelanjaPaguService struct {
	Repository spi_pbj.JenisBelanjaPagu
	Db         *gorm.DB
}

func NewJenisBelanjaPaguService(f *factory.Factory) *jenisBelanjaPaguService {
	jenisBelanjaPaguRepository := f.JenisBelanjaPaguRepository
	db := f.Db
	return &jenisBelanjaPaguService{jenisBelanjaPaguRepository, db}

}

func (s *jenisBelanjaPaguService) Get(ctx *abstraction.Context, payload *spi_pbj2.JenisBelanjaPaguGetRequest) (*spi_pbj2.JenisBelanjaPaguGetResponses, error) {
	var result *spi_pbj2.JenisBelanjaPaguGetResponses

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisBelanjaPagues, info, err := s.Repository.Find(ctx, &payload.JenisBelanjaPaguFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		jenisBelanjaPaguResponses := &[]spi_pbj2.JenisBelanjaPaguResponse{}
		for _, jenisBelanjaPagu := range jenisBelanjaPagues {

			*jenisBelanjaPaguResponses = append(*jenisBelanjaPaguResponses,
				spi_pbj2.JenisBelanjaPaguResponse{
					ID:                     abstraction.ID{ID: jenisBelanjaPagu.ID.ID},
					JenisBelanjaPaguEntity: jenisBelanjaPagu.JenisBelanjaPaguEntity,
				})
		}
		result = &spi_pbj2.JenisBelanjaPaguGetResponses{
			Datas:          jenisBelanjaPaguResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
