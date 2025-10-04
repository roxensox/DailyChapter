package api

import (
	"encoding/json"
	"github.com/roxensox/dailychapter/internal/auth"
	"net/http"
	"time"
)

func (cfg *ApiConfig) POSTLogin(writer http.ResponseWriter, req *http.Request) {
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
	match, err := auth.CheckPasswordHash(rcv.Password, resp.HashedPassword)
	if err != nil {
		http.Error(writer, "Failed to check password", http.StatusInternalServerError)
		return
	}
	if !match {
		http.Error(writer, "Incorrect password", http.StatusUnauthorized)
		return
	}
	token, err := auth.MakeJWT(resp.ID, cfg.Secret, time.Hour)
	if err != nil {
		http.Error(writer, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	out := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}
	outJson, err := json.Marshal(out)
	if err != nil {
		http.Error(writer, "Failed to marshal data", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(200)
	writer.Write(outJson)
}
