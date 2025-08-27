package services

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports"
)

type AuthService struct {
	repo ports.AuthRepository 
}

func NewAuthService(repo ports.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Register(ctx context.Context, client *entities.Client) (*entities.Client, error) {
	if client.CompanyName == "" || client.ContactEmail == "" {
		return nil, errors.New("company name and email are required")
	}

	return s.repo.Register(ctx, client)
}


func (s *AuthService) Login(ctx context.Context, email string, password string) (*entities.Client, error) {
	if email == "" {
		return nil, errors.New("email required")
	}
	return s.repo.Login(ctx, email, password)
}

func (s *AuthService) Logout(clientID string) (string, error) {
	return s.repo.Logout(clientID)
}