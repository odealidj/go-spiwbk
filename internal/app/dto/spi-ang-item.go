package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Save
type SpiAngItemSaveRequest struct {
	SatkerID int `json:"satkerID"`
	ThnAngID int `json:"thnAngID"`
}

type SpiAngItemResponse struct {
	abstraction.ID
	model.SpiAngItemEntity
}
