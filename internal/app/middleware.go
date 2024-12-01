package app

import (
	"context"
	"net/http"

	"super-heroes/internal/pkg/log"

	"github.com/google/uuid"
)

const requestIDHeader = "X-Request-ID"

func RequestIDMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := r.Header.Get(requestIDHeader)
		if requestID == "" {
			requestID = uuid.New().String()
		}

		ctx := context.WithValue(r.Context(), log.RequestIDKey, requestID)
		r = r.WithContext(ctx)

		w.Header().Set(requestIDHeader, requestID)

		next(w, r)
	})
}
