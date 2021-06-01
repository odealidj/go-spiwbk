package sample

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

var err error

func NewHandler() *handler {
	dbConnection, err := database.DBManager(strings.ToUpper("sample2"))
	if err != nil {
		panic("there is no database for your account, Please contact admin handler")
	}
	return &handler{dbConnection: dbConnection}
}

// Get godoc
// @Summary Get samples
// @Description Get samples
// @Tags samples
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param page query int false "Page of pagination"
// @Param page_size query int false "Page Size of pagination"
// @Param sort query string false "Sort"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /samples [get]
func (h *handler) Get(c echo.Context) error {
	h.service = NewService(h.dbConnection)

	//logic
	payload := new(dto.SampleGetRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}

	data, info, err := h.service.Find(payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	// return res.SuccessResponse(data).Send(c)
	return res.CustomSuccessBuilder(200, data, "Data has been retrieve", info).Send(c)
}

// Get By ID godoc
// @Summary Get samples by id
// @Description Get samples by id
// @Tags samples
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "resource id"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /samples/{id} [get]
func (h *handler) GetByID(c echo.Context) error {
	h.service = NewService(h.dbConnection)

	//logic
	payload := new(dto.SampleGetByIDRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		response := res.ErrorBuilder(res.Constant.Error.Validation, err)
		return response.Send(c)
	}

	data, err := h.service.FindByID(payload.ID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := dto.SampleGetByIDResponse{
		abstractions.Model{
			ID:         data.ID,
			CreatedAt:  data.CreatedAt,
			CreatedBy:  data.CreatedBy,
			ModifiedAt: data.ModifiedAt,
			ModifiedBy: data.ModifiedBy,
		},
		dto.Sample{
			Key:   data.Key,
			Value: data.Value,
		},
	}

	return res.SuccessResponse(result).Send(c)
}

// Create godoc
// @Summary Create samples
// @Description Create samples
// @Tags samples
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param create body dto.SampleStoreRequest true "request body"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /samples [post]
func (h *handler) Store(c echo.Context) error {
	h.service = NewService(h.dbConnection)
	trx := h.dbConnection.Begin()
	c.Set("db_transaction", trx)

	//logic
	payload := new(dto.SampleStoreRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.service.WithTrx(trx).Create(payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := dto.SampleStoreResponse{
		abstractions.Model{
			ID:         data.ID,
			CreatedAt:  data.CreatedAt,
			CreatedBy:  data.CreatedBy,
			ModifiedAt: data.ModifiedAt,
			ModifiedBy: data.ModifiedBy,
		},
		dto.Sample{
			Key:    data.Key,
			Value:  data.Value,
			UserId: data.UserId,
		},
	}

	return res.SuccessResponse(result).Send(c)
}

// Update godoc
// @Summary Update samples
// @Description Update samples
// @Tags samples
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param update body dto.SampleUpdateRequest true "request body"
// @Param id path int true "resource id"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /samples/{id} [patch]
func (h *handler) Update(c echo.Context) error {
	h.service = NewService(h.dbConnection)
	trx := h.dbConnection.Begin()
	c.Set("db_transaction", trx)

	//logic
	payload := new(dto.SampleUpdateRequest)
	// id, _ := strconv.Atoi(c.Param("id"))

	if err := c.Bind(&payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.service.WithTrx(trx).Update(payload.ID, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := dto.SampleUpdateResponse{
		abstractions.Model{
			ID:         data.ID,
			CreatedAt:  data.CreatedAt,
			CreatedBy:  data.CreatedBy,
			ModifiedAt: data.ModifiedAt,
			ModifiedBy: data.ModifiedBy,
		},
		dto.Sample{
			Key:    data.Key,
			Value:  data.Value,
			UserId: data.UserId,
		},
	}

	return res.SuccessResponse(result).Send(c)
}

// Delete godoc
// @Summary Delete samples
// @Description Delete samples
// @Tags samples
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "resource id"
// @Success 200 {object} res.successResponse
// @Failure 400 {object} res.errorResponse
// @Failure 404 {object} res.errorResponse
// @Failure 500 {object} res.errorResponse
// @Router /samples/{id} [delete]
func (h *handler) Delete(c echo.Context) error {
	h.service = NewService(h.dbConnection)
	trx := h.dbConnection.Begin()
	c.Set("db_transaction", trx)

	//logic
	payload := new(dto.SampleDeleteRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.BadRequest, err).Send(c)
	}

	if err := c.Validate(payload); err != nil {
		return res.ErrorBuilder(res.Constant.Error.Validation, err).Send(c)
	}

	data, err := h.service.WithTrx(trx).Delete(payload.ID)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	result := dto.SampleUpdateResponse{
		abstractions.Model{
			ID:         data.ID,
			CreatedAt:  data.CreatedAt,
			CreatedBy:  data.CreatedBy,
			ModifiedAt: data.ModifiedAt,
			ModifiedBy: data.ModifiedBy,
		},
		dto.Sample{
			Key:    data.Key,
			Value:  data.Value,
			UserId: data.UserId,
		},
	}

	return res.SuccessResponse(result).Send(c)
}
