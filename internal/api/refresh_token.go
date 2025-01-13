package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (h *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log = helpers.Logger
		ctx = c.Request.Context()
	)
	refreshToken := c.Request.Header.Get("Authorization")
	claim, ok := c.Get("token")
	if !ok {
		log.Error("failed to get token from context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedGetToken, nil)
		c.Abort()
		return
	}

	tokenClaim, ok := claim.(*helpers.ClaimToken)
	if !ok {
		log.Error("failed to parse token")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedGetToken, nil)
		c.Abort()
		return
	}

	req := &models.RefreshTokenRequest{
		Username:     tokenClaim.Username,
		Fullname:     tokenClaim.Fullname,
		RefreshToken: refreshToken,
	}
	resp, err := h.RefreshTokenService.RefreshToken(ctx, req)
	if err != nil {
		log.Error("error in service = ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedRefreshToken, nil)
		c.Abort()
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
