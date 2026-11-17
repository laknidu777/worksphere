package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/worksphere/worksphere/apps/api/internal/config"
	"github.com/worksphere/worksphere/apps/api/internal/database"
	"github.com/worksphere/worksphere/apps/api/internal/server"
)

func main() {
	cfg := config.Load()

	// Initialize structured logging
	var logHandler slog.Handler
	if cfg.Environment == "production" {
		logHandler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})
	} else {
		logHandler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})
	}
	slog.SetDefault(slog.New(logHandler))

	slog.Info("starting WorkSphere API server",
		slog.String("env", cfg.Environment),
		slog.String("port", cfg.Port),
	)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Connect to PostgreSQL (non-fatal if offline in local dev mode, but readiness probe will report not ready)
	db, err := database.Connect(ctx, cfg.DatabaseURL)
	if err != nil {
		slog.Warn("database connection failed on startup; readiness probe will report not ready until connected",
			slog.String("error", err.Error()),
		)
	} else {
		defer db.Close()
	}

	// Connect to Redis (optional/cache; non-fatal if offline)
	redisClient, err := database.ConnectRedis(ctx, cfg.RedisURL)
	if err != nil {
		slog.Warn("redis connection failed on startup; running without redis cache",
			slog.String("error", err.Error()),
		)
	} else {
		defer redisClient.Close()
	}

	srv := server.New(cfg, db, redisClient)
	httpServer := srv.HTTPServer()

	// Start HTTP server in goroutine
	serverErrors := make(chan error, 1)
	go func() {
		slog.Info("listening for HTTP requests", slog.String("addr", httpServer.Addr))
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErrors <- err
		}
	}()

	// Listen for shutdown signals
	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		slog.Error("fatal server error", slog.String("error", err.Error()))
		os.Exit(1)
	case sig := <-shutdown:
		slog.Info("shutdown signal received", slog.String("signal", sig.String()))

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), cfg.ShutdownTimeoutSeconds)
		defer shutdownCancel()

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			slog.Error("graceful shutdown failed", slog.String("error", err.Error()))
			_ = httpServer.Close()
		}
		slog.Info("server exited cleanly")
	}
}
