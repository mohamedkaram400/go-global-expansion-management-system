package services

import (
	"context"
	"errors"

	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities/v1"
	ports "github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports/v1/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/pkg"
	requests "github.com/mohamedkaram400/go-global-expansion-management-system/requests/v1/auth"
)

type ClientAuthService struct {
	repo ports.ClientAuthRepository
}

func NewClientAuthService(repo ports.ClientAuthRepository) *ClientAuthService {
	return &ClientAuthService{repo: repo}
}

func (svc *ClientAuthService) Register(ctx context.Context, req *requests.ClientRegisterRequest) (*entities.Client, error) {
	// Check if company name exists
	existing, _ := svc.repo.GetClientByCompanyName(ctx, req.CompanyName)
	if existing != nil {
		return nil, errors.New("company name already exists")
	}

	// Hash password
	hashedPwd, err := pkg.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create client model
	client := &entities.Client{
		CompanyName:  req.CompanyName,
		ContactEmail: req.ContactEmail,
		Password:     hashedPwd,
	}

	// Save to DB via repo
	return svc.repo.Register(ctx, client)
}

func (svc *ClientAuthService) Login(ctx context.Context, req *requests.ClientLoginRequest) (*entities.Client, string, string, error) {
	config := config.LoadConfig()

	accessHours := config.AccessTokenTime
	refreshDays := config.RefrashTokenTime

	// Get company name exists
	client, err := svc.repo.GetClientByCompanyName(ctx, req.CompanyName)
	if err != nil || client == nil {
		return nil, "", "", errors.New("company not found")
	}

	if err := pkg.CheckPassword(client.Password, req.Password); err != nil {
		return nil, "", "", errors.New("invalid password")
	}

	accessToken, refreshToken, err := pkg.IssueTokens("client_id", client.ID, accessHours, refreshDays)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access and refresh token")
	}

	// Store refresh token in Redis or DB
	if err := pkg.StoreRefreshToken(ctx, client.ID, refreshToken, refreshDays); err != nil {
		return nil, "", "", errors.New("failed to store refresh token")
	}

	return client, accessToken, refreshToken, nil
}

func (svc *ClientAuthService) Logout(clientID string) error {
	err := pkg.DeleteRefreshToken(clientID)
	if err != nil {
		return errors.New("failed to store refresh token")
	}
	return errors.New("")
}
