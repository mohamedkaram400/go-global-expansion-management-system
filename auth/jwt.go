package auth

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("YOUR_SECRET_KEY"))

type TokenType string

const (
	AccessToken  TokenType = "access"
	RefreshToken TokenType = "refresh"
)

func GenerateToken(clientID uint, CompanyName string, duration time.Duration, tokenType TokenType) (string, error) {

	claims := jwt.MapClaims{
		"client_id"		: clientID,
		"company_name"	: CompanyName,
		"exp"			: time.Now().Add(duration).Unix(),
		"type"			: tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func GenerateAccessToken(clientID uint, CompanyName string, hours int) (string, error) {
	return GenerateToken(clientID, CompanyName, time.Duration(hours)*time.Hour, AccessToken)
}

func GenerateRefreshToken(clientID uint, CompanyName string, days int) (string, error) {
	return GenerateToken(clientID, CompanyName, time.Duration(days)*24*time.Hour, RefreshToken)
}

func ValidateJWT(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid claims")
	}
	
	clientIDFloat, ok := claims["client_id"].(float64)
	if !ok {
		return 0, errors.New("client id not found in token")
	}
	
	return uint(clientIDFloat), nil
}