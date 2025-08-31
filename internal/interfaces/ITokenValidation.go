package interfaces

import (
	"context"
	"ewallet-ums/cmd/proto/tokenValidation"
	"ewallet-ums/helpers"
)

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

type ITokenValidationHandler interface {
	ValidateToken(ctx context.Context, req *tokenValidation.TokenRequest) (*tokenValidation.TokenResponse, error)
}
