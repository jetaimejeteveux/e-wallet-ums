package services

import (
	"context"
	"errors"
	"time"

	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
)

type AuthService struct {
	UserRepo interfaces.IUserRepository
}

func (s *AuthService) ValidateSession(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	userSession, err := s.UserRepo.GetUserSessionByToken(ctx, token)
	if err != nil {
		return nil, err
	}
	if userSession == nil {
		return nil, errors.New(constants.ErrSessionNotFound)
	}

	claim, err := helpers.ValidateToken(ctx, token)
	if err != nil {
		return nil, errors.New(constants.ErrUnauthorized)
	}
	if time.Now().Unix() > claim.ExpiresAt.Unix() {
		return nil, errors.New(constants.ErrUnauthorized)
	}

	return claim, nil
}
