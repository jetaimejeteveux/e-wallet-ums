package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/constants"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/models"
)

type Register struct {
	RegisterService interfaces.IRegisterService
}

func (api *Register) RegisterHandler(c *gin.Context) {
	ctx := c.Request.Context()

	var (
		log = helpers.Logger
	)
	req := models.User{}

	if err := c.ShouldBindJSON(&req); err != nil {
		log.Info("failed to parse request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
		return
	}

	if err := req.Validate(); err != nil {
		log.Info("failed to validate request: ", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedValidateRequest, nil)
		return
	}

	resp, err := api.RegisterService.Register(ctx, req)
	if err != nil {
		log.Infof("failed to register new user: %v", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrFailedRegisterUser, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, resp)
}
