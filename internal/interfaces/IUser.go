package interfaces

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type IUserRepository interface {
	InsertUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
	InsertUserSession(ctx context.Context, sessions *models.UserSession) error
}
