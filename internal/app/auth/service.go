package auth

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/auth/dto"
	"codeid-boiler/internal/app/auth/model"
	"codeid-boiler/internal/app/auth/repository"
	"codeid-boiler/internal/factory"
	"fmt"

	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	//Login(*abstraction.Context, *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(*abstraction.Context, *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type service struct {
	Repository repository.Auth
	Db         *gorm.DB
}

func NewService(f *factory.Factory) *service {
	repository := f.AuthRepository
	db := f.Db
	return &service{repository, db}
}

func (s *service) Login(ctx *abstraction.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	var result *dto.LoginResponse

	fmt.Println("s0")
	data, err := s.Repository.FindByUsername(ctx, &payload.Username)
	if err != nil {
		fmt.Println("s1")
		return result, res.ErrorBuilder(&res.ErrorConstant.Unauthorized, err)
	}

	fmt.Println(data.Username)
	
	if err = bcrypt.CompareHashAndPassword([]byte(data.Passwordhash), []byte(payload.Password)); err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}
	

	token, err := data.GenerateToken()
	if err != nil {
		return result, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result = &dto.LoginResponse{
		Token:         token,
		UserAppEntity: data.UserApp.UserAppEntity,
	}

	return result, nil
}

func (s *service) Register(ctx *abstraction.Context, payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var result *dto.RegisterResponse
	var data *model.UserApp

	if err = trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		data, err = s.Repository.Create(ctx, &model.UserApp{UserAppEntity: payload.UserAppEntity},
			&model.LoginApp{LoginAppEntity: model.LoginAppEntity{Username: payload.Username, Password: payload.Password}},
		)
		if err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.UnprocessableEntity, err)
		}

		return nil
	}); err != nil {
		return result, err
	}

	result = &dto.RegisterResponse{
		ID:            data.ID,
		UserAppEntity: data.UserAppEntity,
	}

	return result, nil
}
