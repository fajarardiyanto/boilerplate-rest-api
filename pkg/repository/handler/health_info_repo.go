package repo

import "net/http"

type HealthInfoRepositoryHandler interface {
	HealthInfo(w http.ResponseWriter, r *http.Request) error
}
