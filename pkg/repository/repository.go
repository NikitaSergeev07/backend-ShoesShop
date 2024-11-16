package repository

import (
	"FamilyEmo"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user FamilyEmo.User) (int, error)
	GetUser(email, password string) (FamilyEmo.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
