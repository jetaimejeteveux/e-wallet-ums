package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
)

type ILogoutService interface {
	Logout(ctx context.Context, req *models.UserSession) error
}

type ILogoutHandler interface {
	Logout(c *gin.Context)
}
