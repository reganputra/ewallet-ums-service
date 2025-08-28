package service

import (
	"context"
	"errors"
	"ewallet-ums/helpers"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (models.LoginResponse, error) {

	var resp models.LoginResponse
	now := time.Now()

	userDetail, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		return resp, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		return resp, errors.New("failed compare password")
	}

	token, err := helpers.GenerateToken(ctx, userDetail.Id, userDetail.Username, userDetail.FullName, "access", now)
	if err != nil {
		return resp, errors.New("failed generate token")
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.Id, userDetail.Username, userDetail.FullName, "refresh", now)
	if err != nil {
		return resp, errors.New("failed generate refresh token")
	}

	userSession := &models.UserSession{
		UserId:              uint(userDetail.Id),
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTokenTypes["access"]),
		RefreshTokenExpired: now.Add(helpers.MapTokenTypes["refresh"]),
	}

	err = s.UserRepo.InsertNewUserSession(ctx, userSession)
	if err != nil {
		return resp, errors.New("failed insert user session")
	}

	resp.UserID = userDetail.Id
	resp.Username = userDetail.Username
	resp.FullName = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken

	return resp, nil
}
