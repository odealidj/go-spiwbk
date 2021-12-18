package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisRekapitulasiResponse struct {
	abstraction.ID
	model.JenisRekapitulasiEntity
}

//Get
type JenisRekapitulasiGetRequest struct {
	abstraction.Pagination
	model.JenisRekapitulasiFilter
}

type JenisRekapitulasiGetResponses struct {
	Datas          *[]JenisRekapitulasiResponse
	PaginationInfo *abstraction.PaginationInfo
}
