package main

import (
	"os"

	"golang.org/x/exp/slog"
	"urlshort.ru/m/internal/config"
	"urlshort.ru/m/internal/lib/logger/sl"
	"urlshort.ru/m/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envProd  = "prod"
	envDev   = "dev"
)

func main() {
	// TODO: init config: cleanenv

	cfg := config.MustLoad("./config/local.yaml")
	// TODO: init logger: slog

	log := setupLogger(cfg.Env)
	log.Info("Starting...", slog.String("env", cfg.Env))
	log.Debug("Debug messages are enabled")

	// TODO: init db: sqlite

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("Failed to create storage", sl.Err(err))
		os.Exit(1)
	}
	_ = storage
	// TODO: init http server
	// TODO: init router: chi, "chi render"
	// TODO: run http server
}

func setupLogger(env string) *slog.Logger {

	var log *slog.Logger

	switch env {
	case envLocal:
		// TODO: init logger: slog
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		// TODO: init logger: slog
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
