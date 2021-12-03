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

type RkaklService interface {
	Save(*abstraction.Context, *dto.RkaklSaveRequest, *multipart.FileHeader) (*dto.RkaklResponse, error)
	Update(*abstraction.Context, *dto.RkaklUpdateRequest, *multipart.FileHeader) (*dto.RkaklResponse, error)
	Delete(*abstraction.Context, *dto.RkaklDeleteRequest) (*dto.RkaklDeleteResponse, error)
	Get(*abstraction.Context, *dto.RkaklGetRequest) (*dto.RkaklGetInfoResponse, error)
	GetByID(*abstraction.Context, *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error)
	GetByID2(*abstraction.Context, *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error)
}

type rkaklService struct {
	Repository          repository.Rkakl
	RkaklFileRepository repository.RkaklFile
	Db                  *gorm.DB
}

func NewRkaklService(f *factory.Factory) *rkaklService {
	repository := f.RkaklRepository
	rkaklFileRepository := f.RkaklFileRepository
	db := f.Db
	return &rkaklService{repository, rkaklFileRepository, db}

}

func (s *rkaklService) Save(ctx *abstraction.Context, payload *dto.RkaklSaveRequest, file *multipart.FileHeader) (*dto.RkaklResponse, error) {

	var result *dto.RkaklResponse
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		src, err := file.Open()
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
		}
		defer src.Close()

		rkakl, err := s.Repository.Create(ctx, &model.Rkakl{
			Context:     ctx,
			RkaklEntity: payload.RkaklEntity,
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		fileName := fmt.Sprintf("rkakl%d%s", rkakl.ID, filepath.Ext(file.Filename))

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

		/*
			f, err := excelize.OpenFile("./" + payload.FilePath)
			if err != nil {
				//fmt.Println(22)
				//fmt.Println(err)
				return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
				//return
			}

			cell, err := f.GetCellValue("SPI-SDM", "B2")
			if err != nil {
				//fmt.Println(err)
				return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
			}
		*/
		rkaklFile, err := s.RkaklFileRepository.Create(ctx, &model.RkaklFile{
			Entity:          abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}},
			RkaklFileEntity: model.RkaklFileEntity{Filepath: strings.TrimSpace(strings.Split(payload.FilePath, "/")[1])},
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *rkaklService) Update(ctx *abstraction.Context, payload *dto.RkaklUpdateRequest, file *multipart.FileHeader) (*dto.RkaklResponse, error) {

	var result *dto.RkaklResponse
	var oldFile, fileLocation string
	//var data *model.ThnAng

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkakl, err = s.Repository.Update(ctx, &model.Rkakl{Context: ctx,
			EntityInc:   abstraction.EntityInc{IDInc: abstraction.IDInc{ID: rkakl.ID}},
			RkaklEntity: payload.RkaklEntity,
		})
		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{Context: ctx,
			Entity: abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		if file != nil {

			src, err := file.Open()
			if err != nil {
				return res.ErrorBuilder(&res.ErrorConstant.UploadFileSrcError, err)
			}
			defer src.Close()

			fileName := fmt.Sprintf("rkakl%d%s", rkakl.ID, filepath.Ext(file.Filename))

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

			oldFile = path.Join("upload", rkaklFile.Filepath)

			fileLocation = destinationPath
			//payload.ImagePath = destinationPath
		} else {
			//fmt.Println("No file")
			fileLocation = rkaklFile.Filepath
		}

		rkaklFile, err = s.RkaklFileRepository.Update(ctx, &model.RkaklFile{
			Entity:          abstraction.Entity{ID: abstraction.ID{ID: rkakl.ID}},
			RkaklFileEntity: model.RkaklFileEntity{Filepath: strings.TrimSpace(strings.Split(fileLocation, "/")[1])},
		})

		if err != nil {
			if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
				return res.ErrorBuilder(&res.ErrorConstant.Duplicate, err)
			} else if strings.Contains(strings.ToLower(err.Error()), "foreign key") {
				return res.ErrorBuilder(&res.ErrorConstant.ForeignKey, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		//tidak diperlukan karena nama file nya sama
		//if len(oldFile) > 0 {
		//	_ = os.Remove(oldFile)
		//}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil

}

func (s *rkaklService) Delete(ctx *abstraction.Context, payload *dto.RkaklDeleteRequest) (*dto.RkaklDeleteResponse, error) {
	var result *dto.RkaklDeleteResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{Context: ctx, EntityInc: abstraction.EntityInc{
			IDInc: abstraction.IDInc{ID: payload.ID.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{Context: ctx, Entity: abstraction.Entity{
			ID: abstraction.ID{ID: rkakl.ID}}})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		fmt.Println(path.Join("upload", rkaklFile.Filepath))

		if len(rkaklFile.Filepath) > 0 {

			_ = os.Remove(path.Join("upload", rkaklFile.Filepath))
		}

		_, err = s.Repository.Delete(ctx, &model.Rkakl{Context: ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: rkakl.ID}},
		})
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklDeleteResponse{ID: payload.ID}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) Get(ctx *abstraction.Context, payload *dto.RkaklGetRequest) (*dto.RkaklGetInfoResponse, error) {
	var result *dto.RkaklGetInfoResponse

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakls, info, err := s.Repository.Find(ctx, &payload.RkaklFilter, &payload.Pagination)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
		}

		result = &dto.RkaklGetInfoResponse{
			Datas:          &rkakls,
			PaginationInfo: info,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) GetByID(ctx *abstraction.Context, payload *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error) {
	var result *dto.RkaklResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakl, err := s.Repository.FindByID(ctx, &model.Rkakl{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}
		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{
			Context: ctx,
			Entity:  abstraction.Entity{ID: abstraction.ID{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			Filepath:    &rkaklFile.Filepath,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *rkaklService) GetByID2(ctx *abstraction.Context, payload *dto.RkaklGetByIDRequest) (*dto.RkaklResponse, error) {
	var result *dto.RkaklResponse

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {
		rkakl, err := s.Repository.FindThnAngSatkerByID(ctx, &model.Rkakl{
			Context:   ctx,
			EntityInc: abstraction.EntityInc{IDInc: abstraction.IDInc{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				//result = &dto.RkaklResponse{}
				return nil
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		//if rkakl.ID != 0 && rkakl.SatkerID != 0 {

		rkaklFile, err := s.RkaklFileRepository.FindByID(ctx, &model.RkaklFile{
			Context: ctx,
			Entity:  abstraction.Entity{ID: abstraction.ID{ID: payload.ID.ID}},
		})
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
			}
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		result = &dto.RkaklResponse{
			ID:          abstraction.ID{ID: rkakl.ID},
			RkaklEntity: rkakl.RkaklEntity,
			ThnAngYear:  &rkakl.ThnAng.Year,
			SatkerName:  &rkakl.Satker.Name,
			Filepath:    &rkaklFile.Filepath,
		}
		//} else {
		//result = &dto.RkaklResponse{}
		//}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
