package models

import (
	"github.com/Joe-Degs/abeg_mock_backend/api/security"
	"gorm.io/gorm"
)

// User represents a single user who will own an account.
type User struct {
	gorm.Model
	FullName    string   `gorm:"unique;not null" json:"full_name"`
	PhoneNumber string   `gorm:"uniqueIndex;not null" json:"phone_number"`
	Email       string   `gorm:"unique;not null" json:"email"`
	Password    string   `gorm:"not null" json:"password"`
	UserData    UserData `gorm:"foreignkey:PhoneNumber"`
	Account     Account  `gorm:"foreignkey:PhoneNumber"`
}

type UserData struct {
	gorm.Model
	PhoneNumber   string    `gorm:"uniqueIndex"`
	Country       string    `gorm:"size:50"`
	UserName      string    `gorm:"unique;size:50"`
	Gender        string    `gorm:"size:17"` // prefer not to say is valid
	ImageFileName string    `gorm:"unique;size:14;"`
	NextOfKin     NextOfKin `gorm:"foreignKey:PhoneNumber"`
}

type NextOfKin struct {
	gorm.Model
	PhoneNumber string `gorm:"uniqueIndex"`
	FullName    string `gorm:"size:50;"`
	Gender      string `gorm:"size:17"`
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

// models dont seem to work as i expected them to and thats because i quite dont understand
// how to to model relationships between different gorm models.
// my expectations are to create a user_table, user_data table, nextofkin table and an account table.
// the user_table will have a unique index on its phone_number field which will then serve
// as a foreignkey reference in all related tables. Not to forget, the id(which is a primary_key)
// field of the user_table will serve as a foreignkey in the nextofkin table as nextofkins are
// required to have their own phone_numbers.
// The tables user_data and account will have a phone_number field that will serve as as a foreignkey
// reference to the phone_number field in the user_table.
