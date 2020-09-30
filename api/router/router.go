package router

import (
	"github.com/Joe-Degs/abeg_mock_backend/api/middlewares"
	"github.com/Joe-Degs/abeg_mock_backend/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.Use(
		middlewares.SetContentType("Content-Type", "application/json"),
	)
	return routes.SetupRoutes(r)
}
