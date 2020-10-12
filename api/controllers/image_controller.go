package controllers

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Joe-Degs/abeg_mock_backend/api/database"
	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/repository/crud"
	"github.com/Joe-Degs/abeg_mock_backend/api/responses"
	"gorm.io/gorm"
)

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
		if errors.Is(err, gorm.ErrRecordNotFound) {
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
		filename := fmt.Sprintf("%s.%s", phoneNumber, imgFormat)
		file, err := os.OpenFile("./test_images/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
		defer file.Close()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		if _, err = io.Copy(file, imgFile); err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		//TODO
		// if copy successful, save filename to database

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
