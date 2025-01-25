package interfaces

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
)

type IRefreshTokenService interface {
	RefreshToken(ctx context.Context, req *models.RefreshTokenRequest) (*models.RefreshTokenResponse, error)
	ValidateRefreshToken(ctx context.Context, token string) (*helpers.ClaimToken, error)
}

type IRefreshTokenHandler interface {
	RefreshToken(*gin.Context)
}
