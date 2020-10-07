package models

import (
	"github.com/Joe-Degs/abeg_mock_backend/api/security"
	"gorm.io/gorm"
)

// User represents a single user who will own an account.
type User struct {
	gorm.Model
	FullName    string `gorm:"unique;not null" json:"full_name"`
	PhoneNumber string `gorm:"uniqueIndex;not null" json:"phone_number"`
	Email       string `gorm:"unique;not null" json:"email"`
	Password    string `gorm:"not null" json:"password"`
}

type UserData struct {
	PhoneNumber   string `gorm:"primarykey;size:20;not null"`
	Country       string `gorm:"not null"`
	UserName      string `gorm:"not null"`
	Gender        string `gorm:"size:17"` // prefer not to say is valid
	ImageFileName string `gorm:"unique;size:14;"`
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
