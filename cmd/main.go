package main

import (
	"go_chi/cmd/server"
	"go_chi/internal/db"
)

func main() {
	db := db.InitDatabase()
	defer db.Close()

	// Start server
	server.Start(db)
}
