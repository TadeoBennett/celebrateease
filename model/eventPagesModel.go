package model

import "gorm.io/gorm"

type EventPage struct {
	gorm.Model
	Link string
}
