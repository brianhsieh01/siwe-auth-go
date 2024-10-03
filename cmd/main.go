package main

import (
	"log"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/config"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/internal/database"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDatabaseConnection(&cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()
}
