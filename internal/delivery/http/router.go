package http

import (
	"github.com/TomasSorgetti/go-clean/internal/delivery/http/routes"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	routes.RegisterAuthRoutes(router)

	routes.RegisterUserRoutes(router)

	return router
}