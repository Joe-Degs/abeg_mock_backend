package routes

import (
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/controllers"
)

// server's routes declared as slice of Route structs.
var apiRoutes = []Route{
	Route{
		Uri:     "/api/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
	},
	Route{
		Uri:     "/api/signup",
		Method:  http.MethodPost,
		Handler: controllers.Signup,
	},
}
