package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint      `gorm:"primaryKey;autoIncrement"`
	FirstName      string    `gorm:"size:30"`
	LastName       string    `gorm:"size:30"`
	ProfilePicture *string   `gorm:"size:30"` //optional field, allows null value
	PhoneNumber    *string   `gorm:"size:10"` //optional field, allows null value
	DateOfBirth    time.Time //can be null
	Email          string    `gorm:"size:60"`  //optional field, allows null value
	Password       string    `gorm:"size:256"` //optional field, allows null value
	Status         bool      `gorm:"default:1"`
	CreatedAt      time.Time `gorm:"default:current_timestamp"` // Redefine CreatedAt
	UpdatedAt      time.Time `gorm:"default:current_timestamp"` // Redefine UpdatedAt
}
