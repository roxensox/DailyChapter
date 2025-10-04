package api

import (
	"net/http"
)

func (cfg *ApiConfig) POSTReset(writer http.ResponseWriter, req *http.Request) {
	err := cfg.DBConn.Reset(req.Context())
	if err != nil {
		http.Error(writer, "Failed to reset DB", http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(204)
}
