package main

import (
	"github.com/franciscof12/rest-api-thn/internal/db"
	"github.com/franciscof12/rest-api-thn/internal/server"
	"github.com/franciscof12/rest-api-thn/pkg/common"
)

func main() {
	common.LoadEnvironment()
	client := db.SetupDatabase()
	db.SetupModels(client)
	server.Start()
}
