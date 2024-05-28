package tests

import (
	"encoding/json"
	"github.com/franciscof12/rest-api-thn/internal/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFeatureAndMetricsIntegration(t *testing.T) {
	ipAddresses := []string{"192.168.1.1", "192.168.1.1", "10.0.0.1"}
	for _, ip := range ipAddresses {
		reqFeature, err := http.NewRequest("GET", "/feature", nil)
		if err != nil {
			t.Fatal("Error creating request for /feature:", err)
		}
		reqFeature.RemoteAddr = ip + ":12345"
		rrFeature := httptest.NewRecorder()
		featureHandler := http.HandlerFunc(handlers.FeatureHandler)
		featureHandler.ServeHTTP(rrFeature, reqFeature)
		if rrFeature.Code != http.StatusOK {
			t.Errorf("FeatureHandler returned wrong status code: got %v want %v", rrFeature.Code, http.StatusOK)
		}
	}

	testIP := "192.168.1.1"
	reqMetrics, err := http.NewRequest("GET", "/metrics?ip="+testIP, nil)
	if err != nil {
		t.Fatal("Error creating request for /metrics:", err)
	}
	rrMetrics := httptest.NewRecorder()
	metricsHandler := http.HandlerFunc(handlers.MetricsHandler)
	metricsHandler.ServeHTTP(rrMetrics, reqMetrics)

	if rrMetrics.Code != http.StatusOK {
		t.Errorf("MetricsHandler returned wrong status code: got %v want %v", rrMetrics.Code, http.StatusOK)
	}

	var actualResponse map[string]interface{}
	err = json.Unmarshal(rrMetrics.Body.Bytes(), &actualResponse)
	if err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedResponse := map[string]interface{}{"count": 2, "ip": "192.168.1.1"}
	if actualResponse["ip"] != expectedResponse["ip"] || int(actualResponse["count"].(float64)) != expectedResponse["count"].(int) {
		t.Errorf("MetricsHandler returned unexpected body: got %v want %v", actualResponse, expectedResponse)
	}
}
