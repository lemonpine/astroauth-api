package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type License struct {
	ID        uint   `gorm:"primaryKey"`
	AppID     string `json:"app_id" `
	License   string
	Length    uint `json:"length" `
	Level     uint `json:"level" `
	UsedBy    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (l *License) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	l.License = uuid.NewString()
	return
}

func (l License) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.AppID, validation.Required),
		validation.Field(&l.Length, validation.Required),
		validation.Field(&l.Level, validation.Required),
	)
}
