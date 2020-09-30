package models

import (
	"github.com/Joe-Degs/abeg_mock_backend/api/security"
	"gorm.io/gorm"
)

// User represents a single user who will own an account.
type User struct {
	gorm.Model
	FullName    string `gorm:"unique;not null" json:"full_name"`
	PhoneNumber string `gorm:"unique;not null" json:"phone_number"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"-"`
}

// BeforeSave makes sure to securely hash password before saving in database.
func (u *User) BeforeSave(*gorm.DB) (err error) {
	hashedPass, err := security.Hash(u.Password)
	if err != nil {
		return
	}
	u.Password = string(hashedPass)
	return
}
