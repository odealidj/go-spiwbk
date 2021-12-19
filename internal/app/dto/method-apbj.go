package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type MethodApbjResponse struct {
	abstraction.ID
	model.MethodApbjEntity
}

//Get
type MethodApbjGetRequest struct {
	abstraction.Pagination
	model.MethodApbjFilter
}

type MethodApbjGetResponses struct {
	Datas          *[]MethodApbjResponse
	PaginationInfo *abstraction.PaginationInfo
}
