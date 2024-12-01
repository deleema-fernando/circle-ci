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
			return log.LoggerWithContext(context.Background(), logger)
		},
		Handler: app.Routes(),
	}

	serverShutdownChan := make(chan struct{})

	go func() {
		<-ctx.Done()

		logger.Info("shutting down the server gracefully")

		timeOutCtx, cancelTimeout := context.WithTimeout(ctx, 5*time.Second)
		defer cancelTimeout()

		if err := server.Shutdown(timeOutCtx); err != nil {
			logger.Error("server shutdown error", "error", err)
		}

		<-timeOutCtx.Done()
		logger.Info("server shutdown gracefully")
		time.Sleep(1 * time.Second)
		serverShutdownChan <- struct{}{}
	}()

	if err := server.ListenAndServe(); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			logger.Warn("server closed")
			return
		}

		logger.Error("server error", "error", err)

		return
	}

	<-serverShutdownChan
}
