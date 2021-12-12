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

type SpiAngService interface {
	Save(*abstraction.Context, *dto.SpiAngSaveRequest) (*dto.SpiAngResponse, error)
	Update(*abstraction.Context, *dto.SpiAngUpdateRequest) (*dto.SpiAngResponse, error)
	Delete(*abstraction.Context, *dto.SpiAngDeleteRequest) (*dto.SpiAngResponse, error)
	Get(ctx *abstraction.Context, payload *dto.SpiAngGetRequest) (*dto.SpiAngGetResponse, error)
	GetByID(*abstraction.Context, *dto.SpiAngGetByIDRequest) (*dto.SpiAngResponse, error)
}

type spiAngService struct {
	Repository repository.SpiAng
	Db         *gorm.DB
}

func NewSpiAngService(f *factory.Factory) *spiAngService {
	repository := f.SpiAngRepository
	db := f.Db
	return &spiAngService{repository, db}

}

func (s *spiAngService) Save(ctx *abstraction.Context, payload *dto.SpiAngSaveRequest) (*dto.SpiAngResponse, error) {

	var result *dto.SpiAngResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdm, err := s.Repository.Create(ctx, &model.SpiAng{
			Context:      ctx,
			SpiAngEntity: payload.SpiAngEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiAngResponse{
			ID:           abstraction.ID{ID: spisdm.ID},
			SpiAngEntity: spisdm.SpiAngEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiAngService) Update(ctx *abstraction.Context, payload *dto.SpiAngUpdateRequest) (*dto.SpiAngResponse, error) {

	var result *dto.SpiAngResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiAng, err := s.Repository.FindByID(ctx, &model.SpiAng{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiAng, err = s.Repository.Update(ctx, &model.SpiAng{
			Context:      ctx,
			EntityInc:    abstraction.EntityInc{IDInc: abstraction.IDInc{ID: spiAng.ID}},
			SpiAngEntity: payload.SpiAngEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiAngResponse{
			ID:           abstraction.ID{ID: spiAng.ID},
			SpiAngEntity: spiAng.SpiAngEntity,
			//ThnAngYear:   spiSdm.ThnAng.Year,
			//SatkerName:   spiSdm.Satker.Name,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiAngService) Delete(ctx *abstraction.Context, payload *dto.SpiAngDeleteRequest) (*dto.SpiAngResponse, error) {

	var result *dto.SpiAngResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiAng, err := s.Repository.FindByID(ctx, &model.SpiAng{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiAng, err = s.Repository.Delete(ctx, &model.SpiAng{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: spiAng.ID}},
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiAngResponse{
			ID: abstraction.ID{ID: spiAng.ID},
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiAngService) Get(ctx *abstraction.Context, payload *dto.SpiAngGetRequest) (*dto.SpiAngGetResponse, error) {
	var result *dto.SpiAngGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spiAngs, info, err := s.Repository.Find(ctx, &payload.SpiAngFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		//if len(*spisdms) == 0 {
		//	return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		//}

		spiAngResponses := &[]dto.SpiAngResponses{}
		spiAngResponse := &dto.SpiAngResponses{}
		for _, spiang := range *spiAngs {
			spiAngResponse.ID.ID = spiang.ID
			spiAngResponse.SpiAngEntity = spiang.SpiAngEntity
			spiAngResponse.ThnAngYear = spiang.ThnAng.Year
			spiAngResponse.SatkerName = spiang.Satker.Name
			*spiAngResponses = append(*spiAngResponses, *spiAngResponse)
		}
		result = &dto.SpiAngGetResponse{
			Datas:          spiAngResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *spiAngService) GetByID(ctx *abstraction.Context, payload *dto.SpiAngGetByIDRequest) (*dto.SpiAngResponse, error) {
	var result *dto.SpiAngResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spi, err := s.Repository.FindByID(ctx, &model.SpiAng{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		result = &dto.SpiAngResponse{
			ID:           abstraction.ID{ID: spi.ID},
			SpiAngEntity: spi.SpiAngEntity,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
