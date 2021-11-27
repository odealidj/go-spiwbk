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
	"fmt"
	"gorm.io/gorm"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type SpiSdmService interface {
	Save(*abstraction.Context, *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error)
	SaveWithFile(*abstraction.Context, *dto.SpiSdmSaveRequest, *multipart.FileHeader) (*dto.SpiSdmResponse, error)
	Update(*abstraction.Context, *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error)
	Delete(*abstraction.Context, *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error)
	Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error)
	GetByID(*abstraction.Context, *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error)
}

type spiSdmService struct {
	Repository           repository.SpiSdm
	SpiSdmFileRepository repository.SpiSdmFile
	Db                   *gorm.DB
}

func NewSpiSdmService(f *factory.Factory) *spiSdmService {
	repository := f.SpiSdmRepository
	spiSdmFileRepository := f.SpiSdmFileRepository
	db := f.Db
	return &spiSdmService{repository, spiSdmFileRepository, db}

}

func (s *spiSdmService) Save(ctx *abstraction.Context, payload *dto.SpiSdmSaveRequest) (*dto.SpiSdmResponse, error) {

	var result *dto.SpiSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdm, err := s.Repository.Create(ctx, &model.SpiSdm{
			Context:      ctx,
			SpiSdmEntity: payload.SpiSdmEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmResponse{
			ID:           abstraction.ID{ID: spisdm.ID},
			SpiSdmEntity: spisdm.SpiSdmEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmService) SaveWithFile(ctx *abstraction.Context, payload *dto.SpiSdmSaveRequest, file *multipart.FileHeader) (*dto.SpiSdmResponse, error) {

	var result *dto.SpiSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		src, err := file.Open()
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
		}
		defer src.Close()

		spisdm, err := s.Repository.Create(ctx, &model.SpiSdm{
			Context:      ctx,
			SpiSdmEntity: payload.SpiSdmEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		fileName := fmt.Sprintf("%d%s", spisdm.ID, filepath.Ext(file.Filename))

		//pathApp, _ := os.Getwd()

		// Destination
		destinationPath := path.Join("upload", fileName)

		dst, err := os.Create(destinationPath)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileCreateError, err)
			//return res.ErrorResponse(err).Send(c)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			//fmt.Println("no copy")
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileDestError, err)
		}

		payload.FilePath = destinationPath

		spiSdmFile, err := s.SpiSdmFileRepository.Create(ctx, &model.SpiSdmFile{
			Entity:           abstraction.Entity{ID: abstraction.ID{ID: spisdm.ID}},
			SpiSdmFileEntity: model.SpiSdmFileEntity{FilePath: payload.FilePath},
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmResponse{
			ID:           abstraction.ID{ID: spisdm.ID},
			SpiSdmEntity: spisdm.SpiSdmEntity,
			FilePath:     spiSdmFile.FilePath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmService) Update(ctx *abstraction.Context, payload *dto.SpiSdmUpdateRequest) (*dto.SpiSdmResponse, error) {

	var result *dto.SpiSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiSdm, err := s.Repository.FindByID(ctx, &model.SpiSdm{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiSdm, err = s.Repository.Update(ctx, &model.SpiSdm{
			Context:      ctx,
			EntityInc:    abstraction.EntityInc{IDInc: abstraction.IDInc{ID: spiSdm.ID}},
			SpiSdmEntity: payload.SpiSdmEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmResponse{
			ID:           abstraction.ID{ID: spiSdm.ID},
			SpiSdmEntity: spiSdm.SpiSdmEntity,
			//ThnAngYear:   spiSdm.ThnAng.Year,
			//SatkerName:   spiSdm.Satker.Name,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmService) Delete(ctx *abstraction.Context, payload *dto.SpiSdmDeleteRequest) (*dto.SpiSdmResponse, error) {

	var result *dto.SpiSdmResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		spiSdm, err := s.Repository.FindByID(ctx, &model.SpiSdm{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		spiSdm, err = s.Repository.Delete(ctx, &model.SpiSdm{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: spiSdm.ID}},
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.SpiSdmResponse{
			ID: abstraction.ID{ID: spiSdm.ID},
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *spiSdmService) Get(ctx *abstraction.Context, payload *dto.SpiSdmGetRequest) (*dto.SpiSdmGetResponse, error) {
	var result *dto.SpiSdmGetResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdms, info, err := s.Repository.Find(ctx, &payload.SpiSdmFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}
		if len(*spisdms) == 0 {
			return res.ErrorBuilder(&res.ErrorConstant.NotFound, errors.New("Data Not Found!"))
		}

		spiSdmResponses := &[]dto.SpiSdmResponses{}
		spiSdmResponse := &dto.SpiSdmResponses{}
		for _, spisdm := range *spisdms {
			spiSdmResponse.ID.ID = spisdm.ID
			spiSdmResponse.SpiSdmEntity = spisdm.SpiSdmEntity
			spiSdmResponse.ThnAngYear = spisdm.ThnAng.Year
			spiSdmResponse.SatkerName = spisdm.Satker.Name
			*spiSdmResponses = append(*spiSdmResponses, *spiSdmResponse)
		}
		result = &dto.SpiSdmGetResponse{
			Datas:          spiSdmResponses,
			PaginationInfo: info,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *spiSdmService) GetByID(ctx *abstraction.Context, payload *dto.SpiSdmGetByIDRequest) (*dto.SpiSdmResponse, error) {
	var result *dto.SpiSdmResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		spisdm, err := s.Repository.FindByID(ctx, &model.SpiSdm{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		result = &dto.SpiSdmResponse{
			ID:           abstraction.ID{ID: spisdm.ID},
			SpiSdmEntity: spisdm.SpiSdmEntity,
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
