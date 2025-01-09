package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type ILoginService interface {
	Login(ctx context.Context, req models.LoginRequest) (*models.LoginResponse, error)
}

type ILoginHandler interface {
	Login(c *gin.Context)
}
