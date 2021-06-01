package auth

import (
	"code-boiler/database"
	"code-boiler/internal/abstractions"
	"code-boiler/internal/dto"
	res "code-boiler/pkg/util/response"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	service      *service
	dbConnection *gorm.DB
}

func NewHandler() *handler {
	dbConnection, err := database.DBManager(strings.ToUpper("sample2"))
	if err != nil {
		panic("there is no database for your account, Please contact admin handler")
	}
	return &handler{dbConnection: dbConnection}
}

// Login godoc
// @Summary Login user account
// @Description Login user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param login body dto.AuthLoginRequest true "request body login"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /auth/login [post]
func (h *handler) Login(c echo.Context) (err error) {
	h.service = NewService(h.dbConnection)

	//logic
	payload := new(dto.AuthLoginRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	user, token, err := h.service.Login(payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	data := &dto.AuthLoginResponse{
		Token: token,
		User: dto.User{
			Name:     user.Name,
			Phone:    user.Phone,
			Email:    user.Email,
			Status:   user.Status,
			IsActive: user.IsActive,
		},
	}

	return res.SuccessResponse(data).Send(c)
}

// Register godoc
// @Summary Register user account
// @Description Register user account
// @Tags auth
// @Accept  json
// @Produce  json
// @Param register body dto.AuthRegisterRequest true "request body register"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /auth/register [post]
func (h *handler) Register(c echo.Context) (err error) {
	h.service = NewService(h.dbConnection)
	trx := h.dbConnection.Begin()
	c.Set("db_transaction", trx)

	//logic
	payload := new(dto.AuthRegisterRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	//register -> store user
	user, err := h.service.WithTrx(trx).Register(payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := dto.AuthRegisterResponse{
		abstractions.Model{
			ID:         user.ID,
			CreatedAt:  user.CreatedAt,
			CreatedBy:  user.CreatedBy,
			ModifiedAt: user.ModifiedAt,
			ModifiedBy: user.ModifiedBy,
		},
		dto.User{
			Name:     user.Name,
			Phone:    user.Phone,
			Email:    user.Email,
			Status:   user.Status,
			IsActive: user.IsActive,
		},
	}
	return res.SuccessResponse(result).Send(c)
}
