package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type SatkerResponse struct {
	abstraction.ID
	model.SatkerEntity
}

//Save
type SatkerSaveRequest struct {
	model.SatkerEntity
}

//Update
type SatkerUpdateRequest struct {
	abstraction.ID
	model.SatkerEntity
}

//Delete
type SatkerDeleteRequest struct {
	abstraction.ID
}

//Get
type SatkerGetRequest struct {
	abstraction.Pagination
	model.SatkerFilter
}

//GetByID
type SatkerGetByIDRequest struct {
	abstraction.ID
}

type SatkerGet2Request struct {
	abstraction.PaginationArr
	model.SatkerFilter
}

type SatkerGetResponse struct {
	Datas          []SatkerResponse
	PaginationInfo abstraction.PaginationInfo
}

type SatkerGet2Response struct {
	Datas          []SatkerResponse
	PaginationInfo abstraction.PaginationInfoArr
}
