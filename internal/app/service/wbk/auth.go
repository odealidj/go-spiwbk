package wbk

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/model"
	"codeid-boiler/internal/app/repository"
	"codeid-boiler/internal/factory"
	"codeid-boiler/pkg/util/email"
	res "codeid-boiler/pkg/util/response"
	"codeid-boiler/pkg/util/trxmanager"
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
)

type AuthService interface {
	Login(*abstraction.Context, *dto.LoginRequest) (*dto.LoginResponse, error)
	Register(*abstraction.Context, *dto.RegisterRequest) (*dto.RegisterResponse, error)
}

type authservice struct {
	LoginAppRepository repository.LoginApp
	UserAppRepository  repository.UserApp
	Db                 *gorm.DB
}

func NewAuthService(f *factory.Factory) *authservice {
	repositoryLoginApp := f.LoginAppRepository
	repositoryUserApp := f.UserAppRepository
	db := f.Db
	return &authservice{repositoryLoginApp, repositoryUserApp, db}
}

type StructLoginApp struct {
	*model.LoginApp
	ErrConst *res.Error
	Err      *error
}
type StructUserApp struct {
	*model.UserApp
	ErrConst *res.Error
	Err      *error
}

func (s *authservice) Login(ctx *abstraction.Context, payload *dto.LoginRequest) (*dto.LoginResponse, error) {
	var result *dto.LoginResponse
	var structLoginApp *StructLoginApp
	var structUserApp *StructUserApp
	//var dataLogin *model.LoginApp
	//var dataUser *model.UserApp
	var token string
	var wg sync.WaitGroup

	app_id, err := strconv.ParseInt(os.Getenv("WBK"), 10, 8)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		ChLoginApp := make(chan *StructLoginApp)
		ChUserApp := make(chan *StructUserApp)
		defer close(ChLoginApp)
		defer close(ChUserApp)

		wg.Add(3)

		go func() {
			defer wg.Done()
			dataLogin, err := s.LoginAppRepository.FindLoginAppByAppIDAndUsername(ctx, &model.LoginApp{LoginAppEntity: model.LoginAppEntity{AppID: int(app_id), Username: payload.Username, Password: payload.Password}})
			if err != nil {
				ChLoginApp <- &StructLoginApp{
					LoginApp: nil, ErrConst: &res.ErrorConstant.Unauthorized, Err: &err,
				}
				return
			}

			if err = bcrypt.CompareHashAndPassword([]byte(dataLogin.Passwordhash), []byte(payload.Password)); err != nil {
				ChLoginApp <- &StructLoginApp{
					LoginApp: nil, ErrConst: &res.ErrorConstant.Unauthorized, Err: &err,
				}
				return
			}

			token, err = dataLogin.GenerateToken()
			if err != nil {
				ChLoginApp <- &StructLoginApp{
					LoginApp: nil, ErrConst: &res.ErrorConstant.InternalServerError, Err: &err,
				}
				return
			}
			ChLoginApp <- &StructLoginApp{
				LoginApp: dataLogin, ErrConst: nil, Err: nil,
			}
		}()

		go func() {
			defer wg.Done()
			counter := 0
			for {

				select {
				case structLoginApp = <-ChLoginApp:

					userApp := &StructUserApp{
						UserApp: nil, ErrConst: structLoginApp.ErrConst, Err: structLoginApp.Err,
					}

					if structLoginApp.LoginApp != nil {
						dataUser, err := s.UserAppRepository.Find(ctx, &model.UserApp{Entity: abstraction.Entity{
							ID: abstraction.ID{ID: structLoginApp.ID},
						}})
						if err != nil {
							userApp.Err = &err
						} else {
							userApp.UserApp = dataUser
							userApp.Err = nil
						}
					}
					ChUserApp <- userApp
				}
				counter++
				if counter == 1 {
					break
				}
			}
		}()

		go func() {
			defer wg.Done()
			counter := 0
			for {
				select {
				case structUserApp = <-ChUserApp:
					counter++

				}
				if counter == 1 {
					break
				}
			}
		}()

		wg.Wait()

		if structUserApp.Err != nil {
			return res.ErrorBuilder(structUserApp.ErrConst, *structUserApp.Err)
		}

		result = &dto.LoginResponse{
			Token:         token,
			ID:            abstraction.ID{ID: structUserApp.ID.ID},
			UserAppEntity: structUserApp.UserAppEntity,
		}

		return nil

	}); err != nil {
		return nil, err
	}

	return result, nil
}

