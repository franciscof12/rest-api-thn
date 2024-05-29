package server

import (
	"github.com/franciscof12/rest-api-thn/internal/handlers"
	"log/slog"
	"net/http"
	"os"
)

func Start() {
	mux := http.NewServeMux()
	mux.HandleFunc("/feature", handlers.FeatureHandler)
	mux.HandleFunc("/metrics", handlers.MetricsHandler)
	slog.Info("Server started: ", "port", "8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		slog.Error("Error starting server:", "error", err.Error())
		os.Exit(1)
	}
}
