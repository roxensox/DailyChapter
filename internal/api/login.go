package api

import (
	"encoding/json"
	"github.com/roxensox/dailychapter/internal/auth"
	"net/http"
	"time"
)

func (cfg *ApiConfig) POSTLogin(writer http.ResponseWriter, req *http.Request) {
	// Handles POST request to /login endpoint
	//TODO: If a refresh token was provided, check if it's valid. If it isn't,
	//	revoke any active refresh token for this user and create a new one, return
	//	the new refresh token and access token.

	writer.Header().Set("Content-Type", "application/json")

	rcv := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	json.NewDecoder(req.Body).Decode(&rcv)
	resp, err := cfg.DBConn.GetUserByEmail(req.Context(), rcv.Email)
	if err != nil {
		http.Error(writer, "User not found", http.StatusNotFound)
		return
	}

	if rcv.Password == "" {
		http.Error(writer, "Must provide a password", http.StatusBadRequest)
		return
	}

	match, err := auth.CheckPasswordHash(rcv.Password, resp.HashedPassword)
	if err != nil {
		http.Error(writer, "Failed to check password", http.StatusInternalServerError)
		return
	}
	if !match {
		http.Error(writer, "Incorrect password", http.StatusUnauthorized)
		return
	}
	access_token, err := auth.MakeJWT(resp.ID, cfg.PrivateKey, time.Hour)
	if err != nil {
		http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	out := struct {
		RefreshToken string `json:"refresh_token"`
		AccessToken  string `json:"access_token"`
	}{
		AccessToken: access_token,
	}
	outJson, err := json.Marshal(out)
	if err != nil {
		http.Error(writer, "Failed to marshal data", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(200)
	writer.Write(outJson)
}
