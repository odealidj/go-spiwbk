package abstractions

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID int `json:"id" gorm:"primaryKey;autoIncrement;"`

	CreatedAt  time.Time `json:"created_at"`
	CreatedBy  string    `json:"created_by"`
	ModifiedAt time.Time `json:"modified_at"`
	ModifiedBy string    `json:"modified_by"`

	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
