package services

import (
	"context"
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/mohamedkaram400/go-global-expansion-management-system/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/core/entities"
	"github.com/mohamedkaram400/go-global-expansion-management-system/internal/ports"
	"github.com/mohamedkaram400/go-global-expansion-management-system/pkg"
	"github.com/mohamedkaram400/go-global-expansion-management-system/requests"
)

type AuthService struct {
	repo ports.AuthRepository 
}

func NewAuthService(repo ports.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (svc *AuthService) Register(ctx context.Context, req *requests.RegisterRequest) (*entities.Client, error) {
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
		CompanyName:    req.CompanyName,
		ContactEmail:   req.ContactEmail,
		Password:   	hashedPwd,
	}

	// Save to DB via repo
	return svc.repo.Register(ctx, client)
}

func (svc *AuthService) Login(ctx context.Context, req *requests.LoginRequest) (*entities.Client, string, string, error) {
	accessHours, err := strconv.Atoi(os.Getenv("ACCESS_TOKEN_TIME"))
	if err != nil {
		return nil, "", "", errors.New("invalid ACCESS_TOKEN_TIME in env")
	}

	refreshDays, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_TIME"))
	if err != nil {
		return nil, "", "", errors.New("invalid REFRESH_TOKEN_TIME in env")
	}

	// Get company name exists
	client, err := svc.repo.GetClientByCompanyName(ctx, req.CompanyName)
	if err != nil || client == nil {
		return nil, "", "", errors.New("company not found")
	}

	if err := pkg.CheckPassword(client.Password, req.Password); err != nil {
		return nil, "", "", errors.New("invalid password")
	}

	// Access token (short-lived, 15 min)
	accessToken, err := auth.GenerateAccessToken(client.ID, client.CompanyName, accessHours)
	if err != nil {
		return nil, "", "", errors.New("could not generate access token")
	}

	// Refresh token (long-lived, 7 days)
	refreshToken, err := auth.GenerateRefreshToken(client.ID, client.CompanyName, refreshDays) 
	if err != nil {
		return nil, "", "", errors.New("could not generate refresh token")
	}

	// Store refresh token in Redis or DB
	err = conn.RedisClient.Set(ctx, strconv.Itoa(int(client.ID)), refreshToken, 7*24*time.Hour).Err()
	if err != nil {
		return nil, "", "", errors.New("failed to store refresh token")
	}

	return client, accessToken, refreshToken, nil
}

func (svc *AuthService) Logout(clientID uint) error {
	return conn.RedisClient.Del(context.Background(), strconv.FormatUint(uint64(clientID), 10)).Err()
}