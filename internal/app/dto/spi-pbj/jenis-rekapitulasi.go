package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
)

type JenisRekapitulasiResponse struct {
	abstraction.ID
	spi_pbj.JenisRekapitulasiEntity
}

//Get
type JenisRekapitulasiGetRequest struct {
	abstraction.Pagination
	spi_pbj.JenisRekapitulasiFilter
}

type JenisRekapitulasiGetResponses struct {
	Datas          *[]JenisRekapitulasiResponse
	PaginationInfo *abstraction.PaginationInfo
}
