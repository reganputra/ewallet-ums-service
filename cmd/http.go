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

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}

type Dependency struct {
	// Repositories
	UserRepo        repository.UserRepository
	HealthCheckRepo *repository.HealthCheckRepo

	// Services
	LoginService    service.LoginService
	RegisterService service.RegisterService
	HealthCheckSvc  *service.HealthCheck

	// API Handlers
	LoginAPI       api.LoginHandler
	RegisterAPI    api.RegisterHandler
	HealthCheckAPI api.HealthCheck
}

func InitializeDependencies() *Dependency {
	deps := &Dependency{}

	// Initialize Repositories
	deps.HealthCheckRepo = repository.NewHealthCheckRepo()
	deps.UserRepo = repository.UserRepository{
		DB: helpers.DB,
	}

	// Initialize Services
	deps.HealthCheckSvc = &service.HealthCheck{
		HealthCheckRepository: deps.HealthCheckRepo,
	}

	deps.RegisterService = service.RegisterService{
		UserRepo: &deps.UserRepo,
	}

	deps.LoginService = service.LoginService{
		UserRepo: &deps.UserRepo,
	}

	// Initialize API Handlers
	deps.HealthCheckAPI = api.HealthCheck{
		HealthCheckService: deps.HealthCheckSvc,
	}

	deps.RegisterAPI = api.RegisterHandler{
		RegisterService: &deps.RegisterService,
	}

	deps.LoginAPI = api.LoginHandler{
		LoginService: &deps.LoginService,
	}

	return deps
}
