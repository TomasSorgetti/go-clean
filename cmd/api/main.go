package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TomasSorgetti/go-clean/internal/config"
	httpdelivery "github.com/TomasSorgetti/go-clean/internal/delivery/http"
	_ "github.com/lib/pq"
)

func main() {
	connStr := config.GetDBConfig()

	db, err := InitDB(connStr)
	if err != nil {
		log.Fatal("No se pudo conectar a la base de datos")
	}
	defer db.Close()

	router := httpdelivery.NewRouter()

	fmt.Println("Servidor corriendo en http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println("Error iniciando el servidor:", err)
	}
}

func InitDB(connStr string) (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
		return nil, err
	}

	fmt.Println("Conexión a la base de datos exitosa")
	return db, nil
}
