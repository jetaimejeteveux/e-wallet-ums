package cmd

import (
	"ewallet-framework-1/helpers"
	"ewallet-framework-1/internal/api"
	"ewallet-framework-1/internal/services"
	"github.com/gin-gonic/gin"
	"log"
)

func ServeHTTP() {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	r := gin.Default()

	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
