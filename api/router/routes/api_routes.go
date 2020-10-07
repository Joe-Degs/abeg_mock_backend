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
		Uri:     "/api/search",
		Method:  http.MethodPost,
		Handler: controllers.FindFriends,
	},
	Route{
		Uri:     "/api/image",
		Method:  http.MethodPost,
		Handler: controllers.UploadImage,
	},
	Route{
		Uri:     "/api/image",
		Method:  http.MethodPut,
		Handler: controllers.UpdateImage,
	},
}
