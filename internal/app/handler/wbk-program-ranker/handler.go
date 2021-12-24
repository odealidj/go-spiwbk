package wbk_program_ranker

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/dto"
	"codeid-boiler/internal/app/service"
	"codeid-boiler/internal/factory"
	res "codeid-boiler/pkg/util/response"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service service.WbkProgramRankerService
}

var err error

func NewHandler(f *factory.Factory) *handler {
	service := service.NewWbkProgramRankerService(f)
	return &handler{service}
}

func (h *handler) GetByThnAngIDAndSatkerID(c echo.Context) error {
	cc := c.(*abstraction.Context)

	payload := new(dto.WbkProgramRankerGetRequest)
	if err := c.Bind(payload); err != nil {
		//fmt.Println(1)
		return res.ErrorBuilder(&res.ErrorConstant.BadRequest, err).Send(c)
	}
	if err = c.Validate(payload); err != nil {
		//fmt.Println(2)
		return res.ErrorBuilder(&res.ErrorConstant.Validation, err).Send(c)
	}

	result, err := h.service.GetByThnAngIDAndSatkerID(cc, payload)
	if err != nil {
		return res.ErrorResponse(err).Send(c)
	}

	return res.CustomSuccessBuilder(200, result.Datas, "Get datas success", result.PaginationInfo).Send(c)
}
