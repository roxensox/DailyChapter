package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/roxensox/dailychapter/internal/api"
	"github.com/roxensox/dailychapter/internal/database"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Failed to open database")
		os.Exit(1)
	}

	cfg := api.ApiConfig{}
	dbQueries := database.New(db)
	cfg.DBConn = dbQueries
	sMux := http.NewServeMux()
	handler := http.FileServer(http.Dir("."))
	server := http.Server{
		Handler: sMux,
		Addr:    ":8080",
	}

	sMux.Handle("/app/", handler)
	sMux.HandleFunc("POST /users", cfg.POSTUsers)

	server.ListenAndServe()
}
