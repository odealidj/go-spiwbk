package sample

import (
	"code-boiler/internal/abstractions"
	"code-boiler/internal/dto"
	"code-boiler/internal/model"
	"code-boiler/internal/repository"
	res "code-boiler/pkg/util/response"

	"gorm.io/gorm"
)

type Service interface {
	Find(*dto.SampleGetRequest) ([]*model.Sample, *abstractions.PaginationInfo, error)
	FindByID(id int) (*model.Sample, error)

	Create(data dto.SampleStoreRequest) (*model.Sample, error)
	Update(id int, data dto.SampleUpdateRequest) (*model.Sample, error)
	Delete(id int) (*model.Sample, error)

	WithTrx(trx *gorm.DB) *service
}

type service struct {
	Repository repository.Sample
}

func NewService(dbConnection *gorm.DB) *service {
	repository := repository.NewSample(dbConnection)
	return &service{repository}
}

func (s *service) WithTrx(trx *gorm.DB) *service {
	repository := repository.NewSample(trx)
	return &service{
		Repository: repository,
	}
}

func (s *service) Find(payload *dto.SampleGetRequest) ([]*model.Sample, *abstractions.PaginationInfo, error) {
	var companies []*model.Sample
	var err error

	companies, info, err := s.Repository.Find(payload)
	if err != nil {
		return nil, info, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	return companies, info, nil
}

func (s *service) FindByID(id int) (*model.Sample, error) {
	sample, err := s.Repository.FindByID(id)
	if err != nil {
		if err.Error() == "record not found" {
			return nil, res.ErrorBuilder(res.Constant.Error.NotFound, err)
		}
		return nil, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	return sample, nil
}

func (s *service) Create(data *dto.SampleStoreRequest) (*model.Sample, error) {
	sample, err := s.Repository.FindByKey(data.Key)
	if sample != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.Duplicate, nil)
	}

	if err != nil {
		if err.Error() != "record not found" {
			return nil, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
		}
	}

	sample = &model.Sample{
		Key:    data.Key,
		Value:  data.Value,
		UserId: data.UserId,
	}

	sample, err = s.Repository.Create(sample)
	if err != nil {
		return sample, res.ErrorBuilder(res.Constant.Error.UnprocessableEntity, err)
	}
	return sample, nil
}

func (s *service) Update(id int, data *dto.SampleUpdateRequest) (*model.Sample, error) {
	if data.Key != "" {
		sample, _ := s.Repository.FindByKey(data.Key)
		if sample != nil {
			return nil, res.ErrorBuilder(res.Constant.Error.Duplicate, nil)
		}
	}

	_, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.BadRequest, err)
	}
	sample, err := s.Repository.Update(id, &model.Sample{
		Key:    data.Key,
		Value:  data.Value,
		UserId: data.UserId,
	})
	if err != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	return sample, nil
}

func (s *service) Delete(id int) (*model.Sample, error) {
	_, err := s.Repository.FindByID(id)
	if err != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.BadRequest, err)
	}

	sample, err := s.Repository.Delete(id)
	if err != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}
	return sample, nil
}
