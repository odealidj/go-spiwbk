package service

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"gorm.io/gorm"
	"net/http"
)

type SpiAngItemService interface {
	Save(*abstraction.Context, *dto.SpiAngItemSaveRequest) ([]dto.SpiAngItemResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	//Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type spiAngItemService struct {
	Repository         repository.SpiAngItem
	SpiAngRepository   repository.SpiAng
	KomponenRepository repository.Komponen
	Db                 *gorm.DB
}

func NewSpiAngItemService(f *factory.Factory) *spiAngItemService {
	repository := f.SpiAngItemRepository
	spiAngRepository := f.SpiAngRepository
	komponenRepository := f.KomponenRepository
	db := f.Db
	return &spiAngItemService{repository, spiAngRepository,
		komponenRepository, db}

}

func (s *spiAngItemService) Save(ctx *abstraction.Context, payload *dto.SpiAngItemSaveRequest) ([]dto.SpiAngItemResponse, error) {

	var result []dto.SpiAngItemResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiAng, err := s.SpiAngRepository.Create(ctx, &model.SpiAng{Context: ctx, SpiAngEntity: model.SpiAngEntity{
			ThnAngID: uint16(payload.ThnAngID), SatkerID: uint16(payload.SatkerID),
		}})
		if err != nil {
			//if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
			//	return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
			//		"Duplicate spi ang", "Invalid spi ang")
			//}

			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid spi ang", "Invalid spi ang")
		}

		komponens, err := s.KomponenRepository.FindByThnAngIDAndSatkerID(ctx, &dto.KomponenFindByThnAngIDAndSatkerIDRequest{
			ThnAngID: payload.ThnAngID, SatkerID: payload.SatkerID,
		})
		if err != nil {
			return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
				"Invalid komponen", err.Error())
		}

		for _, komponen := range *komponens {

			spiAngItem, err := s.Repository.Create(ctx, &model.SpiAngItem{Context: ctx, SpiAngItemEntity: model.SpiAngItemEntity{
				SpiAngID: int(spiAng.ID), KomponenID: int(komponen.ID),
			}})
			if err != nil {
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid spi ang item", err.Error())
			}

			result = append(result, dto.SpiAngItemResponse{
				ID:               abstraction.ID{ID: spiAngItem.ID},
				SpiAngItemEntity: spiAngItem.SpiAngItemEntity,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}
