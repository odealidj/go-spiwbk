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

type SatkerService interface {
	Save(*abstraction.Context, *dto.SatkerSaveRequest) (*dto.SatkerResponse, error)
	Update(*abstraction.Context, *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error)
	Delete(*abstraction.Context, *dto.SatkerDeleteRequest) (*dto.SatkerResponse, error)
	Get(ctx *abstraction.Context, payload *dto.SatkerGetRequest) (*dto.SatkerGetResponse, error)
	Get2(ctx *abstraction.Context, payload *dto.SatkerGet2Request) (*dto.SatkerGet2Response, error)
	GetByID(*abstraction.Context, *dto.SatkerGetByIDRequest) (*dto.SatkerResponse, error)
	GetCount(ctx *abstraction.Context) (*int64, error)
}

type satkerService struct {
	Repository repository.Satker
	Db         *gorm.DB
}

func NewSatkerService(f *factory.Factory) *satkerService {
	repository := f.SatkerRepository
	db := f.Db
	return &satkerService{repository, db}

}

func (s *satkerService) Save(ctx *abstraction.Context, payload *dto.SatkerSaveRequest) (*dto.SatkerResponse, error) {

	var result *dto.SatkerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		satker, err := s.Repository.Create(ctx, &model.Satker{
			Context:      ctx,
			SatkerEntity: payload.SatkerEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SatkerResponse{
			ID:           abstraction.ID{ID: satker.ID},
			SatkerEntity: satker.SatkerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *satkerService) Update(ctx *abstraction.Context, payload *dto.SatkerUpdateRequest) (*dto.SatkerResponse, error) {

	var result *dto.SatkerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		satker, err := s.Repository.FindByID(ctx, &model.Satker{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		satker, err = s.Repository.Update(ctx, &model.Satker{
			Context:      ctx,
			EntityInc:    abstraction.EntityInc{IDInc: abstraction.IDInc{ID: satker.ID}},
			SatkerEntity: payload.SatkerEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SatkerResponse{
			ID:           abstraction.ID{ID: satker.ID},
			SatkerEntity: satker.SatkerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *satkerService) Delete(ctx *abstraction.Context, payload *dto.SatkerDeleteRequest) (*dto.SatkerResponse, error) {

	var result *dto.SatkerResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		satker, err := s.Repository.FindByID(ctx, &model.Satker{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		satker, err = s.Repository.Delete(ctx, &model.Satker{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: satker.ID}},
		})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SatkerResponse{
			ID:           abstraction.ID{ID: satker.ID},
			SatkerEntity: satker.SatkerEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *satkerService) Get(ctx *abstraction.Context, payload *dto.SatkerGetRequest) (*dto.SatkerGetResponse, error) {
	var result *dto.SatkerGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.Repository.Find(ctx, payload.SatkerFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*datas) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		var satkerResponse []dto.SatkerResponse
		for _, satker := range *datas {
			satkerResponse = append(satkerResponse, dto.SatkerResponse{
				ID:           abstraction.ID{ID: satker.ID},
				SatkerEntity: satker.SatkerEntity,
			})
		}

		result = &dto.SatkerGetResponse{
			Datas:          satkerResponse,
			PaginationInfo: *info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *satkerService) Get2(ctx *abstraction.Context, payload *dto.SatkerGet2Request) (*dto.SatkerGet2Response, error) {
	var result *dto.SatkerGet2Response

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		datas, info, err := s.Repository.Find2(ctx, payload.SatkerFilter, &payload.PaginationArr)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*datas) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		var satkerResponse []dto.SatkerResponse
		for _, satker := range *datas {
			satkerResponse = append(satkerResponse, dto.SatkerResponse{
				ID:           abstraction.ID{ID: satker.ID},
				SatkerEntity: satker.SatkerEntity,
			})
		}

		result = &dto.SatkerGet2Response{
			Datas:          satkerResponse,
			PaginationInfo: *info,
		}
		return nil
	}); err != nil {
		result = &dto.SatkerGet2Response{}
		return result, err
	}

	return result, nil
}

func (s *satkerService) GetByID(ctx *abstraction.Context, payload *dto.SatkerGetByIDRequest) (*dto.SatkerResponse, error) {
	var result *dto.SatkerResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		satker, err := s.Repository.FindByID(ctx, &model.Satker{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		result = &dto.SatkerResponse{
			ID:           abstraction.ID{ID: satker.ID},
			SatkerEntity: satker.SatkerEntity,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *satkerService) GetCount(ctx *abstraction.Context) (*int64, error) {
	var result *int64
	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		count, err := s.Repository.Count(ctx)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = count

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}
