package services

import (
	"fmt"
	"github.com/fajarardiyanto/boilerplate-rest-api/config"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/repo"
	"github.com/shirou/gopsutil/v3/host"
)

type service struct{}

func NewHealthInfoService() repo.HealthInfoRepositoryService {
	return &service{}
}

func (*service) BuildInfo() (OSBuildName string) {
	if info, err := host.Info(); err == nil {
		OSBuildName = fmt.Sprintf("%s (%s)",
			info.PlatformFamily, info.PlatformVersion)
	}

	return OSBuildName
}

func (*service) HealthInfoDatabase() bool {
	var version string
	err := config.GetDBConn().Orm().Raw("SELECT @@VERSION").Scan(&version).Error
	if err != nil {
		return false
	}

	return true
}

func (*service) HealthInfoRedis() bool {
	if err := config.GetRdbConn().Init(); err != nil {
		return false
	}

	return true
}

func (*service) HealthInfoMongo() bool {
	if err := config.GetMongoConn().Init(); err != nil {
		return false
	}

	return true
}
