package service

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/config"
	repo "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/service"
)

type service struct{}

func NewHealthInfoService() repo.HealthInfoRepositoryService {
	return &service{}
}

func (*service) HealthInfoDatabase() bool {
	var version string
	err := config.GetDBConn().Orm().Raw(`SELECT @@VERSION`).Scan(&version).Error
	if err != nil {
		return false
	}

	return true
}
