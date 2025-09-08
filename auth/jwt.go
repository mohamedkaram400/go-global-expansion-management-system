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

func GenerateToken(subjectKey string, subjectID uint, duration time.Duration, tokenType TokenType) (string, error) {

	claims := jwt.MapClaims{
		subjectKey		: subjectID,
		"exp"			: time.Now().Add(duration).Unix(),
		"type"			: tokenType,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
} 

func GenerateAccessToken(subjectKey string, subjectID uint, hours int) (string, error) {
	return GenerateToken(subjectKey, subjectID,  time.Duration(hours)*time.Hour, AccessToken)
}

func GenerateRefreshToken(subjectKey string, subjectID uint, days int) (string, error) {
	return GenerateToken(subjectKey, subjectID,  time.Duration(days)*24*time.Hour, RefreshToken)
}

func ValidateJWT(tokenString string, key string) (uint, error) {
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
	
	idFloat, ok := claims[key].(float64)
	if !ok {
		return 0, errors.New("id not found in token")
	}
	
	return uint(idFloat), nil
}

// how i can make new middleware check if request come from admin user allow it other wise rejecte 