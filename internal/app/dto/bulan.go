package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type BulanResponse struct {
	abstraction.ID
	model.BulanEntity
}

//Get
type BulanGetRequest struct {
	abstraction.Pagination
	model.BulanFilter
}

type BulanGetResponses struct {
	Datas          *[]BulanResponse
	PaginationInfo *abstraction.PaginationInfo
}
