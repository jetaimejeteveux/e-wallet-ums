package services

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	UserRepo interfaces.IUserRepository
}

func (r *RegisterService) Register(ctx context.Context, request models.User) (interface{}, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	request.Password = string(hashedPassword)

	err = r.UserRepo.InsertUser(ctx, &request)
	if err != nil {
		return nil, err
	}

	resp := request
	resp.Password = ""
	return resp, nil

}
