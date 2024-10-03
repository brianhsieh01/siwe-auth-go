package auth

import (
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	db *pg.DB
}

func NewAuthRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}
