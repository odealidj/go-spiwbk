package auth

import (
	"code-boiler/internal/dto"
	"code-boiler/internal/model"
	"code-boiler/internal/repository"
	res "code-boiler/pkg/util/response"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(payload *dto.AuthRegisterRequest) (model.User, error)
	Login(payload *dto.AuthLoginRequest) (*model.User, string, error)

	WithTrx(trx *gorm.DB) *service
}

type service struct {
	Repository repository.User
}

func NewService(dbConnection *gorm.DB) *service {
	repository := repository.NewUser(dbConnection)
	return &service{Repository: repository}
}

func (s *service) WithTrx(trx *gorm.DB) *service {
	repository := repository.NewUser(trx)
	return &service{
		Repository: repository,
	}
}

func (s *service) Login(payload *dto.AuthLoginRequest) (*model.User, string, error) {
	user, err := s.Repository.FindByEmail(payload.Email)
	if user == nil {
		return nil, "", res.ErrorBuilder(res.Constant.Error.Unauthorized, err)
	}
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(payload.Password)) != nil {
		return nil, "", res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
	}

	return user, token, nil
}

func (s *service) Register(payload *dto.AuthRegisterRequest) (*model.User, error) {
	user, err := s.Repository.FindByEmail(payload.Phone)
	if user != nil {
		return nil, res.ErrorBuilder(res.Constant.Error.Duplicate, err)
	}
	if err != nil {
		if err.Error() != "record not found" {
			return nil, res.ErrorBuilder(res.Constant.Error.InternalServerError, err)
		}
	}

	user = &model.User{}
	user.Name = payload.Name
	user.Email = payload.Email
	user.Phone = payload.Phone
	user.Status = payload.Status
	user.IsActive = payload.IsActive
	user.Password = payload.Password

	user, err = s.Repository.Create(user)
	if err != nil {
		return user, res.ErrorBuilder(res.Constant.Error.UnprocessableEntity, err)
	}

	return user, nil
}
