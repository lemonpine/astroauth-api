package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type App struct {
	AppID   string `json:"app_id" gorm:"primaryKey"`
	Name    string `json:"name" `
	OwnedBy uint
	// User      AppUser `gorm:"foreignKey:OwnedBy"`
	// Banned    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (a *App) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	a.AppID = uuid.NewString()
	return
}
