package config

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/flt-go-logger/interfaces"
	"github.com/fajarardiyanto/flt-go-logger/lib"
)

var logger interfaces.Logger

func GetLogger() interfaces.Logger {
	return logger
}

func init() {
	logger = lib.NewLib()
	logger.Init(dto.GetConfig().Name)
}
