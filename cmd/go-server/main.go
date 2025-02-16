package main

import (
	"fmt"
	"go_server/internal/config"
	"go_server/internal/storage"
	"log/slog"
	"os"
)

var (
	mapStore = storage.CreateStorage()
)

func main() {

	//TODO init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	//todo init logger
	log := loggerSetup(cfg.Env)

	// log.With(slog.String("env", cfg.Env))

	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("debugging working")

	//todo init storage

	//todo init router

	//todo run server

}

func loggerSetup(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		{
			log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}
	case "dev":
		{
			log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
		}

	}
	return log
}
