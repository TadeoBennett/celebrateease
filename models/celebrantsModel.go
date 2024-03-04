package models

import "gorm.io/gorm"

type Celebrant struct {
	gorm.Model
	FirstName string
	LastName string
}