// Package controllers define functions that control how a route handles a request.
package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/database"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository/crud"
	"github.com/Joe-Degs/abeg_mock_backend/api/responses"
)

// Login controls the login route of server
func Login(w http.ResponseWriter, r *http.Request) {
	// read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// convert json body to user struct.
	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// connect to db
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewUsersCrud(db)
	func(repo repository.UsersRepo) {
		user, err := repo.FindUser(user.PhoneNumber)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, user)
	}(repo)
}

// Signup controls the signup route of server
func Signup(w http.ResponseWriter, r *http.Request) {
	// read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// convert json body to user struct.
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// connect to db and try saving data in it.
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	// new users_crud repo to save the data.
	repo := crud.NewUsersCrud(db)
	func(repo repository.UsersRepo) {
		user, err := repo.SaveUser(user)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusCreated, user) // feels unnecessary to send user back to client.
	}(repo)
}

// FindFriends reads body which will be a slice of phone_numbers.
// and checks if they are registered users, then sends back data
// containing info about found users.
func FindFriends(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusNotImplemented, errors.New("come back later :-)"))
}

// UpLoadImage will handle uploading and storing image data of registered users.
// still dont know how to implement this yet, or the most efficient way to upload an image.
// i'll read around on the internet on the best way to do it.
//
func UpLoadImage(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusNotImplemented, errors.New("come back at a longer later :-("))
}
