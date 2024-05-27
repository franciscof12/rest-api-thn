package models

type RequestInfo struct {
	RemoteAddr string
	RealIP     string
	Time       string
	Method     string
	Path       string
	Headers    map[string][]string
}
