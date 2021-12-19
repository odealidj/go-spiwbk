package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type GroupPackageValueResponse struct {
	abstraction.ID
	model.GroupPackageValueEntity
}

//Get
type GroupPackageValueGetRequest struct {
	abstraction.Pagination
	model.GroupPackageValueFilter
}

type GroupPackageValueGetResponses struct {
	Datas          *[]GroupPackageValueResponse
	PaginationInfo *abstraction.PaginationInfo
}
