package response

import "codeid-boiler/internal/abstraction"

type Meta struct {
	Success bool                        `json:"success" default:"true"`
	Message string                      `json:"message" default:"true"`
	Info    *abstraction.PaginationInfo `json:"info"`
}

type Meta2 struct {
	Success bool                           `json:"success" default:"true"`
	Message string                         `json:"message" default:"true"`
	Info    *abstraction.PaginationInfoArr `json:"info"`
}
