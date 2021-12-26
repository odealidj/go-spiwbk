package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type ThnAngRequest struct {
	model.ThnAngEntity
}

type ThnAngRequests struct {
	Year string `json:"year"`
}

type ThnAngUpdateRequest struct {
	abstraction.ID
	model.ThnAngEntity
}

type ThnAngResponse struct {
	abstraction.ID
	model.ThnAngEntity
}

type ThnAngRequestForm struct {
	Year  string  `json:"year" form:"year"`
	Group []int16 `json:"group" form:"group"`
}

//Get
type ThnAngGetRequest struct {
	abstraction.ID
}

//Get
type ThnAngGetByYearRequest struct {
	Year string `json:"year"`
}

//Get All
type ThnAngGetAllRequest struct {
	abstraction.Pagination
	model.ThnAngFilter
}

type ThnAngGetAllResponse struct {
	Datas          *[]ThnAngResponse
	PaginationInfo *abstraction.PaginationInfo
}
