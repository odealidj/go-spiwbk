package dto

import (
	"codeid-boiler/internal/abstraction"
	"codeid-boiler/internal/app/model"
)

//Find ByThnAngIDAndSatkerID
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

//Find ByKomponenID
type SubKomponenAkunFindByKomponenIDRequest struct {
	KomponenID int `json:"komponenID"`
}

type SubKomponenAkunFindByKomponenIDResponse struct {
	abstraction.ID
	model.SubKomponenAkunEntity
	AkunCode string `json:"akun_code"`
	AkunName string `json:"akun_name"`
}
