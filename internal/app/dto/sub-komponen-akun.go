package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Find
type SubKomponenAkunFindByThnAngIDAndSatkerIDRequest struct {
	ThnAngID int `json:"thnAngID"`
	SatkerID int `json:"satkerID"`
}

type SubKomponenAkunFindByThnAngIDAndSatkerIDResponse struct {
	abstraction.ID
	model.SubKomponenAkunEntity
	AkunCode string `json:"akun_code"`
	AkunName string `json:"akun_name"`
}
