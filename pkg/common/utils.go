package common

import (
	"github.com/joho/godotenv"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strings"
)

// FromRequest return client's real public IP address from http request headers.
func FromRequest(r *http.Request) string {
	xRealIP := r.Header.Get("X-Real-Ip")

	if xRealIP == "" {
		var remoteIP string
		if strings.ContainsRune(r.RemoteAddr, ':') {
			remoteIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		} else {
			remoteIP = r.RemoteAddr
		}

		return remoteIP
	}

	return xRealIP
}

func FormatHeaders(headers http.Header) string {
	var formattedHeaders string
	for key, values := range headers {
		formattedHeaders += key + ": " + values[0] + "; "
	}
	return formattedHeaders
}

func LoadEnvironment() {
	err := godotenv.Load()
	if err != nil {
		slog.Error("Error loading .env file", "error", err.Error())
		os.Exit(1) // Terminate if env cannot be loaded
	}
}
