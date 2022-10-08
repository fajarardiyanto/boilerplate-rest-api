package repo

import (
	"net/http"
)

type (
	HealthInfoRepositoryController interface {
		HealthInfo(w http.ResponseWriter, r *http.Request) error
	}

	HealthInfoRepositoryService interface {
		BuildInfo() string
		HealthInfoDatabase() bool
		HealthInfoRedis() bool
		HealthInfoMongo() bool
	}
)
