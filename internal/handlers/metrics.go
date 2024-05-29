package handlers

import (
	"context"
	"github.com/franciscof12/rest-api-thn/internal/models"
	"github.com/franciscof12/rest-api-thn/pkg/common"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := r.URL.Query().Get("ip")
	if ip == "" {
		common.WriteJSON(w, http.StatusBadRequest, map[string]interface{}{"error": "ip parameter is required"})
		return
	}
	filter := bson.M{"realip": ip}
	count, err := models.ResquestCollection.CountDocuments(context.Background(), filter)
	if err != nil {
		common.WriteJSON(w, http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	common.WriteJSON(w, http.StatusOK, map[string]interface{}{"ip": ip, "count": count})
}
