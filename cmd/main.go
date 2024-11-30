package main

import (
	"context"
	"errors"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"super-heroes/internal/app"
	"super-heroes/internal/pkg/log"
)

var version string

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger = logger.With("version", version)

	logger.Info("starting the heroes service")
	defer logger.Info("stopping the heroes service")

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	app := app.New(app.Config{})

	server := &http.Server{
		Addr: ":1990",
		BaseContext: func(l net.Listener) context.Context {
			return log.LoggerWithContext(ctx, logger)
		},
		Handler: app.Routes(),
	}

	server.RegisterOnShutdown(func() {
		logger.Info("server shutting down")
		time.Sleep(2 * time.Second)
	})

	go func() {
		if err := server.ListenAndServe(); err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Warn("server closed")
				return
			}

			logger.Error("server error", "error", err)

			return
		}
	}()

	<-ctx.Done()

	timeOutCtx, cancelTimeout := context.WithTimeout(ctx, 5*time.Second)
	defer cancelTimeout()

	defer func() {
		if err := server.Shutdown(timeOutCtx); err != nil {
			logger.Error("server shutdown error", "error", err)
		}
	}()

	<-timeOutCtx.Done()
}
