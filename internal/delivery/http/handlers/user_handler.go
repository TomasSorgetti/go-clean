package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users := []string{"Tomás", "Sofía", "Lucas"} 
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id, "name": "Tomás"})
}