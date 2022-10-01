package api

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/app/handler"
	"github.com/fajarardiyanto/boilerplate-rest-api/app/service"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/boilerplate-rest-api/pkg/middleware"
	"github.com/fajarardiyanto/flt-go-router/lib"
)

func Start() {
	// init configuration database
	config.Init()

	router := lib.New(dto.GetConfig().Version)
	router.Use(
		middleware.MiddlewareLogger(),
		middleware.MiddlewareCors(),
		middleware.MiddlewareError(),
	)

	// Service
	helloSvc := service.NewHelloService()

	// Handler
	helloHandler := handler.NewHelloHandler(helloSvc)

	// router
	router.GET("", helloHandler.Hello)

	router.Run(dto.GetConfig().Port)
}
