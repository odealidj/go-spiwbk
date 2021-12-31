package frekuensi_ranker

import (
	"codeid-boiler/internal/abstraction"
	wbk2 "codeid-boiler/internal/app/dto/wbk"
	"codeid-boiler/internal/app/service/wbk"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service wbk.FrekuensiRankerService
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := wbk.NewFrekuensiRankerService(f)
	return &handler{service}
}

func (h *handler) Save(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(wbk2.FrekuensiRankerUpsertRequest)
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

func (h *handler) Get(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(wbk2.FrekuensiRankerGetRequest)
	if err := c.Bind(payload); err != nil {
		//fmt.Println(1)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		//fmt.Println(2)
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.Get(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", result.PaginationInfo).Send(c)
}
