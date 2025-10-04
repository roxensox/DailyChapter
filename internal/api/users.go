package api

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/roxensox/dailychapter/internal/auth"
	"github.com/roxensox/dailychapter/internal/database"
	"github.com/roxensox/dailychapter/internal/utils"
	"net/http"
	"time"
)

func (cfg *ApiConfig) POSTUsers(writer http.ResponseWriter, req *http.Request) {
	// Handles POST request to users endpoint, creates new user

	// Writes header content type
	writer.Header().Set("Content-Type", "application/json")

	// Instantiates a custom struct to receive input
	rcv := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&rcv)
	if !utils.ValidateEmail(rcv.Email) {
		http.Error(writer, "Invalid email", http.StatusBadRequest)
		return
	}

	hashPass, err := auth.HashPassword(rcv.Password)

	// Creates parameter object for query
	params := database.CreateUserParams{
		Email:          rcv.Email,
		ID:             uuid.New(),
		HashedPassword: hashPass,
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
	}

	// Runs the query, rejects duplicate users
	resp, err := cfg.DBConn.CreateUser(req.Context(), params)
	if err != nil {
		http.Error(writer, "Unable to create user.", http.StatusInternalServerError)
		return
	}

	// Instantiates a custom struct to transform db response to JSON
	output := struct {
		Email     string    `json:"email"`
		ID        uuid.UUID `json:"id"`
		CreatedAt time.Time `json:"created_at"`
	}{
		Email:     resp.Email,
		ID:        resp.ID,
		CreatedAt: resp.CreatedAt,
	}

	// Marshals output struct to JSON
	outputJson, err := json.Marshal(output)
	if err != nil {
		http.Error(writer, "Unable to marshal data.", http.StatusInternalServerError)
		return
	}

	// Writes success message
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(outputJson)
}

func (cfg *ApiConfig) GETUsers(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	http.Error(writer, "Not implemented yet", http.StatusNotImplemented)
}
