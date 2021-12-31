package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
)

type MethodApbjResponse struct {
	abstraction.ID
	spi_pbj.MethodApbjEntity
}

//Get
type MethodApbjGetRequest struct {
	abstraction.Pagination
	spi_pbj.MethodApbjFilter
}

type MethodApbjGetResponses struct {
	Datas          *[]MethodApbjResponse
	PaginationInfo *abstraction.PaginationInfo
}
