package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"amerta/backend/internal/auth"
	"amerta/backend/internal/config"
	"amerta/backend/internal/database"
	"amerta/backend/internal/httpapi"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	cfg := config.Load()
	if cfg.JWTSecret == config.DevSecret {
		logger.Warn("JWT_SECRET is not set; using development secret")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	pool, err := database.Connect(ctx, cfg.DatabaseURL)
	cancel()
	if err != nil {
		logger.Error("database", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	server := &http.Server{
		Addr: cfg.Addr,
		Handler: httpapi.NewRouter(httpapi.Deps{
			Pool:       pool,
			Tokens:     auth.New(cfg.JWTSecret, cfg.AccessTTL),
			Logger:     logger,
			AccessTTL:  cfg.AccessTTL,
			RefreshTTL: cfg.RefreshTTL,
			Secure:     cfg.CookieSecure,
			Origins:    cfg.Origins,
		}),
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		logger.Info("server started", "addr", cfg.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("server", "error", err)
			os.Exit(1)
		}
	}()

	stop, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	<-stop.Done()

	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Error("shutdown", "error", err)
	}
}
