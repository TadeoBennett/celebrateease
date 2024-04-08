package model

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	ID               uint      `gorm:"primaryKey;autoIncrement"`
	EventName        string    `gorm:"size:30;not null"`
	EventPicture     *string   `gorm:"size:30"` //optional field, allows null value
	EventDescription string    `gorm:"size:60;not null"`
	Status           bool      `gorm:"default:1"`
	CreatedAt        time.Time `gorm:"default:current_timestamp"` // Redefine CreatedAt
	UpdatedAt        time.Time `gorm:"default:current_timestamp"` // Redefine UpdatedAt
}
