package handler

import (
	"github.com/fajarardiyanto/boilerplate-rest-api/internal/dto"
	repo "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/handler"
	repoService "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/service"
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"net/http"
)

type healthInfoHandler struct {
	svc repoService.HealthInfoRepositoryService
}

func NewHealthInfoHandler(svc repoService.HealthInfoRepositoryService) repo.HealthInfoRepositoryHandler {
	return &healthInfoHandler{
		svc: svc,
	}
}

func (h *healthInfoHandler) HealthInfo(w http.ResponseWriter, r *http.Request) error {
	return interfaces.JSON(w, http.StatusOK, dto.ResponseHealthInfo{
		ServiceName:    dto.GetConfig().Name,
		ServiceVersion: dto.GetConfig().Version,
		Database: dto.Database{
			Mysql: h.svc.HealthInfoDatabase(),
		},
	})
}
