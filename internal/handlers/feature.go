package handlers

import (
	"github.com/franciscof12/rest-api-thn/internal/models"
	"github.com/franciscof12/rest-api-thn/pkg/common"
	"log/slog"
	"net/http"
	"sync"
	"time"
)

var (
	RequestInfo []models.RequestInfo
	mutex       sync.Mutex
)

func FeatureHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("Hello THN!"))
	if err != nil {
		slog.Error("Error writing response:", "error", err.Error())
	}

	request := createRequestInfo(r)
	mutex.Lock()
	defer mutex.Unlock()
	RequestInfo = append(RequestInfo, request)

	slog.Info("Request saved",
		"RemoteAddress", request.RemoteAddr,
		"RealIP", request.RealIP,
		"Time", request.Time,
		"Method", request.Method,
		"Path", request.Path,
		"Headers", common.FormatHeaders(request.Headers),
	)
}

func createRequestInfo(r *http.Request) models.RequestInfo {
	realIP := common.FromRequest(r)
	remoteAddr := r.RemoteAddr
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	method := r.Method
	path := r.URL.Path
	header := r.Header

	return models.RequestInfo{
		RemoteAddr: remoteAddr,
		RealIP:     realIP,
		Time:       currentTime,
		Method:     method,
		Path:       path,
		Headers:    header,
	}
}
