// Package controllers define functions that control how a route handles a request.
package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/models"
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

	userform := NewForm(user.PhoneNumber)
	if err := userform.Get(); err != nil {
		if errors.Is(err, ErrUnregisteredUser) {
			responses.ERROR(w, http.StatusUnauthorized, err)
			return
		} else if errors.Is(err, ErrDbConnection) {
			responses.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if err := userform.Validate(user.Password); err != nil {
		responses.ERROR(w, http.StatusForbidden, err)
		return
	}
	responses.JSON(w, http.StatusOK, *userform.UserModel)
}
