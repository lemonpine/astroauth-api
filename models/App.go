package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type App struct {
	AppID   string `json:"app_id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"uniqueIndex"`
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

func (a App) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Name, validation.Required),
	)
}