func (s *authservice) Register(ctx *abstraction.Context, payload *dto.RegisterRequest) (*dto.RegisterResponse, error) {
	var result *dto.RegisterResponse
	var structLoginApp *StructLoginApp
	var structUserApp *StructUserApp
	var wg sync.WaitGroup

	app_id, err := strconv.ParseInt(os.Getenv("WBK"), 10, 8)
	if err != nil {
		return nil, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if int(*payload.UserAppEntity.SatkerId) == 0 {
		payload.UserAppEntity.SatkerId = nil
	}

	if *payload.UserAppEntity.RoleUserId == 0 {
		payload.UserAppEntity.RoleUserId = nil
	}

	if *payload.UserAppEntity.JabatanId == 0 {
		payload.UserAppEntity.JabatanId = nil
	}

	if *payload.UserAppEntity.Email != "" {
		if email.Valid(*payload.Email) == false {
			return nil, res.CustomErrorBuilderWithData(http.StatusBadRequest,
				"Format email tidak valid", "Format email tidak valid")
		}
	} else {
		payload.UserAppEntity.Email = nil
	}

	if err := trxmanager.New(s.Db).WithTrx(ctx, func(ctx *abstraction.Context) error {

		ChLoginApp := make(chan *StructLoginApp)
		defer close(ChLoginApp)
		ChUserApp := make(chan *StructUserApp)
		defer close(ChUserApp)

		wg.Add(3)

		go func(wgs *sync.WaitGroup) {
			defer wgs.Done()
			loginApp, err := s.LoginAppRepository.CreateLoginApp(ctx, &model.LoginApp{LoginAppEntity: model.LoginAppEntity{AppID: int(app_id), Username: payload.Username, Password: payload.Password}})
			if err != nil {
				if strings.Contains(strings.ToLower(err.Error()), "duplicate") {

					errorDuplicate := errors.New("Invalid parameters or payload")
					ChLoginApp <- &StructLoginApp{
						LoginApp: nil, ErrConst: &res.ErrorConstant.Duplicate, Err: &errorDuplicate,
					}
				} else {

					ChLoginApp <- &StructLoginApp{
						LoginApp: nil, ErrConst: &res.ErrorConstant.UnprocessableEntity, Err: &err,
					}
				}
			} else {
				ChLoginApp <- &StructLoginApp{
					LoginApp: loginApp, ErrConst: nil, Err: nil,
				}
			}

		}(&wg)

		go func(wgs *sync.WaitGroup) {
			defer wgs.Done()
			counter := 0

			for {
				select {
				case structLoginApp = <-ChLoginApp:
					counter++
					structUserApp := &StructUserApp{
						UserApp: nil, ErrConst: &res.ErrorConstant.UnprocessableEntity, Err: structLoginApp.Err,
					}

					switch structLoginApp.LoginApp {
					case nil:
						structUserApp.Err = structLoginApp.Err
					default:
						userApp, err := s.UserAppRepository.CreateUserApp(ctx,
							&model.UserApp{Entity: abstraction.Entity{ID: abstraction.ID{
								ID: structLoginApp.LoginApp.ID}},
								UserAppEntity: payload.UserAppEntity})
						if err != nil {
							if strings.Contains(strings.ToLower(err.Error()), "duplicate") {
								errorDuplicate := errors.New("Invalid parameters or payload")
								structUserApp.Err = &errorDuplicate
								structUserApp.ErrConst = &res.ErrorConstant.Duplicate
								break
							} else {
								structUserApp.Err = &err
								structUserApp.ErrConst = &res.ErrorConstant.UnprocessableEntity
								break
							}
						} else {
							structUserApp.UserApp = userApp
							structUserApp.ErrConst = nil
							structUserApp.Err = nil
							break
						}
					}
					ChUserApp <- structUserApp
					//fmt.Println("end switch")
				}

				if counter == 1 {
					//fmt.Println("counter")
					break
				}
			}

		}(&wg)

		go func(wgs *sync.WaitGroup) {
			defer wgs.Done()
			counter := 0
			for {
				select {
				case structUserApp = <-ChUserApp:
					counter++
				}
				if counter == 1 {
					break
				}
			}
		}(&wg)

		wg.Wait()

		if structLoginApp.Err != nil {
			fmt.Println(1)
			return res.ErrorBuilder(structLoginApp.ErrConst, *structLoginApp.Err)
		}
		if structUserApp.Err != nil {
			fmt.Println(2)
			return res.ErrorBuilder(structUserApp.ErrConst, *structUserApp.Err)
		}

		result = &dto.RegisterResponse{
			ID:            abstraction.ID{ID: structUserApp.ID.ID},
			UserAppEntity: structUserApp.UserAppEntity,
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return result, nil
}
