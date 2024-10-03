package main

import (
	"log"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/config"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/internal/auth"
	"github.com/Larryx-s-Kitchen/siwe-auth-go/internal/database"
	"github.com/labstack/echo/v4"
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

	e := echo.New()

	repo := auth.NewAuthRepository(db)
	service := auth.NewAuthService(repo, cfg)
	handler := auth.NewAuthHandler(service)

	e.GET("/auth/nonce", handler.GetNonce)
	e.POST("/auth/signin", handler.SignIn)

	e.Logger.Fatal(e.Start(":8080"))

}
