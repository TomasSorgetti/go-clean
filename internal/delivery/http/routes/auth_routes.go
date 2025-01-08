package routes

import (
	"github.com/TomasSorgetti/go-clean/internal/delivery/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterAuthRoutes(router *mux.Router) {
	authHandler := handlers.NewAuthHandler()

	router.HandleFunc("/auth/login", authHandler.Login).Methods("POST")

	router.HandleFunc("/auth/register", authHandler.Register).Methods("POST")
}
