package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

type JenisCertificateResponse struct {
	abstraction.ID
	model.JenisCertificateEntity
}

//Save
type JenisCertificateSaveRequest struct {
	model.JenisCertificateEntity
}

//Update
type JenisCertificateUpdateRequest struct {
	abstraction.ID
	model.JenisCertificateEntity
}

//Get
type JenisCertificateGetRequest struct {
	abstraction.Pagination
	model.JenisCertificateFilter
}

type JenisCertificateGetResponse struct {
	Datas          []JenisCertificateResponse
	PaginationInfo abstraction.PaginationInfo
}
