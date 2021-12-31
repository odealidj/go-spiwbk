package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
)

type JenisBelanjaPaguResponse struct {
	abstraction.ID
	spi_pbj.JenisBelanjaPaguEntity
}

//Get
type JenisBelanjaPaguGetRequest struct {
	abstraction.Pagination
	spi_pbj.JenisBelanjaPaguFilter
}

type JenisBelanjaPaguGetResponses struct {
	Datas          *[]JenisBelanjaPaguResponse
	PaginationInfo *abstraction.PaginationInfo
}
