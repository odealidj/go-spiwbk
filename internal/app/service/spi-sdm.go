package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"errors"
	"gorm.io/gorm"
	"strings"
)

type SpiSdmService interface {
	Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	//Update(*abstraction.Context, *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error)
	//Delete(*abstraction.Context, *dto.SatkerID) (*dto.SatkerResponse, error)
	Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error)
}

type spiSdmService struct {
	Repository repository.SpiSdm
	Db         *gorm.DB
}

func NewSpiSdmService(f *factory.Factory) *spiSdmService {
	repository := f.SpiSdmRepository
	db := f.Db
	return &spiSdmService{repository, db}

}

func (s *spiSdmService) Save(ctx *abstraction.Context, payload *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error) {

	var result *dto.SpiSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdm, err := s.Repository.Create(ctx, &model.SpiSdm{
			Context:      ctx,
			SpiSdmEntity: payload.SpiSdmEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmResponse{
			ID:           abstraction.ID{ID: spisdm.ID},
			SpiSdmEntity: spisdm.SpiSdmEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmService) Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error) {
	var result *dto.SpiSdmGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdms, info, err := s.Repository.Find(ctx, &payload.SpiSdmFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*spisdms) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		spiSdmResponses := &[]dto.SpiSdmResponse{}
		spiSdmResponse := &dto.SpiSdmResponse{}
		for _, spisdm := range *spisdms {
			spiSdmResponse.ID.ID = spisdm.ID
			spiSdmResponse.SpiSdmEntity = spisdm.SpiSdmEntity
			spiSdmResponse.ThnAngYear = spisdm.ThnAng.Year
			spiSdmResponse.SatkerName = spisdm.Satker.Name
			*spiSdmResponses = append(*spiSdmResponses, *spiSdmResponse)
		}
		result = &dto.SpiSdmGetResponse{
			Datas:          spiSdmResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
