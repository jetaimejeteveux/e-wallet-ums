package repository

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) InsertUser(ctx context.Context, user *models.User) error {
	return r.DB.Create(&user).Error
}

func (r *UserRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) InsertUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Create(&session).Error
}

func (r *UserRepository) DeleteUserSession(ctx context.Context, session *models.UserSession) error {
	return r.DB.Where("token = ?", session.Token).Delete(&session).Error
}

func (r *UserRepository) GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error) {
	var resp *models.UserSession
	err := r.DB.Where("token = ?", token).First(&resp).Error
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (r *UserRepository) UpdateTokenByRefreshToken(ctx context.Context, req *models.RefreshTokenRequest) error {
	return r.DB.Model(&models.UserSession{}).Where("refresh_token = ?", req.RefreshToken).Update("token", req.AccessToken).Error
}

func (r *UserRepository) GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error) {
	var resp *models.UserSession
	err := r.DB.Where("refresh_token = ?", refreshToken).First(&resp).Error
	if err != nil {
		return nil, err
	}
	return resp, nil
}
