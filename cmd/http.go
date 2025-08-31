package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/interfaces"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	healthCheckRepo := repository.NewHealthCheckRepo()
	healthCheckSvc := &service.HealthCheck{
		HealthCheckRepository: healthCheckRepo,
	}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	// Initialize all dependencies
	deps := InitializeDependencies()

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandler)

	userV1 := r.Group("/users/v1")
	userV1.POST("/register", deps.RegisterAPI.Register)
	userV1.POST("/login", deps.LoginAPI.Login)

	userV1.DELETE("/logout", deps.MiddlewareValidateAuth, deps.LogoutAPI.Logout)
	userV1.PUT("/refresh-token", deps.MiddlewareRefreshToken, deps.RefreshTokenAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

type Dependency struct {
	// Repositories
	UserRepository interfaces.IUserRepository

	// Api Handler
	RegisterAPI     interfaces.IRegisterHandler
	LoginAPI        interfaces.ILoginHandler
	LogoutAPI       interfaces.ILogoutHandler
	RefreshTokenAPI interfaces.IRefreshTokenHandler

	TokenValidationAPI api.TokenValidationHandler
}

func InitializeDependencies() Dependency {

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	registerSvc := &service.RegisterService{
		UserRepo: userRepo,
	}
	registerApi := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	loginSvc := &service.LoginService{
		UserRepo: userRepo,
	}
	loginApi := &api.LoginHandler{
		LoginService: loginSvc,
	}

	logoutSvc := &service.LogoutService{
		UserRepo: userRepo,
	}
	logoutApi := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	refreshTokenSvc := &service.RefreshTokenService{
		UserRepo: userRepo,
	}
	refreshTokenApi := &api.RefreshTokenHandler{
		RefreshTokenSvc: refreshTokenSvc,
	}

	tokenValidationSvc := &service.TokenValidationService{
		UserRepo: userRepo,
	}
	tokenValidationApi := api.TokenValidationHandler{
		TokenValidationSvc: tokenValidationSvc,
	}

	return Dependency{
		UserRepository:     userRepo,
		RegisterAPI:        registerApi,
		LoginAPI:           loginApi,
		LogoutAPI:          logoutApi,
		RefreshTokenAPI:    refreshTokenApi,
		TokenValidationAPI: tokenValidationApi,
	}
}
