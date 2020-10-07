// Package controllers define functions that control how a route handles a request.
package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/Joe-Degs/abeg_mock_backend/api/database"
	"github.com/Joe-Degs/abeg_mock_backend/api/helpers/debug"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository/crud"
	"github.com/Joe-Degs/abeg_mock_backend/api/responses"
	"github.com/Joe-Degs/abeg_mock_backend/api/security"
	"gorm.io/gorm"
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
		dbUser, err := repo.FindUser(user.PhoneNumber)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		// check the values of dbUser and user
		debug.Pretty(dbUser)
		debug.Pretty(user)
		// check password is same as hashed one.
		err = security.Verify(dbUser.Password, user.Password)
		if err != nil {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("incorrect password"))
			return
		}
		responses.JSON(w, http.StatusOK, dbUser)
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

// FindFriends reads body which will be a slice(list) of phone_numbers.
// and checks if they are registered users, then sends back data
// containing info about found users.
func FindFriends(w http.ResponseWriter, r *http.Request) {
	friendsPhoneNumbers := struct {
		PhoneNumbers []string `json:"phone_numbers"`
	}{}

	// read request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// convert to json
	err = json.Unmarshal(body, &friendsPhoneNumbers)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	debug.Pretty(friendsPhoneNumbers) // print to console

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	repo := crud.NewUsersCrud(db)

	// loop through all the phone_numbers and find them in the database.
	// but i'm confused as to which should come first the func or loop.
	// should the function be in the loop or loop be in the function LOL!.
	func(repo repository.UsersRepo) {
		registeredFriends := make([]models.User, 0)
		for _, number := range friendsPhoneNumbers.PhoneNumbers {
			user, err := repo.FindUser(number)
			if err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					continue
				}
				responses.ERROR(w, http.StatusUnprocessableEntity, err)
				return
			}
			if user.ID != 0 {
				registeredFriends = append(registeredFriends, user)
			}
		}
		responses.JSON(w, http.StatusOK, registeredFriends)
	}(repo)
}

// UserIsRegistered checks if users data can be found in database and returns
// the user and a boolean value
func UserIsRegistered(phoneNumber string, w http.ResponseWriter) (*models.User, bool) {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return &models.User{}, false
	}
	repo := crud.NewUsersCrud(db)
	user, err := repo.FindUser(phoneNumber)
	if err != nil {
		if user.ID == 0 {
			responses.ERROR(w, http.StatusUnauthorized, errors.New("unregistered user"))
			return &user, false
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return &user, false
	}
	return &user, true
}

// UpLoadImage will handle uploading and storing image data of registered users.
// still dont know how to implement this yet, or the most efficient way to upload an image.
// i'll read around on the internet on the best way to do it.
// Also dont know how to struct models to include image yet.
// sending it via http POST with `content-type: multipart/form-data` seems like the smartest option
// for my usecase.
func UploadImage(w http.ResponseWriter, r *http.Request) {
	phoneNumber := r.FormValue("phone_number")

	// before reading imgFile and anything else, check if user is registered?
	if _, ok := UserIsRegistered(phoneNumber, w); ok {
		r.ParseMultipartForm(10 << 10)
		imgFile, imgHeader, err := r.FormFile("image")
		defer imgFile.Close()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		imgFormat := strings.Split(imgHeader.Header["Content-Type"][0], "/")[1]
		filename := fmt.Sprintf("./test_images/%s.%s", phoneNumber, imgFormat)
		file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		if _, err = io.Copy(file, imgFile); err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, struct {
			Message string `json:"message"`
		}{
			Message: "successful",
		})
	}

}

func UpdateImage(w http.ResponseWriter, r *http.Request) {
	responses.ERROR(w, http.StatusNotImplemented, errors.New("working on it"))
}
