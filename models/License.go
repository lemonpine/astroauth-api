package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type License struct {
	ID          uint   `gorm:"primaryKey"`
	AppID       string `json:"app_id" `
	License     string
	LengthHours uint `json:"length" `
	Level       uint `json:"level" `
	UsedBy      uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (l *License) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	l.License = uuid.NewString()
	return
}
