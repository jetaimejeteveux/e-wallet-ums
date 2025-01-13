package services

import (
	"context"
	"errors"
	"time"

	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type RefreshTokenService struct {
	UserRepo interfaces.IUserRepository
}

func (s *RefreshTokenService) RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error) {
	var (
		log = helpers.Logger
		now = time.Now()
	)

	token, err := helpers.GenerateToken(ctx, req.Username, req.Fullname, constants.Token, now)
	if err != nil {
		log.Error("failed to generate token = ", err)
		return nil, errors.New(constants.ErrFailedGenerateToken)
	}
	req.AccessToken = token

	err = s.UserRepo.UpdateTokenByRefreshToken(ctx, req)
	if err != nil {
		log.Error("failed updating token in db = ", err)
		return nil, errors.New(constants.ErrFailedUpdateToken)
	}

	return &models.RefreshTokenResponse{
		Token: token,
	}, nil
}

func (s *RefreshTokenService) ValidateRefreshToken(ctx context.Context, token string) (*helpers.ClaimToken, error) {
	userSession, err := s.UserRepo.GetUserSessionByRefreshToken(ctx, token)
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
