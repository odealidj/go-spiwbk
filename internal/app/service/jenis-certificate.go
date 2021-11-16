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

type JenisCertificateService interface {
	Save(*abstraction.Context, *dto.JenisCertificateSaveRequest) (*dto.JenisCertificateResponse, error)
	//Update(*abstraction.Context, *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error)
	//Delete(*abstraction.Context, *dto.SatkerID) (*dto.SatkerResponse, error)
	Get(ctx *abstraction.Context, payload *dto.JenisCertificateGetRequest) (*dto.JenisCertificateGetResponse, error)
}

type jenisCertificateService struct {
	Repository repository.JenisCertificate
	Db         *gorm.DB
}

func NewJenisCertificateService(f *factory.Factory) *jenisCertificateService {
	repository := f.JenisCertificateRepository
	db := f.Db
	return &jenisCertificateService{repository, db}

}

func (s *jenisCertificateService) Save(ctx *abstraction.Context, payload *dto.JenisCertificateSaveRequest) (*dto.JenisCertificateResponse, error) {

	var result *dto.JenisCertificateResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenissdm, err := s.Repository.Create(ctx, &model.JenisCertificate{
			Context:                ctx,
			JenisCertificateEntity: payload.JenisCertificateEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.JenisCertificateResponse{
			ID:                     abstraction.ID{ID: jenissdm.ID},
			JenisCertificateEntity: jenissdm.JenisCertificateEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *jenisCertificateService) Get(ctx *abstraction.Context, payload *dto.JenisCertificateGetRequest) (*dto.JenisCertificateGetResponse, error) {
	var result *dto.JenisCertificateGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		jenisCertificates, info, err := s.Repository.Find(ctx, &payload.JenisCertificateFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*jenisCertificates) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		var jenisCertificateResponses []dto.JenisCertificateResponse
		for _, jenisCertificate := range *jenisCertificates {
			jenisCertificateResponses = append(jenisCertificateResponses, dto.JenisCertificateResponse{
				ID:                     abstraction.ID{ID: jenisCertificate.ID},
				JenisCertificateEntity: jenisCertificate.JenisCertificateEntity,
			})
		}
		result = &dto.JenisCertificateGetResponse{
			Datas:          jenisCertificateResponses,
			PaginationInfo: *info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
