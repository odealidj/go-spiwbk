package thnang

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/service"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"fmt"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service service.ThnAngService
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := service.NewThnAngService(f)
	return &handler{service}
}

func (h *handler) Save(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {

		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.Save(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) SaveFrom(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngRequestForm)

	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {

		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	data, err := h.service.SaveForm(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) SaveBatch(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payloads := &[]dto.ThnAngRequests{}
	if err = c.Bind(payloads); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}

	for _, payload := range *payloads {

		if err = c.Validate(payload); err != nil {
			return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
		}
	}

	data, err := h.service.SaveBatch(cc, *payloads)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) Update(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngUpdateRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	fmt.Println(1)
	data, err := h.service.Update(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) Delete(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(abstraction.ID)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	fmt.Println(1)
	data, err := h.service.Delete(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) GetAll(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngGetAllRequest)
	if err := c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.GetAll(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", result.PaginationInfo).Send(c)

}

func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngGetRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.Get(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) GetByYear(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.ThnAngGetByYearRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.GetByYear(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}
