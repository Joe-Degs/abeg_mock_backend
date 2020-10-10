package models

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	PhoneNumber string  `gorm:"uniqueIndex"`
	Balance     float64 `gorm:"not null"`
}
