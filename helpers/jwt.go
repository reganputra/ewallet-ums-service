package helpers

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ClaimToken struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

var MapTokenTypes = map[string]time.Duration{
	"access":  time.Minute * 15,
	"refresh": time.Hour * 24,
}

func GenerateToken(ctx context.Context, userId int, username, email, fullName, tokenType string, now time.Time) (string, error) {

	secret := []byte(GetEnv("APP_SECRET", ""))

	claims := ClaimToken{
		UserID:   userId,
		Username: username,
		FullName: fullName,
		Email:    email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    GetEnv("APP_NAME", ""),
			IssuedAt:  jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(MapTokenTypes[tokenType])),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return tokenString, errors.New("failed to generate token")
	}
	return tokenString, nil
}

func ValidateToken(ctx context.Context, token string) (*ClaimToken, error) {

	secret := []byte(GetEnv("APP_SECRET", ""))

	parsedToken, err := jwt.ParseWithClaims(token, &ClaimToken{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := parsedToken.Claims.(*ClaimToken); ok && parsedToken.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token claims")
}
