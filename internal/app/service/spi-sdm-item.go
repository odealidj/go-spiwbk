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

type SpiSdmItemService interface {
	Save(*abstraction.Context, *dto.SpiSdmItemSaveRequest) (*dto.SpiSdmItemResponse, error)
	Update(*abstraction.Context, *dto.SpiSdmItemUpdateRequest) (*dto.SpiSdmItemResponse, error)
	Delete(*abstraction.Context, *dto.SpiSdmItemDeleteRequest) (*dto.SpiSdmItemResponse, error)
	//Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error)
	GetSpiSdmItemBySpiSdmID(*abstraction.Context, *dto.SpiSdmItemViewBySpiSdmIDRequest) ([]dto.SpiSdmItemViewBySpiSdmIDResponse, error)
}

type spiSdmItemService struct {
	Repository repository.SpiSdmItem
	Db         *gorm.DB
}

func NewSpiSdmItemService(f *factory.Factory) *spiSdmItemService {
	repository := f.SpiSdmItemRepository
	db := f.Db
	return &spiSdmItemService{repository, db}

}

func (s *spiSdmItemService) Save(ctx *abstraction.Context, payload *dto.SpiSdmItemSaveRequest) (*dto.SpiSdmItemResponse, error) {

	var result *dto.SpiSdmItemResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdm, err := s.Repository.Create(ctx, &model.SpiSdmItem{
			Context:          ctx,
			SpiSdmItemEntity: payload.SpiSdmItemEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmItemResponse{
			ID:               abstraction.ID{ID: spisdm.ID},
			SpiSdmItemEntity: spisdm.SpiSdmItemEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmItemService) Update(ctx *abstraction.Context, payload *dto.SpiSdmItemUpdateRequest) (*dto.SpiSdmItemResponse, error) {

	var result *dto.SpiSdmItemResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiSdmItem, err := s.Repository.FindByID(ctx, &model.SpiSdmItem{
			Context:          ctx,
			EntityInc:        abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
			SpiSdmItemEntity: payload.SpiSdmItemEntity,
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiSdmItem, err = s.Repository.Update(ctx, &model.SpiSdmItem{
			Context:          ctx,
			EntityInc:        abstraction.EntityInc{IDInc: abstraction.IDInc{ID: spiSdmItem.ID}},
			SpiSdmItemEntity: payload.SpiSdmItemEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmItemResponse{
			ID:               abstraction.ID{ID: spiSdmItem.ID},
			SpiSdmItemEntity: spiSdmItem.SpiSdmItemEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmItemService) Delete(ctx *abstraction.Context, payload *dto.SpiSdmItemDeleteRequest) (*dto.SpiSdmItemResponse, error) {

	var result *dto.SpiSdmItemResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiSdmItem, err := s.Repository.FindByID(ctx, &model.SpiSdmItem{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID},
		}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiSdmItem, err = s.Repository.Delete(ctx, &model.SpiSdmItem{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID},
		}})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmItemResponse{
			ID: abstraction.ID{ID: spiSdmItem.EntityInc.ID},
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmItemService) GetSpiSdmItemBySpiSdmID(ctx *abstraction.Context, payload *dto.SpiSdmItemViewBySpiSdmIDRequest) ([]dto.SpiSdmItemViewBySpiSdmIDResponse, error) {
	var result []dto.SpiSdmItemViewBySpiSdmIDResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		result, err = s.Repository.ViewSpiSdmItemBySpiSdmID(ctx, payload)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	return result, nil
}
