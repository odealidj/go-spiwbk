package spi_pbj_paket_jenis_belanja_pagu

import (
	"codeid-boiler/internal/abstraction"
	spi_pbj2 "codeid-boiler/internal/app/dto/spi-pbj"
	"codeid-boiler/internal/app/service/spi-pbj"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"github.com/labstack/echo/v4"
)

var err error

type handler struct {
	service spi_pbj.SpiPbjPaketJenisBelanjaPaguService
}

func NewHandler(f *factory.Factory) *handler {
	service := spi_pbj.NewSpiPbjPaketJenisBelanjaPaguService(f)
	return &handler{service}
}

func (h *handler) Save(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(spi_pbj2.SpiPbjPaketJenisBelanjaPaguSaveRequest)
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

func (h *handler) GetSpiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(spi_pbj2.SpiPbjPaketJenisBelanjaPaguGetRequest)
	if err := c.Bind(payload); err != nil {
		//fmt.Println(1)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		//fmt.Println(2)
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.GetSpiPbjPaketJenisBelanjaPaguByThnAngIDAndSatkerID(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", result.PaginationInfo).Send(c)
}
