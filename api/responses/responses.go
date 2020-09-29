// Package responses defines handy functions for encoding and writing data to client.
package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSON encodes json data and writes to client.
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
	}
}

// ERROR encodes errors as json and sends it to user.
func ERROR(w http.ResponseWriter, statusCode int, err error) {
	if err != nil {
		JSON(w, statusCode, struct {
			Error string `json:"error"`
		}{
			Error: err.Error(),
		})
		return
	}
}
