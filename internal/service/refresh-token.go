package service

import (
	"context"
	"errors"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error) {

	resp := models.RefreshTokenResponse{}
	token, err := helpers.GenerateToken(ctx, tokenClaim.UserID, tokenClaim.Username, tokenClaim.Email, tokenClaim.FullName, "access", time.Now())
	if err != nil {
		return resp, errors.New("failed to generate new token")
	}

	err = s.UserRepo.UpdateTokenWByRefreshToken(ctx, token, refreshToken)
	if err != nil {
		return resp, errors.New("failed to update new token")
	}
	resp.Token = token
	return resp, nil

}
