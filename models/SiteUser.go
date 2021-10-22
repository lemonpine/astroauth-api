package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type SiteUser struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `json:"email" `
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u SiteUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required),
	)
}
