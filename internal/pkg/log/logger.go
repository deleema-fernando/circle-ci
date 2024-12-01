package log

import (
	"context"
	"log/slog"
	"os"
)

type contextKey string

const (
	loggerKeyTypeKey contextKey = "logger"
	RequestIDKey     contextKey = "request_id"
)

func LoggerWithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKeyTypeKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	if ctx == nil {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	logger, ok := ctx.Value(loggerKeyTypeKey).(*slog.Logger)
	if !ok {
		return slog.New(slog.NewTextHandler(os.Stdout, nil))
	}

	requestID, ok := ctx.Value(RequestIDKey).(string)
	if ok {
		logger = logger.With("request_id", requestID)
	}

	return logger
}
