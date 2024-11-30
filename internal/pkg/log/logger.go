package log

import (
	"context"
	"log/slog"
)

type loggerKeyType string

const loggerKeyTypeKey loggerKeyType = "logger"

func LoggerWithContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKeyTypeKey, logger)
}

func LoggerFromContext(ctx context.Context) *slog.Logger {
	logger := ctx.Value(loggerKeyTypeKey).(*slog.Logger)

	return logger
}
