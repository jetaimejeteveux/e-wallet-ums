package services

import (
	"context"
	"errors"

	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
)

type TokenValidationService struct {
	UserRepository interfaces.IUserRepository
}

func (s *TokenValidationService) ValidateToken(ctx context.Context, token string) (*helpers.ClaimToken, error) {

	claimToken, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		return claimToken, errors.New(constants.ErrFailedValidateToken)
	}

	_, err = s.UserRepository.GetUserSessionByToken(ctx, token)
	if err != nil {
		return claimToken, errors.New(constants.ErrFailedGetUserSession)
	}

	return claimToken, nil
}
