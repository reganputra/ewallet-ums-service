package service

import "ewallet-ums/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.HealthCheckRepository
}

func (h *HealthCheck) HealthCheckService() (string, error) {
	return h.HealthCheckRepository.HealthCheckRepository()
}
