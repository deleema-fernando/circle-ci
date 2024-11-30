package log

import (
	"context"
	"log/slog"
	"os"
)

type loggerKeyType string

const loggerKeyTypeKey loggerKeyType = "logger"

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

	return logger
}
