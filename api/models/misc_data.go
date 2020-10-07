package models

import "gorm.io/gorm"

// Misc holds other user credentials that might be provided by user.
type Misc struct {
	gorm.Model
	Tag   string `gorm:"unique;size:20;not null" json:"tag"`
	Image []byte `json:"image"`
}
