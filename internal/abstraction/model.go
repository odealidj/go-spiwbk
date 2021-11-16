package abstraction

import (
	"codeid-boiler/pkg/util/date"
	"time"

	"gorm.io/gorm"
)

type IDInc struct {
	ID uint16 `json:"id" validate:"number" gorm:"primaryKey;"`
}
type ID struct {
	ID uint16 `json:"id" param:"id" form:"id" query:"id" validate:"number" gorm:"primaryKey;autoIncrement:false;"`
}

type CreateBy struct {
	CreatedAt  time.Time  `json:"created_at" gorm:"<-:create"`
	CreatedBy  string     `json:"created_by"`
	ModifiedAt *time.Time `json:"modified_at" gorm:"<-:update"`
	ModifiedBy *string    `json:"modified_by"`

	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
type EntityInc struct {
	IDInc
	CreateBy
}

type Entity struct {
	ID
	CreateBy
}

type Filter struct {
	CreatedAt  *time.Time `query:"created_at"`
	CreatedBy  *string    `query:"created_by"`
	ModifiedAt *time.Time `query:"modified_at"`
	ModifiedBy *string    `query:"modified_by"`
}

func (m *Entity) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}

func (m *Entity) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}

func (m *EntityInc) BeforeUpdate(tx *gorm.DB) (err error) {
	m.ModifiedAt = date.DateTodayLocal()
	return
}

func (m *EntityInc) BeforeCreate(tx *gorm.DB) (err error) {
	m.CreatedAt = *date.DateTodayLocal()
	return
}
