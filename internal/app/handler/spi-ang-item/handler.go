package spi_ang_item

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/service"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"github.com/labstack/echo/v4"
)

var err error

type handler struct {
	service service.SpiAngItemService
}

func NewHandler(f *factory.Factory) *handler {
	service := service.NewSpiAngItemService(f)
	return &handler{service}
}

func (h *handler) Save(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SpiAngItemSaveRequest)
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
