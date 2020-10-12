package controllers

/*
   Code in this file is an attempt to add a layer of abstraction
   to activities involving databases and users.
   Making it less verbose to query database when the need arises in
   the controller functions handling requests and giving responses.
*/

import (
	"errors"
	"fmt"

	"github.com/Joe-Degs/abeg_mock_backend/api/database"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository/crud"
	"github.com/Joe-Degs/abeg_mock_backend/api/security"
	"gorm.io/gorm"
)

// UserForm holds data to make querying user table easier.
type UserForm struct {
	Number         string               // number is a phone_number from user table
	UserModel      *models.User         // points to user that has number in user_table
	repoConnection repository.UsersRepo // interface for querying database easily.
	// kept here to save me from opening multiple connections to database.
}

var (
	ErrUnregisteredUser  error = errors.New("Unregistered User")
	ErrDbConnection            = errors.New("error connecting to database")
	ErrNoFriend                = errors.New("non of your friends are users")
	ErrIncorrectPassword       = errors.New("Incorrect Password")
	ErrUserRegistered          = errors.New("User already registered")
)

// NewForm returns a new UserForm with given phoneNumber.
func NewForm(phoneNumber string) *UserForm {
	return &UserForm{Number: phoneNumber}
}

// openConnection connects to database and returns a repo for querying the database.
func openConnection() (repository.UsersRepo, error) {
	db, err := database.Connect()
	if err != nil {
		return nil, ErrDbConnection
	}
	repo := crud.NewUsersCrud(db)
	return repo, nil
}

// open stores a repo on UserForm for reuse.
func (u *UserForm) open() error {
	repo, err := openConnection()
	if err != nil {
		return err
	}
	u.repoConnection = repo
	return nil
}

// Close sets repo on UserForm to nil.
func (u *UserForm) Close() {
	u.repoConnection = nil
	fmt.Printf("Repo is supposed to be nil, what did i get? %v", u.repoConnection)
}

// Get opens connection to database and retrieves user associated with Number field.
func (u *UserForm) Get() (error, bool) {
	err := u.open()
	if err != nil {
		return err, false
	}
	user, err := u.repoConnection.FindUser(u.Number)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrUnregisteredUser, false
		}
		return err, false
	}
	u.UserModel = &user
	return nil, true
}

// RegisterNewUser stores a new user in database.
func (u *UserForm) RegisterNewUser(user models.User) error {
	if _, ok := u.Get(); ok {
		return ErrUserRegistered
	}
	defer u.Close()
	saved, err := u.repoConnection.SaveUser(user)
	if err != nil {
		return err
	}
	u.UserModel = &saved
	return nil
}

// Validate compares a password and a hash and sees.
// this method is only called after calling Get.
func (u *UserForm) Validate(password string) error {
	err := security.Verify(u.UserModel.Password, password)
	if err != nil {
		return ErrIncorrectPassword
	}
	return err
}

func fetch(repo repository.UsersRepo, number string) (*models.User, error) {
	user, err := repo.FindUser(number)
	if err != nil {
		// some type of logging might happen here
		return nil, err
	}
	return &user, nil
}

func GetUsers(numbers []string) ([]models.User, error) {
	//return users, errors.New("")
	repo, err := openConnection()
	if err != nil {
		return nil, err
	}
	users := make([]models.User, 0)
	for _, number := range numbers {
		user, _ := fetch(repo, number)
		if user != nil {
			users = append(users, *user)
			continue
		}
	}
	if len(users) < 1 {
		return nil, ErrNoFriend
	}
	return users, nil
}
