package spi_pbj

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model/spi-pbj"
)

type GroupPackageValueResponse struct {
	abstraction.ID
	spi_pbj.GroupPackageValueEntity
}

//Get
type GroupPackageValueGetRequest struct {
	abstraction.Pagination
	spi_pbj.GroupPackageValueFilter
}

type GroupPackageValueGetResponses struct {
	Datas          *[]GroupPackageValueResponse
	PaginationInfo *abstraction.PaginationInfo
}
