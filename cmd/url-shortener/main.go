package main

import (
	"fmt"
	"log/slog"
	"os"
	"url-shortener/internal/config"
	"url-shortener/internal/lib/logger/sl"
	"url-shortener/internal/storage/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(*cfg)

	log := setupLogger(&cfg.Env)
	log.Info("Starting url-shortener", slog.String("env", cfg.Env))

	storage, err := sqlite.InitStorage(cfg.StoragePath)
	if err != nil {
		log.Error("Storage can't be initialized!", sl.Err(err))
		os.Exit(1)
	}

	_ = storage

	
	// init router: chi (or net/http), render

	// run server
}

func setupLogger(env *string) *slog.Logger {
	var log *slog.Logger

	switch *env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	default:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		*env = envLocal
	}

	return log

}
