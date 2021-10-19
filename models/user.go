package models

import (
	"time"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `json:"email" `
	Password  string `json:"password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
