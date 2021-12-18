package spi_pbj_rekapitulasi

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
	service service.SpiPbjRekapitulasiService
}

func NewHandler(f *factory.Factory) *handler {
	service := service.NewSpiPbjRekapitulasiService(f)
	return &handler{service}
}

func (h *handler) Upsert(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SpiPbjRekapitulasiUpsertRequest)
	if err = c.Bind(payload); err != nil {
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {

		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}
	data, err := h.service.Upsert(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.SuccessResponse(data).Send(c)
}

func (h *handler) GetSpiPbjRekapitulasiByID(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.SpiPbjRekapitulasiGetRequest)
	if err := c.Bind(payload); err != nil {
		//fmt.Println(1)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		//fmt.Println(2)
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.GetSpiPbjRekapitulasiByID(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", result.PaginationInfo).Send(c)
}
