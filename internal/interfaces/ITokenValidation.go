package interfaces

import (
	"context"
	"ewallet-ums/helpers"
)

type ITokenValidationService interface {
	TokenValidation(ctx context.Context, token string) (*helpers.ClaimToken, error)
}
