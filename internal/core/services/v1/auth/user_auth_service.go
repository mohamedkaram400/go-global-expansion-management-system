package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/mohamedkaram400/go-global-expansion-management-system/auth"
	"github.com/mohamedkaram400/go-global-expansion-management-system/config"
	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
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

	// Access token (short-lived, 15 min)
	accessToken, err := auth.GenerateAccessToken("user_id", user.ID, accessHours)
	if err != nil {
		return nil, "", "", errors.New("could not generate access token")
	}

	// Refresh token (long-lived, 7 days)
	refreshToken, err := auth.GenerateRefreshToken("user_id", user.ID, refreshDays) 
	if err != nil {
		return nil, "", "", errors.New("could not generate refresh token")
	}

	// Store refresh token in Redis or DB
	err = conn.RedisClient.Set(ctx, strconv.Itoa(int(user.ID)), refreshToken, 7*24*time.Hour).Err()
	if err != nil {
		return nil, "", "", errors.New("failed to store refresh token")
	}

	return user, accessToken, refreshToken, nil
}

func (svc *UserAuthService) Logout(userID uint) error {
	fmt.Println(userID)
	return conn.RedisClient.Del(context.Background(), strconv.FormatUint(uint64(userID), 10)).Err()
}