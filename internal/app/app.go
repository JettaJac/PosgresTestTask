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

// / TODO:  Добавить гоурутину
func Run(config *config.Config) error {
	log := sl.SetupLogger(config.Env)

	storage, err := sqlstore.NewDB(config.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err)) // /  пока почему то не подключаеться лог/слог
		os.Exit(1)
	}
	log.Info("Сonnected to the database", slog.String("env", config.Env))
	defer storage.CloseDB() /// TODO:  Возможно сделать без вызова доп функции, а сразу закрыть здесь

	// Created server
	/*srv := */
	srv := server.NewServer(config, storage, log)
	log.Info("Starting server", slog.String("address", config.Address))

	_ = srv //!!!

	// srv.ListenAndServe(config.Address, nil)
	// fmt.Println("TTTTTTT", "YYYYYY")

	return http.ListenAndServe(config.Address, srv)
}
