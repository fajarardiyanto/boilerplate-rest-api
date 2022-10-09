package main

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/router"
	"github.com/fajarardiyanto/boilerplate-rest-api/pkg/middleware"
	"github.com/fajarardiyanto/flt-go-router/lib"
)

func main() {
	//if err := Run(os.Args[1:]); err != nil {
	//	config.GetLogger().Error("Unable to run the command %s ", err.Error())
	//}
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
}

//
//var rootCmd = &cobra.Command{
//	Use:   dto.GetConfig().Name,
//	Short: dto.GetConfig().Name,
//}
//
//func init() {
//	rootCmd.AddCommand(api.CmdAPI)
//	rootCmd.AddCommand(generate.CmdConfig)
//}
//
//// Run function lets you run the commands
//func Run(args []string) error {
//	rootCmd.SetArgs(args)
//	return rootCmd.Execute()
//}
