package api

import (
	"context"
	"ewallet-ums/cmd/proto/tokenValidation"
	"ewallet-ums/constant"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"fmt"
)

type TokenValidationHandler struct {
	TokenValidationSvc interfaces.ITokenValidationService
	tokenValidation.UnimplementedTokenValidationServer
}

func (s *TokenValidationHandler) ValidateToken(ctx context.Context, req *tokenValidation.TokenRequest) (*tokenValidation.TokenResponse, error) {

	var log = helpers.Logger

	token := req.GetToken()
	if token == "" {
		err := fmt.Errorf("token is required")
		log.WithField("request_id", ctx.Value("request_id")).Error("Token validation failed: empty token")
		return &tokenValidation.TokenResponse{
			Message: "Token is required",
		}, err
	}

	log.WithField("token_length", len(token)).Info("Processing token validation request")

	claimToken, err := s.TokenValidationSvc.TokenValidation(ctx, token)
	if err != nil {
		log.WithField("validation_error", err.Error()).Error("Token validation failed in service layer")
		return &tokenValidation.TokenResponse{
			Message: err.Error(),
		}, err
	}

	log.WithField("user_id", claimToken.UserID).WithField("username", claimToken.Username).Info("Token validation successful")

	return &tokenValidation.TokenResponse{
		Message: constant.Success,
		Data: &tokenValidation.UserData{
			UserId:   int64(claimToken.UserID),
			Username: claimToken.Username,
			FullName: claimToken.FullName,
		},
	}, nil
}
