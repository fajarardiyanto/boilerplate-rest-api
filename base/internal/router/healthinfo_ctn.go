package router

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/controller"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/services"
	"github.com/fajarardiyanto/flt-go-router/interfaces"
)

func HealthInfoRouter(r interfaces.Routers) {
	// Service
	svc := services.NewHealthInfoService()

	// Controller
	health := controller.NewHealthInfoController(svc)

	r.GET("/health-info", health.HealthInfo)
}
