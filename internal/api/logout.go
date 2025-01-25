package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/models"
)

type LogoutHandler struct {
	LogoutSvc interfaces.ILogoutService
}

func (h *LogoutHandler) Logout(c *gin.Context) {
	var (
		log = helpers.Logger
		ctx = c.Request.Context()
	)

	token := c.Request.Header.Get("Authorization")

	if err := h.LogoutSvc.Logout(ctx, &models.UserSession{
		Token: token,
	}); err != nil {
		log.Error("error in service : ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedLogout, nil)
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, nil)

}
