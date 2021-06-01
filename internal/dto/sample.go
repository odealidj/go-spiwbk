package dto

import "code-boiler/internal/abstractions"

type Sample struct {
	Key    string `json:"key"`
	Value  string `json:"value"`
	UserId int    `json:"user_id"`
}

//region Get
type SampleGetRequest struct {
	abstractions.Pagination
	Filter SampleGetFilterRequest
}
type SampleGetFilterRequest struct {
	Key    string `json:"key,omitempty" query:"key"`
	Value  string `json:"value,omitempty" query:"value"`
	UserId int    `json:"user_id,omitempty" query:"user_id"`
}

//endregion

//region SampleGetByID
type SampleGetByIDRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type SampleGetByIDResponse struct {
	abstractions.Model
	Sample
}

//endregion

//region SampleStore
type SampleStoreRequest struct {
	Key    string `json:"key" validate:"required"`
	Value  string `json:"value" validate:"required"`
	UserId int    `json:"user_id" validate:"required"`
}

type SampleStoreResponse struct {
	abstractions.Model
	Sample
}

//endregion

//region SampleUpdate
type SampleUpdateRequest struct {
	ID     int    `param:"id" validate:"required,numeric"`
	Key    string `json:"key,omitempty"`
	Value  string `json:"value,omitempty"`
	UserId int    `json:"user_id,omitempty"`
}

type SampleUpdateResponse struct {
	abstractions.Model
	Sample
}

//endregion

//region SampleUpdate
type SampleDeleteRequest struct {
	ID int `param:"id" validate:"required,numeric"`
}

type SampleDeleteResponse struct {
	abstractions.Model
	Sample
}

//endregion
