package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/franciscof12/thn-rest-api/internal/models"
)

func TestMetricsHandler(t *testing.T) {
	mutex.Lock()
	RequestInfo = []models.RequestInfo{
		{RealIP: "127.0.0.1"},
		{RealIP: "127.0.0.1"},
		{RealIP: "192.168.0.1"},
	}
	mutex.Unlock()

	req, err := http.NewRequest("GET", "/metrics?ip=127.0.0.1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(MetricsHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := map[string]interface{}{"ip": "127.0.0.1", "count": 2}
	var result map[string]interface{}
	err = json.Unmarshal(rr.Body.Bytes(), &result)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	if result["ip"] != expected["ip"] || int(result["count"].(float64)) != expected["count"].(int) {
		t.Errorf("handler returned unexpected body: got %v want %v", result, expected)
	}
}
