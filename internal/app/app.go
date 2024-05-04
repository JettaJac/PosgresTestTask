package app

import (
	"fmt"
	"log/slog"
	"main/internal/config"
	sl "main/internal/lib/logger"

	"main/internal/server"
	"main/internal/storage/sqlstore"
	"net/http"
	"os"
	// "main/internal/config"
	// "fmt"
)

func Run(config *config.Config) error {
	fmt.Println("Запуск приложения")
	log := sl.SetupLogger(config.Env)

	// Created BaseData
	storage, err := sqlstore.NewDBMY(config.StoragePath)
	if err != nil {
		// log.Error("failed to init storage", sl.Err(err)) // /  пока почему то не подключаеться лог/слог
		os.Exit(1)
	}
	defer storage.CloseDB() /// !!! Возможно сделать без вызова доп функции, а сразу закрыть здесь

	// Created server
	/*srv := */
	server.NewServer(config, storage /*, log*/)
	log.Info("starting server", slog.String("address", config.Address))

	/*_ = srv*/

	return http.ListenAndServe(config.Address, nil)
}
