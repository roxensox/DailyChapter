package api

import (
	"database/sql"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/roxensox/dailychapter/internal/auth"
	"github.com/roxensox/dailychapter/internal/database"
	"net/http"
	"time"
)

func (cfg *ApiConfig) POSTBooks(writer http.ResponseWriter, req *http.Request) {
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil || apiKey != cfg.APIKey {
		http.Error(writer, "Invalid API Key", http.StatusUnauthorized)
		return
	}
	rcv := struct {
		Title   string    `json:"title"`
		PubDate time.Time `json:"pub_date"`
	}{}

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&rcv)

	if err != nil {
		http.Error(writer, "Unable to parse date", http.StatusInternalServerError)
		return
	}

	pDate := sql.NullTime{}

	if rcv.PubDate.IsZero() {
		pDate.Time = rcv.PubDate
		pDate.Valid = true
	}

	params := database.CreateBookParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Title:     rcv.Title,
		PubDate:   pDate,
	}

	cfg.DBConn.CreateBook(req.Context(), params)
}
