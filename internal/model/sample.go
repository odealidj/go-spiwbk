package model

import (
	"code-boiler/internal/abstractions"
	"time"

	"gorm.io/gorm"
)

type Sample struct {
	//model standart design
	abstractions.Model

	Key   string `json:"key"`
	Value string `json:"value"`

	//relations
	UserId int `json:"user_id"`

	//relations definitions
	User User `json:"user" gorm:"foreignKey:UserId"`
}

func (m *Sample) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = time.Now()
	return
}
