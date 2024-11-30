package app

import (
	"net/http"
	"super-heroes/internal/pkg/log"
)

type App interface {
	Routes() *http.ServeMux
}

type Config struct{}

type heroes struct{}

func New(conf Config) App {
	return &heroes{}
}

func (r *heroes) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		log := log.LoggerFromContext(r.Context())
		log.Info("health check")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	return mux
}
