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

	// Health Check
	healthCheckRepo := repository.NewHealthCheckRepo()
	healthCheckSvc := &service.HealthCheck{
		HealthCheckRepository: healthCheckRepo,
	}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	// Register
	registerRepo := &repository.RegisterRepository{
		DB: helpers.DB,
	}
	registerService := &service.RegisterService{
		RegisterRepo: registerRepo,
	}
	registerApi := &api.RegisterHandler{
		RegisterService: registerService,
	}

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandler)

	userV1 := r.Group("/users/v1")
	userV1.POST("/register", registerApi.Register)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
