package services

import (
	"context"

	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/repository"
)

type LogoutServices struct {
	UserRepo repository.UserRepository
}

func (s *LogoutServices) Logout(ctx context.Context, req *models.UserSession) error {
	return s.UserRepo.DeleteUserSession(ctx, &models.UserSession{
		Token: req.Token,
	})
}
