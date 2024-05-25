package models

import (
	"time"

	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	ID              uint      `gorm:"primaryKey;autoIncrement"`
	Page_creator_id uint      //is a foreign key
	Event_id        uint      //is a foreign key
	Title           string    `gorm:"size:30;not null"`
	HeaderImage     *string   `gorm:"size:30"` //optional field, allows null value
	ContentImage1   *string   `gorm:"size:30"` //optional field, allows null value
	VoiceNote       *[]byte   //restrict size in the application logic //optional field, allows null value
	BodyTextContent string    `gorm:"size:300;not null"`
	ViewCode        string    `gorm:"size:30;not null"`
	CongratsGIF     bool      `gorm:"default:0"`
	LoveGIF         bool      `gorm:"default:0"`
	Status          bool      `gorm:"default:0"`
	Accessed        uint      `gorm:"default:0"`
	CreatedAt       time.Time `gorm:"default:current_timestamp"` // Redefine CreatedAt
	UpdatedAt       time.Time `gorm:"default:current_timestamp"` // Redefine UpdatedAt
}
