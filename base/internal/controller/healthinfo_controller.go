package controller

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/repo"
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"net/http"
)

type healthInfoController struct {
	svc repo.HealthInfoRepositoryService
}

func NewHealthInfoController(svc repo.HealthInfoRepositoryService) repo.HealthInfoRepositoryController {
	return &healthInfoController{
		svc: svc,
	}
}

func (c *healthInfoController) HealthInfo(w http.ResponseWriter, r *http.Request) error {
	res := &dto.ResponseHealthInfo{
		ServiceName:    dto.GetConfig().Name,
		ServiceVersion: dto.GetConfig().Version,
		Compiler:       c.svc.BuildInfo(),
		Database: dto.Database{
			Mysql: c.svc.HealthInfoDatabase(),
			Redis: c.svc.HealthInfoRedis(),
			Mongo: c.svc.HealthInfoMongo(),
		},
	}
	return interfaces.JSON(w, http.StatusOK, res)
}
