package handler

import (
	repo "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/handler"
	repoService "github.com/fajarardiyanto/boilerplate-rest-api/pkg/repository/service"
	"github.com/fajarardiyanto/flt-go-router/interfaces"
	"net/http"
)

type helloHandler struct {
	svc repoService.HelloRepositoryService
}

func NewHelloHandler(svc repoService.HelloRepositoryService) repo.HelloRepositoryHandler {
	return &helloHandler{svc}
}

func (h *helloHandler) Hello(w http.ResponseWriter, r *http.Request) error {
	return interfaces.JSON(w, http.StatusOK, map[string]interface{}{
		"name": h.svc.Hello("Hello"),
	})
}
