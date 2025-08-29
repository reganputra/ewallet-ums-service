package service

import (
	"context"
	"errors"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"time"
)

type TokenValidationService struct {
	UserRepo interfaces.IUserRepository
}

func (s *TokenValidationService) TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	var log = helpers.Logger

	claimToken, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		log.WithField("token_validation_error", err.Error()).Error("JWT token validation failed")
		return nil, errors.New("token validation failed: " + err.Error())
	}

	// Validate token type - only access tokens should be validated here
	if claimToken.TokenType != "access" {
		log.WithField("token_type", claimToken.TokenType).Error("Invalid token type for validation")
		return nil, errors.New("invalid token type: only access tokens can be validated")
	}

	session, err := s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		log.WithField("session_error", err.Error()).Error("Failed to retrieve user session")
		return nil, errors.New("session validation failed: invalid or expired session")
	}

	// Check if session has expired
	if session.TokenExpired.Before(time.Now()) {
		log.WithField("user_id", session.UserId).Error("Session has expired")
		return nil, errors.New("session has expired")
	}

	log.WithField("user_id", claimToken.UserID).Info("Token validation successful")
	return claimToken, nil
}
