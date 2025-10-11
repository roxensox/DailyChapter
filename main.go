package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/roxensox/dailychapter/internal/api"
	"github.com/roxensox/dailychapter/internal/database"
	"github.com/roxensox/dailychapter/internal/utils"
)

func main() {
	// Loads the .env file into the enviroment
	godotenv.Load()

	// Pulls the database url from the environment
	dbURL := os.Getenv("DB_URL")

	// Opens a DB connection
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println("Failed to open database")
		os.Exit(1)
	}

	// Creates config object
	cfg := api.ApiConfig{}

	// Starts query engine and adds it to config
	dbQueries := database.New(db)
	cfg.DBConn = dbQueries
	cfg.Secret = os.Getenv("JWT_SECRET")
	cfg.APIKey = os.Getenv("API_KEY")
	cfg.PrivateKey = utils.GetPrivateKey()
	cfg.PublicKey = utils.GetPublicKey()

	// Creates a new serve mux instance
	sMux := http.NewServeMux()

	// Creates server instance
	server := http.Server{
		Handler: sMux,
		Addr:    ":8080",
	}

	// Registers function handlers for POST methods
	sMux.HandleFunc("POST /users", cfg.POSTUsers)
	sMux.HandleFunc("POST /reset", cfg.POSTReset)
	sMux.HandleFunc("POST /login", cfg.POSTLogin)
	sMux.HandleFunc("POST /books", cfg.POSTBooks)
	sMux.HandleFunc("POST /books/{bookID}/subscribe", cfg.POSTBooksIDSubscribe)
	sMux.HandleFunc("POST /refresh", cfg.POSTRefresh)

	// Registers function handlers for GET methods
	sMux.HandleFunc("GET /users", cfg.GETUsers)
	sMux.HandleFunc("GET /books", cfg.GETBooks)

	// Runs the server loop
	server.ListenAndServe()
}
