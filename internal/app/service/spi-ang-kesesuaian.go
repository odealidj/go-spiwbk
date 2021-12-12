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

type SpiAngKesesuaianService interface {
	Save(*abstraction.Context, *dto.SpiAngKesesuaianSaveRequest) ([]dto.SpiAngKesesuaianResponse, error)
	//Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	//Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	//Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error)
	//GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
	GetBySpiSdmID(*abstraction.Context, *dto.SpiAngKesesuaianGetRequest) (*dto.SpiAngKesesuaianGetInfoResponse, error)
}

type spiAngKesesuaianService struct {
	SpiAngItemRepository       repository.SpiAngItem
	SpiAngRepository           repository.SpiAng
	KomponenRepository         repository.Komponen
	SpiAngKesesuaianRepository repository.SpiAngKesesuaian
	Db                         *gorm.DB
}

func NewSpiAngKesesuaianService(f *factory.Factory) *spiAngKesesuaianService {

	spiAngItemRepository := f.SpiAngItemRepository
	spiAngRepository := f.SpiAngRepository
	komponenRepository := f.KomponenRepository
	spiAngKesesuaianRepository := f.SpiAngKesesuaianRepository
	db := f.Db
	return &spiAngKesesuaianService{spiAngItemRepository, spiAngRepository,
		komponenRepository, spiAngKesesuaianRepository, db}

}

func (s *spiAngKesesuaianService) Save(ctx *abstraction.Context, payload *dto.SpiAngKesesuaianSaveRequest) ([]dto.SpiAngKesesuaianResponse, error) {

	var result []dto.SpiAngKesesuaianResponse
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

			spiAngItem, err := s.SpiAngItemRepository.Create(ctx, &model.SpiAngItem{Context: ctx, SpiAngItemEntity: model.SpiAngItemEntity{
				SpiAngID: int(spiAng.ID), KomponenID: int(komponen.ID),
			}})
			if err != nil {
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid spi ang item", err.Error())
			}

			spiAngKesesuaian, err := s.SpiAngKesesuaianRepository.Create(ctx, &model.SpiAngKesesuaian{Context: ctx,
				SpiAngKesesuaianEntity: model.SpiAngKesesuaianEntity{
					SpiAngItemID:      int(spiAngItem.ID),
					JenisKesesuaianID: payload.JenisKesesuaianID,
					JenisPengendaliID: payload.JenisPengendaliID, IsCheck: payload.IsCheck,
				}})
			if err != nil {
				return res.CustomErrorBuilderWithData(http.StatusUnprocessableEntity,
					"Invalid spi ang kesesuaian", err.Error())
			}

			result = append(result, dto.SpiAngKesesuaianResponse{
				ID:                     abstraction.ID{ID: spiAngKesesuaian.ID},
				SpiAngKesesuaianEntity: spiAngKesesuaian.SpiAngKesesuaianEntity,
				SatkerID:               payload.SatkerID,
				ThnAngID:               payload.ThnAngID,
			})
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiAngKesesuaianService) GetBySpiSdmID(ctx *abstraction.Context, payload *dto.SpiAngKesesuaianGetRequest) (
	*dto.SpiAngKesesuaianGetInfoResponse, error) {

	var result *dto.SpiAngKesesuaianGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		datas, info, err := s.SpiAngKesesuaianRepository.
			FindSpiKesesuaianByThnAngIDAndSatkerID(ctx, &payload.SpiAngKesesuaianFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		/*
			if len(dealers) <= 0 {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data not found"))
			}
		*/
		result = &dto.SpiAngKesesuaianGetInfoResponse{
			Datas:          &datas,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
