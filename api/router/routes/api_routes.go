package routes

import (
	"net/http"

	"github.com/Joe-Degs/abeg_mock_backend/api/controllers"
)

// represents all of the servers public routes.
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
	Route{
		Uri:     "/api/findfriends",
		Method:  http.MethodPost,
		Handler: controllers.FindFriends,
	},
	Route{
		Uri:     "/api/uploadimage",
		Method:  http.MethodPost,
		Handler: controllers.UpLoadImage,
	},
}
