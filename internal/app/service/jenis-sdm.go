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

type JenisSdmService interface {
	Save(*abstraction.Context, *dto.JenisSdmSaveRequest) (*dto.JenisSdmResponse, error)
	//Update(*abstraction.Context, *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error)
	//Delete(*abstraction.Context, *dto.SatkerID) (*dto.SatkerResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisSdmGetRequest) (*dto.JenisSdmGetResponse, error)
}

type jenissdmservice struct {
	Repository repository.JenisSdm
	Db         *gorm.DB
}

func NewJenisSdmSdmService(f *factory.Factory) *jenissdmservice {
	repository := f.JenisSdmRepository
	db := f.Db
	return &jenissdmservice{repository, db}

}

func (s *jenissdmservice) Save(ctx *abstraction.Context, payload *dto.JenisSdmSaveRequest) (*dto.JenisSdmResponse, error) {

	var result *dto.JenisSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenissdm, err := s.Repository.Create(ctx, &model.JenisSdm{
			Context:        ctx,
			JenisSdmEntity: payload.JenisSdmEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.JenisSdmResponse{
			ID:             abstraction.ID{ID: jenissdm.ID},
			JenisSdmEntity: jenissdm.JenisSdmEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *jenissdmservice) Get(ctx *abstraction.Context, payload *dto.JenisSdmGetRequest) (*dto.JenisSdmGetResponse, error) {
	var result *dto.JenisSdmGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenissdms, info, err := s.Repository.Find(ctx, &payload.JenisSdmFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*jenissdms) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		var jenisSdmResponses []dto.JenisSdmResponse
		for _, jenissdm := range *jenissdms {
			jenisSdmResponses = append(jenisSdmResponses, dto.JenisSdmResponse{
				ID:             abstraction.ID{ID: jenissdm.ID},
				JenisSdmEntity: jenissdm.JenisSdmEntity,
			})
		}
		result = &dto.JenisSdmGetResponse{
			Datas:          jenisSdmResponses,
			PaginationInfo: *info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
