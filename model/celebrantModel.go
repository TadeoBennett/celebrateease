package model

import (
	"time"

	"gorm.io/gorm"
)

type Celebrant struct {
	gorm.Model
	ID          uint `gorm:"primaryKey;autoIncrement"`
	Creator_id  uint
	FirstName   string     `gorm:"size:30;not null"`
	LastName    *string    `gorm:"size:30"`
	Email       *string    `gorm:"size:60"` //optional field, allows null value
	DateOfBirth *time.Time //can be null
	Status      bool       `gorm:"default:1"`
	CreatedAt   time.Time  `gorm:"default:current_timestamp"` // Redefine CreatedAt
	UpdatedAt   time.Time  `gorm:"default:current_timestamp"` // Redefine UpdatedAt
}
