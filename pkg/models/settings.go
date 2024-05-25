package models

import (
	"gorm.io/gorm"
)

type Settings struct {
	gorm.Model
	ID                 uint `gorm:"primaryKey;autoIncrement"`
	User_id            uint //foreign key
	EmailNotifications bool `gorm:"default:0"`
	DarkMode           bool `gorm:"default:0"`
}
