package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisBelanjaPaguResponse struct {
	abstraction.ID
	model.JenisBelanjaPaguEntity
}

//Get
type JenisBelanjaPaguGetRequest struct {
	abstraction.Pagination
	model.JenisBelanjaPaguFilter
}

type JenisBelanjaPaguGetResponses struct {
	Datas          *[]JenisBelanjaPaguResponse
	PaginationInfo *abstraction.PaginationInfo
}
