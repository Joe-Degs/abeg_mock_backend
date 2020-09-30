//routes defines routes and registers them.
package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route describes a complete route.
type Route struct {
	Uri     string                                   // http request uri format eg: `/api/login`
	Method  string                                   // http request method
	Handler func(http.ResponseWriter, *http.Request) // function to handle requests
}

// expose the slice of routes.
func Load() []Route {
	return apiRoutes
}

// SetupRoutes takes a slice of routes and registers them to handler requests.
func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range apiRoutes {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
		// TODO
		// define middlewares for logging and other things you might
		// think about later.
	}
	return r
}
