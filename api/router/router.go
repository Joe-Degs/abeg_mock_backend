package router

import (
	"log"
	"os"

	"github.com/Joe-Degs/abeg_mock_backend/api/middlewares"
	"github.com/Joe-Degs/abeg_mock_backend/api/router/routes"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	r.Use(
		middlewares.SetContentType("Content-Type", "application/json"),
		middlewares.Logging(logger),
	)
	return routes.SetupRoutes(r)
}
