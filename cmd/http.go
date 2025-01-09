package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jetaimejeteveux/e-wallet-ums/helpers"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/api"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/interfaces"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/repository"
	"github.com/jetaimejeteveux/e-wallet-ums/internal/services"
)

func ServeHTTP() {
	dependency := dependencyInject()

	r := gin.Default()
	// r.GET("/health", dependency.HealthcheckHandlerHTTP)

	userV1 := r.Group("/user/v1")
	userV1.POST("/register", dependency.RegisterAPI.Register)
	userV1.POST("/login", dependency.LoginAPI.Login)

	err := r.Run(":" + helpers.GetEnv("PORT", ""))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository
	RegisterAPI    interfaces.IRegisterHandler
	LoginAPI       interfaces.ILoginHandler
}

func dependencyInject() Dependency {
	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &services.RegisterService{
		UserRepo: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &services.LoginService{
		UserRepo: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
	}
}
