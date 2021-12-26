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

type ThnAngService interface {
	Save(*abstraction.Context, *dto.ThnAngRequest) (*dto.ThnAngResponse, error)
	SaveForm(*abstraction.Context, *dto.ThnAngRequestForm) (*dto.ThnAngResponse, error)
	SaveBatch(*abstraction.Context, []dto.ThnAngRequests) ([]dto.ThnAngResponse, error)
	Update(*abstraction.Context, *dto.ThnAngUpdateRequest) (*dto.ThnAngResponse, error)
	Delete(*abstraction.Context, *abstraction.ID) (*dto.ThnAngResponse, error)
	GetAll(*abstraction.Context, *dto.ThnAngGetAllRequest) (*dto.ThnAngGetAllResponse, error)
	Get(*abstraction.Context, *dto.ThnAngGetRequest) (*dto.ThnAngResponse, error)
	GetByYear(*abstraction.Context, *dto.ThnAngGetByYearRequest) (*dto.ThnAngResponse, error)
}

type thnAngService struct {
	Repository repository.ThnAng
	Db         *gorm.DB
}

func NewThnAngService(f *factory.Factory) *thnAngService {
	repository := f.ThnAngRepository
	db := f.Db
	return &thnAngService{repository, db}

}

func (s *thnAngService) Save(ctx *abstraction.Context, payload *dto.ThnAngRequest) (*dto.ThnAngResponse, error) {

	var result *dto.ThnAngResponse
	var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		data, err = s.Repository.Create(ctx, &model.ThnAng{Context: ctx, ThnAngEntity: payload.ThnAngEntity})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ThnAngResponse{
		ID: abstraction.ID{ID: data.EntityInc.ID}, ThnAngEntity: data.ThnAngEntity,
	}

	return result, nil

}

func (s *thnAngService) SaveForm(ctx *abstraction.Context, payload *dto.ThnAngRequestForm) (*dto.ThnAngResponse, error) {

	var result *dto.ThnAngResponse
	var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err = s.Repository.Create(ctx, &model.ThnAng{Context: ctx,
			ThnAngEntity: model.ThnAngEntity{
				Year: payload.Year,
			}})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.ThnAngResponse{
		ID: abstraction.ID{ID: data.EntityInc.ID}, ThnAngEntity: data.ThnAngEntity,
	}

	return result, nil

}

func (s *thnAngService) SaveBatch(ctx *abstraction.Context, payloads []dto.ThnAngRequests) ([]dto.ThnAngResponse, error) {

	var results []dto.ThnAngResponse
	var data []model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		var TngAngs []model.ThnAng
		//var ThnAngRequests []dto.ThnAngRequest
		for _, temp := range payloads {
			//ThnAngRequests = append(ThnAngRequests, temp)
			thnAgn := model.ThnAng{
				ThnAngEntity: model.ThnAngEntity{
					Year: temp.Year,
				},
			}
			TngAngs = append(TngAngs, thnAgn)
		}

		data, err = s.Repository.CreateBatch(ctx, TngAngs)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return results, err
	}

	for _, thn := range data {
		results = append(results, dto.ThnAngResponse{
			ID: abstraction.ID{ID: thn.ID}, ThnAngEntity: thn.ThnAngEntity,
		})
	}

	return results, nil

}

func (s *thnAngService) Update(ctx *abstraction.Context, payload *dto.ThnAngUpdateRequest) (*dto.ThnAngResponse, error) {

	var result *dto.ThnAngResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		thnAng, err := s.Repository.FindByID(ctx, &model.ThnAng{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{
				ID: payload.ID.ID,
			},
		}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		data, err := s.Repository.Update(ctx, &model.ThnAng{Context: ctx,
			EntityInc: abstraction.EntityInc{
				IDInc: abstraction.IDInc{
					ID: thnAng.EntityInc.IDInc.ID,
				},
			},
			ThnAngEntity: payload.ThnAngEntity})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.ThnAngResponse{
			ID: abstraction.ID{ID: data.EntityInc.ID}, ThnAngEntity: data.ThnAngEntity,
		}

		return nil
	}); err != nil {

		return result, err
	}

	return result, nil

}

func (s *thnAngService) Delete(ctx *abstraction.Context, payload *abstraction.ID) (*dto.ThnAngResponse, error) {

	var result *dto.ThnAngResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err := s.Repository.FindByID(ctx, &model.ThnAng{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID},
		}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		data, err = s.Repository.Delete(ctx, &model.ThnAng{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID},
		}})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.ThnAngResponse{
			ID: abstraction.ID{ID: data.EntityInc.ID}, ThnAngEntity: data.ThnAngEntity,
		}

		return nil
	}); err != nil {
		return result, err
	}

	return result, nil

}

func (s *thnAngService) GetAll(ctx *abstraction.Context, payload *dto.ThnAngGetAllRequest) (*dto.ThnAngGetAllResponse, error) {
	var result *dto.ThnAngGetAllResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		thnAngs, info, err := s.Repository.Find(ctx, &payload.ThnAngFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*thnAngs) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		thnAngResponses := &[]dto.ThnAngResponse{}
		thnAngResponse := &dto.ThnAngResponse{}
		for _, thnAng := range *thnAngs {
			thnAngResponse.ID.ID = thnAng.ID
			thnAngResponse.ThnAngEntity = thnAng.ThnAngEntity
			*thnAngResponses = append(*thnAngResponses, *thnAngResponse)
		}
		result = &dto.ThnAngGetAllResponse{
			Datas:          thnAngResponses,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return result, err
	}

	return result, nil

}

func (s *thnAngService) Get(ctx *abstraction.Context, payload *dto.ThnAngGetRequest) (*dto.ThnAngResponse, error) {
	var result *dto.ThnAngResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		thnang, err := s.Repository.FindByID(ctx, &model.ThnAng{Context: ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.ThnAngResponse{
			ID:           abstraction.ID{ID: thnang.ID},
			ThnAngEntity: thnang.ThnAngEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *thnAngService) GetByYear(ctx *abstraction.Context, payload *dto.ThnAngGetByYearRequest) (*dto.ThnAngResponse, error) {
	var result *dto.ThnAngResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		thnang, err := s.Repository.FindByYear(ctx, &model.ThnAng{Context: ctx,
			ThnAngEntity: model.ThnAngEntity{Year: payload.Year}})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		if thnang.ID == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}

		result = &dto.ThnAngResponse{
			ID:           abstraction.ID{ID: thnang.ID},
			ThnAngEntity: thnang.ThnAngEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}
