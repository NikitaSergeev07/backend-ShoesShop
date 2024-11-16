package service

import (
	"FamilyEmo"
	"FamilyEmo/pkg/repository"
)

type Authorization interface {
	CreateUser(user FamilyEmo.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
