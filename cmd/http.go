package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/api"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/repository"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/services"
)

func ServeHTTP() {
	healthcheckSvc := &services.Healthcheck{}
	healthcheckAPI := &api.Healthcheck{
		HealthcheckServices: healthcheckSvc,
	}

	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerSvc := &services.RegisterService{
		RegisterRepo: registerRepo,
	}
	registerApi := &api.Register{
		RegisterService: registerSvc,
	}

	r := gin.Default()
	r.GET("/health", healthcheckAPI.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", registerApi.RegisterHandler)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}
