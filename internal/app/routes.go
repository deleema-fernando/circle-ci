package app

import (
	"net/http"

	"super-heroes/internal/pkg/log"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type App interface {
	Routes() *http.ServeMux
}

type Config struct {
	NR *newrelic.Application
}

type heroes struct {
	nr *newrelic.Application
}

func New(conf Config) App {
	return &heroes{
		nr: conf.NR,
	}
}

func (r *heroes) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", r.applyMiddlewares(r.healthHandler))

	return mux
}

func (r *heroes) applyMiddlewares(handler http.HandlerFunc) http.HandlerFunc {
	return RequestIDMiddleware(handler)
}

func (h *heroes) healthHandler(w http.ResponseWriter, r *http.Request) {
	txn := h.nr.StartTransaction("health")
	defer txn.End()

	log := log.LoggerFromContext(r.Context())
	log.Info("health check")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write([]byte("OK")); err != nil {
		log.Error("failed to write response", "error", err)
		w.WriteHeader(http.StatusInternalServerError)

		return
	}
}
