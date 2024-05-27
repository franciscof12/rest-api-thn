package common

import (
	"encoding/json"
	"log/slog"
	"net/http"
)

func WriteJSON(response http.ResponseWriter, statusCode int, data any) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	if err := json.NewEncoder(response).Encode(data); err != nil {
		slog.Error("Error writing response", "error: ", err.Error())
	}
}

func WriteError(response http.ResponseWriter, statusCode int, message string) {
	WriteJSON(response, statusCode, map[string]string{"error": message})
}
