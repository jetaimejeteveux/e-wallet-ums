package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
)

type AuthHandler struct {
	AuthService interfaces.IAuthService
}

func (h *AuthHandler) MiddlewareValidateAuth(c *gin.Context) {
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
