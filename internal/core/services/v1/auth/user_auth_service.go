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

type UserAuthService struct {
	repo ports.UserAuthRepository 
}

func NewUserAuthService(repo ports.UserAuthRepository) *UserAuthService {
	return &UserAuthService{repo: repo}
}

func (svc *UserAuthService) Login(ctx context.Context, req *requests.UserLoginRequest) (*entities.User, string, string, error) {
	config := config.LoadConfig()

	accessHours := config.AccessTokenTime
	refreshDays := config.RefrashTokenTime

	// Get company name exists
	user, err := svc.repo.GetUserByEmail(ctx, req.Email)
	if err != nil || user == nil {
		return nil, "", "", errors.New("email not found")
	}

	if err := pkg.CheckPassword(user.Password, req.Password); err != nil {
		return nil, "", "", errors.New("invalid password")
	}

	accessToken, refreshToken, err := pkg.IssueTokens("user_id", user.ID, accessHours, refreshDays)
	if err != nil {
		return nil, "", "", errors.New("failed to generate access and refresh token")
	}

	// Store refresh token in Redis or DB
	if err := pkg.StoreRefreshToken(ctx, user.ID, refreshToken, refreshDays); err != nil {
		return nil, "", "", errors.New("failed to store refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (svc *UserAuthService) Logout(userID uint) error {
	err := pkg.DeleteRefreshToken(userID)
	if err != nil {
		return errors.New("failed to store refresh token")
	}
	return errors.New("")
}