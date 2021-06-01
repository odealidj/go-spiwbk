package response

import "code-boiler/internal/abstractions"

type Meta struct {
	Success bool                         `json:"success" default:"true"`
	Message string                       `json:"message" default:"true"`
	Info    *abstractions.PaginationInfo `json:"info"`
}

type responseHelper struct {
	Error   errorHelper
	Success successHelper
}

var Constant responseHelper = responseHelper{
	Error:   errorConstant,
	Success: successConstant,
}
