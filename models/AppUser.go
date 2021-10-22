package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type AppUser struct {
	ID        uint      `gorm:"primaryKey"`
	Email     string    `json:"email" `
	Username  string    `json:"username" `
	Password  string    `json:"password"`
	HWID      string    `json:"hwid" gorm:"column:hwid"`
	AppID     string    `json:"app_id" `
	License   string    `json:"license" `
	Expiry    time.Time `json:"expiry" `
	Level     uint      `json:"level" `
	Banned    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (u AppUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required),
		validation.Field(&u.Password, validation.Required),
		validation.Field(&u.HWID, validation.Required),
		validation.Field(&u.AppID, validation.Required),
	)
}
