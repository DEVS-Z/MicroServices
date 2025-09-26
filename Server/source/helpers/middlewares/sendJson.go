package middleware

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func sendJSON(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(status))

	response := ErrorResponse{
		Status: status,
		Error:  message,
	}

	json.NewEncoder(w).Encode(response)
}
