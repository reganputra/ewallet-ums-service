package service

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RegisterService) Register(ctx context.Context, req models.User) (interface{}, error) {

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashPwd)

	err = s.UserRepo.InsertNewUser(ctx, &req)
	if err != nil {
		return nil, err
	}

	resp := req
	resp.Password = ""
	return resp, nil
}
