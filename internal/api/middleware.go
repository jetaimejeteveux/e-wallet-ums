package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
)

type MiddlewareHandler struct {
	AuthService         interfaces.IAuthService
	RefreshTokenService interfaces.IRefreshTokenService
}

func (h *MiddlewareHandler) MiddlewareValidateAuth(c *gin.Context) {
	ctx := c.Request.Context()
	tokenAuth := c.Request.Header.Get("Authorization")
	if tokenAuth == "" {
		log.Println("empty authorization")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}

	claim, err := h.AuthService.ValidateSession(ctx, tokenAuth)
	if err != nil {
		log.Println("empty authorization")
		helpers.SendResponseHTTP(c, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		c.Abort()
		return
	}

	c.Set(constants.Token, claim)

	c.Next()
}

func (h *MiddlewareHandler) MiddlewareRefreshToken(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		ctx.Abort()
		return
	}

	claims, err := h.RefreshTokenService.ValidateRefreshToken(ctx, auth)
	if err != nil {
		log.Println("token validation failed:", err)
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnauthorized, nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claims)
	ctx.Next()
}
