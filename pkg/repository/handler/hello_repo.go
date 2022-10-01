package repo

import "net/http"

type HelloRepositoryHandler interface {
	Hello(w http.ResponseWriter, r *http.Request) error
}
