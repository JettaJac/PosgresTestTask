package app

import (
	"log/slog"
	"main/internal/config"
	"main/internal/lib/logger"
	"main/internal/server"
	"main/internal/storage/sqlstore"
	"net/http"
	"os"
)

// Run starts the application
func Run(config *config.Config) error {
	log := sl.SetupLogger(config.Env)

	storage, err := sqlstore.NewDB(config.DatabaseURL)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}
	log.Info("Сonnected to the database", slog.String("env", config.Env))
	defer storage.CloseDB()

	// Created server
	srv := server.NewServer(config, storage, log)
	log.Info("Starting server", slog.String("address", config.Address))

	return http.ListenAndServe(config.Address, srv)
}
