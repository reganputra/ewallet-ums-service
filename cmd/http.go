package cmd

import (
	"ewallet-ums/helpers"
	"ewallet-ums/internal/api"
	"ewallet-ums/internal/repository"
	"ewallet-ums/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func ServerHttp() {

	// Initialize all dependencies
	deps := InitializeDependencies()

	r := gin.Default()
	r.GET("/health", deps.HealthCheckAPI.HealthCheckHandler)

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
	UserRepo        repository.UserRepository
	HealthCheckRepo repository.HealthCheckRepo

	// Services
	LoginService        service.LoginService
	RegisterService     service.RegisterService
	HealthCheckSvc      service.HealthCheck
	LogoutService       service.LogoutService
	RefreshTokenService service.RefreshTokenService

	// API Handlers
	LoginAPI        api.LoginHandler
	RegisterAPI     api.RegisterHandler
	HealthCheckAPI  api.HealthCheck
	LogoutAPI       api.LogoutHandler
	RefreshTokenAPI api.RefreshTokenHandler
}

func InitializeDependencies() *Dependency {
	deps := &Dependency{}

	// Initialize Repositories
	deps.HealthCheckRepo = *repository.NewHealthCheckRepo()
	deps.UserRepo = repository.UserRepository{DB: helpers.DB}

	// Initialize Services
	deps.HealthCheckSvc = service.HealthCheck{
		HealthCheckRepository: &deps.HealthCheckRepo,
	}
	deps.RegisterService = service.RegisterService{
		UserRepo: &deps.UserRepo,
	}
	deps.LoginService = service.LoginService{
		UserRepo: &deps.UserRepo,
	}
	deps.LogoutService = service.LogoutService{
		UserRepo: &deps.UserRepo,
	}
	deps.RefreshTokenService = service.RefreshTokenService{
		UserRepo: &deps.UserRepo,
	}

	// Initialize API Handlers
	deps.HealthCheckAPI = api.HealthCheck{
		HealthCheckService: &deps.HealthCheckSvc,
	}
	deps.RegisterAPI = api.RegisterHandler{
		RegisterService: &deps.RegisterService,
	}
	deps.LoginAPI = api.LoginHandler{
		LoginService: &deps.LoginService,
	}
	deps.LogoutAPI = api.LogoutHandler{
		LogoutService: &deps.LogoutService,
	}
	deps.RefreshTokenAPI = api.RefreshTokenHandler{
		RefreshTokenSvc: &deps.RefreshTokenService,
	}

	return deps
}
