package pkg

import (
	"context"
	"strconv"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/mohamedkaram400/go-global-expansion-management-system/auth"

	"github.com/mohamedkaram400/go-global-expansion-management-system/conn"
)

func CheckPassword(hashed, plain string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func StoreRefreshToken(ctx context.Context, id uint, token string, days int) error {
	return conn.RedisClient.Set(ctx,
		strconv.Itoa(int(id)),
		token,
		time.Duration(days)*24*time.Hour,
	).Err()
}

func DeleteRefreshToken(id uint) error {
	return conn.RedisClient.Del(context.Background(),
		strconv.FormatUint(uint64(id), 10),
	).Err()
}

func IssueTokens(subjectKey string, subjectID uint, accessHours, refreshDays int) (string, string, error) {

	// Access token (short-lived, 15 min)
	accessToken, err := auth.GenerateAccessToken(subjectKey, subjectID, accessHours)
	if err != nil {
		return "", "", err
	}

	// Refresh token (long-lived, 7 days)
	refreshToken, err := auth.GenerateRefreshToken(subjectKey, subjectID, refreshDays)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}