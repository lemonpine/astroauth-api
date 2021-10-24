package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"gorm.io/gorm"
)

type SiteUser struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `json:"email" gorm:"uniqueIndex"`
	Password  string `json:"password"`
	MaxApp    uint   `json:"max_app"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SiteUserRequest struct {
	Email    string `json:"email" gorm:"uniqueIndex"`
	Password string `json:"password"`
}

func (u *SiteUser) BeforeCreate(tx *gorm.DB) (err error) {
	// UUID version 4
	u.MaxApp = 1
	return
}

func (u SiteUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
	)
}
