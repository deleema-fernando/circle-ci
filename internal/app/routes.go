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

		if _, err := w.Write([]byte("OK")); err != nil {
			log.Error("failed to write response", "error", err)
			w.WriteHeader(http.StatusInternalServerError)

			return
		}
	})

	return mux
}
