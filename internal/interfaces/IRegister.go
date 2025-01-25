package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
)

type IRegisterService interface {
	Register(ctx context.Context, request models.User) (interface{}, error)
}
type IRegisterHandler interface {
	Register(*gin.Context)
}
