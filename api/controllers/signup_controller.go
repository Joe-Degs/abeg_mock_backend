package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/models"
	"github.com/Joe-Degs/abeg_mock_backend/api/responses"
)

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
	//	db, err := database.Connect()
	//	if err != nil {
	//		responses.ERROR(w, http.StatusInternalServerError, err)
	//		return
	//	}
	//	// new users_crud repo to save the data.
	//	repo := crud.NewUsersCrud(db)
	//	func(repo repository.UsersRepo) {
	//		user, err := repo.SaveUser(user)
	//		if err != nil {
	//			responses.ERROR(w, http.StatusUnprocessableEntity, err)
	//			return
	//		}
	//		responses.JSON(w, http.StatusCreated, user) // feels unnecessary to send user back to client.
	//	}(repo)

	userform := NewForm(user.PhoneNumber)
	if err = userform.RegisterNewUser(user); err != nil {
		if errors.Is(err, ErrUserRegistered) {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	responses.JSON(w, http.StatusCreated, *(userform.UserModel))
}
