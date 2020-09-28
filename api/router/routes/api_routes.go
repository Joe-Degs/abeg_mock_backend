package routes

import "net/http"

// server's routes defined as slice of Route struct.
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
