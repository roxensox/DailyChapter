package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/roxensox/dailychapter/internal/auth"
	"github.com/roxensox/dailychapter/internal/database"
	"github.com/roxensox/dailychapter/internal/utils"
)

func (cfg *ApiConfig) POSTBooks(writer http.ResponseWriter, req *http.Request) {
	// Gets APIKey from the header and checks it against config
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil || apiKey != cfg.APIKey {
		http.Error(writer, fmt.Sprintf("Invalid API Key: %s", apiKey), http.StatusUnauthorized)
		return
	}

	// Prepares struct instance to receive request body
	rcv := struct {
		Title   string `json:"title"`
		PubDate string `json:"pub_date"`
	}{}

	// Decodes body into rcv
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&rcv)

	parsedDate, err := utils.ParseDate(rcv.PubDate)

	// Initializes nulltime for query parameter
	pDate := sql.NullTime{}

	// Sets up nulltime if no date was found
	if err == nil {
		pDate.Time = parsedDate
		pDate.Valid = true
	}

	// Sets up parameters
	params := database.CreateBookParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     rcv.Title,
		PubDate:   pDate,
	}

	// Adds the book to the DB
	_, err = cfg.DBConn.CreateBook(req.Context(), params)
	if err != nil {
		http.Error(writer, "Failed to add book", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(200)
}

func (cfg *ApiConfig) GETBooks(writer http.ResponseWriter, req *http.Request) {
	// Handler for the /books endpoint, returns all books

	// Sets response content type to json
	writer.Header().Set("Content-Type", "application/json")

	// Queries the DB for all books
	results, err := cfg.DBConn.GetBooks(req.Context())
	if err != nil || len(results) == 0 {
		http.Error(writer, "No books found.", http.StatusNotFound)
		return
	}

	// Prepares a custom object for output
	out := []struct {
		Title   string    `json:"title"`
		PubDate time.Time `json:"pub_date"`
	}{}

	// Loops through results and parses them into output object
	for _, book := range results {
		curr_struct := struct {
			Title   string    `json:"title"`
			PubDate time.Time `json:"pub_date"`
		}{
			Title:   book.Title,
			PubDate: book.PubDate.Time,
		}
		out = append(out, curr_struct)
	}

	// Marshals output to json
	jsonOut, err := json.Marshal(out)
	if err != nil {
		http.Error(writer, "Failed to marshal data.", http.StatusInternalServerError)
		return
	}

	// Writes success response
	writer.WriteHeader(200)
	writer.Write(jsonOut)
}

func (cfg *ApiConfig) POSTBooksIDSubscribe(writer http.ResponseWriter, req *http.Request) {

	writer.Header().Set("Content-Type", "application/json")

	book_id := req.PathValue("bookID")
	book_uuid, err := uuid.Parse(book_id)
	if err != nil {
		fmt.Printf("Invalid book ID: %s\n\t%v\n", book_id, err)
		http.Error(writer, "Invalid book ID", http.StatusBadRequest)
		return
	}

	token, err := auth.GetJWT(req.Header)
	if err != nil {
		fmt.Printf("Failed to retrieve JWT\n\t%v\n", err)
		http.Error(writer, "JWT not found", http.StatusBadRequest)
		return
	}
	uid, err := auth.ValidateJWT(token, cfg.PublicKey)
	if err != nil {
		fmt.Println(err)
		http.Error(writer, "Invalid token", http.StatusUnauthorized)
		return
	}

	params := database.SubscribeParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    uid,
		BookID:    book_uuid,
	}
	err = cfg.DBConn.Subscribe(req.Context(), params)
	if err != nil {
		http.Error(writer, "Failed to execute query", http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	writer.WriteHeader(204)
}
