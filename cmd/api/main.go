package main

import (
	"fmt"
	"net/http"

	httpdelivery "github.com/TomasSorgetti/go-clean/internal/delivery/http"
)

func main() {
	router := httpdelivery.NewRouter()

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error iniciando el servidor:", err)
	}
}