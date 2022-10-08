package api

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/router"
	"github.com/fajarardiyanto/boilerplate-rest-api/pkg/middleware"
	"github.com/fajarardiyanto/flt-go-router/lib"
	"github.com/spf13/cobra"
)

var CmdAPI = &cobra.Command{
	Use:   "api",
	Short: "Start api server",
	RunE:  Api,
}

func Api(cmd *cobra.Command, args []string) error {
	// init configuration database
	config.Config()

	r := lib.New(dto.GetConfig().Version)
	r.Use(
		middleware.MiddlewareLogger(),
		middleware.MiddlewareCors(),
		middleware.MiddlewareError(),
	)

	{
		router.HealthInfoRouter(r)
	}

	r.Run(dto.GetConfig().Port)
	return nil
}
