package auth

import (
	"time"
	"errors"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("YOUR_SECRET_KEY") // load from env

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GenerateToken(CompanyName string, duration time.Duration, tokenType TokenType) (string, error) {

	claims := jwt.MapClaims{
		"company_name": CompanyName,
		"exp":         time.Now().Add(duration).Unix(),
		"type":        tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(CompanyName string, hours int) (string, error) {
	return GenerateToken(CompanyName, time.Duration(hours)*time.Hour, AccessToken)
}

func GenerateRefreshToken(CompanyName string, days int) (string, error) {
	return GenerateToken(CompanyName, time.Duration(days)*24*time.Hour, RefreshToken)
}

func ValidateJWT(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return "", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid claims")
	}

	CompanyName, ok := claims["client_id"].(string)
	if !ok {
		return "", errors.New("client_id not found in token")
	}

	return CompanyName, nil
}