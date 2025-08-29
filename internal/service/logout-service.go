package service

import (
	"context"
	"errors"
	"ewallet-ums/internal/interfaces"
)

type LogoutService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LogoutService) Logout(ctx context.Context, token string) error {
	_, err := s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return errors.New("invalid session")
	}
	return s.UserRepo.DeleteUserSession(ctx, token)
}
