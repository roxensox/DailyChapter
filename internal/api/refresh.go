package api

import (
	"encoding/json"
	"github.com/roxensox/dailychapter/internal/auth"
	"net/http"
	"time"
)

func (cfg *ApiConfig) POSTRefresh(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	rec := struct {
		RefreshToken string `json:"refresh_token"`
	}{}

	decoder := json.NewDecoder(req.Body)
	decoder.Decode(&rec)

	if rec.RefreshToken == "" {
		http.Error(writer, "No refresh token provided", http.StatusBadRequest)
		return
	}

	uid, err := cfg.DBConn.CheckRefreshToken(req.Context(), rec.RefreshToken)
	if err != nil {
		http.Error(writer, "No matching refresh token found", http.StatusNotFound)
		return
	}

	out, err := auth.MakeJWT(uid, cfg.PrivateKey, time.Hour*1)
	if err != nil {
		http.Error(writer, "Failed to create new access token", http.StatusInternalServerError)
		return
	}

	outObj := struct {
		token string `json:"token"`
	}{
		token: out,
	}
	outJson, err := json.Marshal(outObj)
	if err != nil {
		http.Error(writer, "Failed to marshal data", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(200)
	writer.Write(outJson)
}
