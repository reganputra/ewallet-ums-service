package service

import (
	"context"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/models"

	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	RegisterRepo interfaces.IRegisterRepository
}

func (s *RegisterService) Register(ctx context.Context, req models.User) (interface{}, error) {

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = s.RegisterRepo.InsertNewUser(ctx, &req)
	if err != nil {
		return nil, err
	}

	req.Password = string(hashPwd)
	resp := req
	resp.Password = ""
	return resp, nil
}
