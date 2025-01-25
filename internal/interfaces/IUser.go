package interfaces

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InsertUserSession(ctx context.Context, sessions *models.UserSession) error
	GetUserSessionByToken(ctx context.Context, token string) (*models.UserSession, error)
	UpdateTokenByRefreshToken(ctx context.Context, req *models.RefreshTokenRequest) error
	GetUserSessionByRefreshToken(ctx context.Context, refreshToken string) (*models.UserSession, error)
}
