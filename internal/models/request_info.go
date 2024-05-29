package models

import "go.mongodb.org/mongo-driver/mongo"

var ResquestCollection *mongo.Collection

func SetRequestCollection(c *mongo.Collection) {
	ResquestCollection = c
}

type RequestInfo struct {
	RemoteAddr string              `json:"remote_addr"`
	RealIP     string              `json:"realip"`
	Time       string              `json:"time"`
	Method     string              `json:"method"`
	Path       string              `json:"path"`
	Headers    map[string][]string `json:"headers"`
}
