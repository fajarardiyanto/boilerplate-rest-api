package dto

import (
	"github.com/fajarardiyanto/flt-go-database/interfaces"
	"github.com/fajarardiyanto/flt-go-utils/flags"
)

var cfg = new(Config)

type Config struct {
	Version  string `yaml:"version"`
	Name     string `yaml:"name"`
	Port     string `yaml:"port"`
	Database struct {
		Mysql interfaces.SQLConfig `yaml:"mysql"`
	} `yaml:"database"`
}

func init() {
	flags.Init("external/config/config.yaml", cfg)
}

func GetConfig() *Config {
	return cfg
}
