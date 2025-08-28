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

	healthCheckRepo := repository.NewHealthCheckRepo()
	healthCheckSvc := &service.HealthCheck{
		HealthCheckRepository: healthCheckRepo,
	}
	healthCheckApi := api.HealthCheck{
		HealthCheckService: healthCheckSvc,
	}

	r := gin.Default()
	r.GET("/health", healthCheckApi.HealthCheckHandler)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal("Failed to start HTTP server:", err)
	}
}
