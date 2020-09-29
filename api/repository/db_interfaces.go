// Package repository defines interfaces for basic c.r.u.d of data.
package repository

import "github.com/Joe-Degs/abeg_mock_backend/api/models"

// UsersRepo defines the basic functionality of a db repo for a user.
type UsersRepo interface {
	// FindUser takes the phone number as a unique parameter
	// and uses it to find the user.
	FindUser(string) (models.User, error)

	// SaveUser create a new record in database with user details.
	SaveUser(models.User) (models.User, error)

	// UpdateUser takes a phonenumber(string) and a user(struct)
	// and updates the fields in the database.
	//UpdateUser(string, models.User) (int, error)

	// DeleteUser takes a phonenumber(string) and deletes
	// the user record associated with it.
	//DeleteUser(string) (int, error)
}

//
//type AccountsRepo interface {
//
//}
