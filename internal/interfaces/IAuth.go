package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
)

type IAuthService interface {
	ValidateSession(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

type IAuthMiddlewareHandler interface {
	MiddlewareValidateAuth(c *gin.Context)
}
