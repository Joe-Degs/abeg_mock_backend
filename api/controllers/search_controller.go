package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/responses"
)

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

	foundFriends, err := GetUsers(friendsPhoneNumbers.PhoneNumbers)
	if err != nil {
		if errors.Is(err, ErrNoFriend) {
			responses.JSON(w, http.StatusOK, err)
			return
		}
		responses.ERROR(w, http.StatusInternalServerError, err)
	}
	responses.JSON(w, http.StatusOK, foundFriends)
}
