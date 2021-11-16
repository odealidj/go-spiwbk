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

type PegawaiService interface {
	Save(*abstraction.Context, *dto.PegawaiSaveRequest) (*dto.PegawaiResponse, error)
	//Update(*abstraction.Context, *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error)
	//Delete(*abstraction.Context, *dto.SatkerID) (*dto.SatkerResponse, error)
	Get(ctx *abstraction.Context, payload *dto.PegawaiGetRequest) (*dto.PegawaiGetResponse, error)
}

type pegawaiservice struct {
	Repository repository.Pegawai
	Db         *gorm.DB
}

func NewPegawaiservice(f *factory.Factory) *pegawaiservice {
	repository := f.PegawaiRepository
	db := f.Db
	return &pegawaiservice{repository, db}

}

func (s *pegawaiservice) Save(ctx *abstraction.Context, payload *dto.PegawaiSaveRequest) (*dto.PegawaiResponse, error) {

	var result *dto.PegawaiResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		pegawai, err := s.Repository.Create(ctx, &model.Pegawai{
			Context: ctx, PegawaiEntity: payload.PegawaiEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.PegawaiResponse{
			ID: abstraction.ID{ID: pegawai.ID}, PegawaiEntity: pegawai.PegawaiEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *pegawaiservice) Get(ctx *abstraction.Context, payload *dto.PegawaiGetRequest) (*dto.PegawaiGetResponse, error) {
	var result *dto.PegawaiGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		pegawaies, info, err := s.Repository.Find(ctx, &payload.PegawaiFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*pegawaies) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		pegawaiResponses := &[]dto.PegawaiResponse{}
		pegawaiResponse := &dto.PegawaiResponse{}
		for _, pegawai := range *pegawaies {
			pegawaiResponse.ID.ID = pegawai.ID
			pegawaiResponse.PegawaiEntity = pegawai.PegawaiEntity
			*pegawaiResponses = append(*pegawaiResponses, *pegawaiResponse)
		}
		result = &dto.PegawaiGetResponse{
			Datas:          pegawaiResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
