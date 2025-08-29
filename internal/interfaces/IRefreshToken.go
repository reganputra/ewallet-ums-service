package interfaces

import (
	"context"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/models"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, refreshToken string, tokenClaim helpers.ClaimToken) (models.RefreshTokenResponse, error)
}
