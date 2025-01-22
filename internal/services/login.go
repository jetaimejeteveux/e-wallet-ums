package services

import (
	"context"
	"errors"
	"time"

	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/models"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	UserRepo interfaces.IUserRepository
}

func (s *LoginService) Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error) {
	var (
		resp   models.LoginResponse
		logger = helpers.Logger
		now    = time.Now()
	)

	userDetail, err := s.UserRepo.GetUserByUsername(ctx, req.Username)
	if err != nil {
		logger.Error("error getting user from db = ", err)
		return nil, errors.New(constants.ErrFailedGetUser)
	}
	if userDetail == nil {
		logger.Error("error user not found")
		return nil, errors.New(constants.ErrUserNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userDetail.Password), []byte(req.Password))
	if err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(ctx, userDetail.Username, userDetail.FullName, constants.Token, now)
	if err != nil {
		logger.Error("failed to generate token: ", err)
		return nil, errors.New(constants.ErrFailedGenerateToken)
	}

	refreshToken, err := helpers.GenerateToken(ctx, userDetail.Username, userDetail.FullName, constants.RefreshToken, now)
	if err != nil {
		logger.Error("failed to generate token: ", err)
		return nil, errors.New(constants.ErrFailedGenerateRefreshToken)
	}

	userSession := &models.UserSession{
		UserID:              userDetail.ID,
		Token:               token,
		RefreshToken:        refreshToken,
		TokenExpired:        now.Add(helpers.MapTypeToken[constants.Token]),
		RefreshTokenExpired: now.Add(helpers.MapTypeToken[constants.RefreshToken]),
	}
	err = s.UserRepo.InsertUserSession(ctx, userSession)
	if err != nil {
		logger.Error("failed to insert session token: ", err)
		return nil, errors.New(constants.ErrFailedInsertSession)
	}

	resp.UserID = userDetail.ID
	resp.Username = userDetail.Username
	resp.Fullname = userDetail.FullName
	resp.Email = userDetail.Email
	resp.Token = token
	resp.RefreshToken = refreshToken

	return &resp, nil

}
