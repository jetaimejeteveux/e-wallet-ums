package api

import (
	"ewallet-framework-1/helpers"
	"ewallet-framework-1/internal/interfaces"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Healthcheck struct {
	HealthcheckServices interfaces.IHealthcheckServices
}

func (api *Healthcheck) HealthcheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthcheckServices.HealthcheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}
	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
