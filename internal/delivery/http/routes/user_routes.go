package routes

import (
	"github.com/TomasSorgetti/go-clean/internal/delivery/http/handlers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	userHandler := handlers.NewUserHandler()

	router.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")

	router.HandleFunc("/users/{id}", userHandler.GetUserByID).Methods("GET")
}
