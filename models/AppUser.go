package models

import (
	"time"
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
