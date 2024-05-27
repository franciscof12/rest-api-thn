package main

import (
	"github.com/franciscof12/rest-api-thn/internal/handlers"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /feature", handlers.FeatureHandler)
	mux.HandleFunc("GET /metrics", handlers.MetricsHandler)
	slog.Info("Server started: ", "port", "8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Error starting server:", "error", err.Error())
	}
}
