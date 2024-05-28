package handlers

import (
	"github.com/franciscof12/rest-api-thn/internal/models"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestFeatureHandler(t *testing.T) {
	mutex.Lock()
	RequestInfo = []models.RequestInfo{}
	mutex.Unlock()

	req, err := http.NewRequest("GET", "/feature", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	http.HandlerFunc(FeatureHandler).ServeHTTP(rr, req)
	if rr.Code != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
	}

	expectedBody := "Hello THN!"
	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expectedBody)
	}

	mutex.Lock()
	defer mutex.Unlock()
	if len(RequestInfo) != 1 {
		t.Fatalf("expected 1 request info, got %d", len(RequestInfo))
	}

	ri := RequestInfo[0]
	verifyRequestDetails(t, ri, req)
}

func verifyRequestDetails(t *testing.T, ri models.RequestInfo, req *http.Request) {
	if ri.RemoteAddr != req.RemoteAddr {
		t.Errorf("expected RemoteAddr %v, got %v", req.RemoteAddr, ri.RemoteAddr)
	}
	if ri.RealIP != "" {
		t.Errorf("expected RealIP to be empty, got %v", ri.RealIP)
	}
	if _, err := time.Parse("2006-01-02 15:04:05", ri.Time); err != nil {
		t.Errorf("expected valid time, got %v", ri.Time)
	}
	if ri.Method != req.Method {
		t.Errorf("expected Method %v, got %v", req.Method, ri.Method)
	}
	if ri.Path != req.URL.Path {
		t.Errorf("expected Path %v, got %v", req.URL.Path, ri.Path)
	}
	if !compareHeaders(ri.Headers, req.Header) {
		t.Errorf("expected Headers to match, got %v", ri.Headers)
	}
}

func compareHeaders(h1, h2 map[string][]string) bool {
	if len(h1) != len(h2) {
		return false
	}

	for key, v1 := range h1 {
		v2, found := h2[key]
		if !found || len(v1) != len(v2) {
			return false
		}
		for i, value := range v1 {
			if value != v2[i] {
				return false
			}
		}
	}

	return true
}
