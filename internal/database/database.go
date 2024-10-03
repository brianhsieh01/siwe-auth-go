package database

import (
	"fmt"

	"github.com/Larryx-s-Kitchen/siwe-auth-go/config"
	"github.com/go-pg/pg/v10"
)

func InitDatabaseConnection(cfg *config.DatabaseConfig) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		User:     cfg.User,
		Password: cfg.Password,
		Database: cfg.Name,
	})

	_, err := db.Exec("SELECT 1")
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database: %w", err)
	}

	return db, nil
}
