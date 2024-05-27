package handlers

import (
	"github.com/franciscof12/rest-api-thn/pkg/common"
	"net/http"
)

func MetricsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ip := r.URL.Query().Get("ip")
	count := 0

	mutex.Lock()
	for _, info := range RequestInfo {
		if info.RealIP == ip {
			count++
		}
	}
	mutex.Unlock()
	common.WriteJSON(w, http.StatusOK, map[string]interface{}{"ip": ip, "count": count})
}
