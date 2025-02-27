package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"go_server/internal/config"
	"go_server/internal/http-server/handlers"
	"go_server/internal/storage"
)

// Invoke-WebRequest -Uri http://localhost:8080/users -Method POST -Headers $headers -Body '{"Id":1,"Name":"Alice","Card_pin":1234}'
// curl "http://localhost:8080/users?id=1"

func main() {
	mapStore := storage.CreateStorage()

	// init config
	cfg := config.MustLoad()
	fmt.Println(cfg)

	// init logger
	log := loggerSetup(cfg.Env)

	// log.With(slog.String("env", cfg.Env))

	log.Info("starting server", slog.String("env", cfg.Env))
	log.Debug("debugging working")

	// init storage

	// init router
	router := chi.NewRouter()

	h := handlers.NewHandler(mapStore, log)

	router.Post("/users", h.AddUserHandler)
	router.Get("/users", h.GetUserHandler)

	//todo run server

	log.Info("starting server...")

	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Error(err.Error())
	}
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
