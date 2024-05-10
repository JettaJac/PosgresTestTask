package main

// TODO:
// тестирование
// + логи прокинуть нормально
// развертование
// паралелльный запуск
// почистить код
// перенести тесты в отдельную папку
// Документацию
// +имя таблицы вывести в константу 3-е 6в 7.02
// + allпереписать на селект, 3-е 9в 7.18
//  инициализировать гит репозиторий
// +? убрать name

// TODO:  MacOS указать в документации
import (
	// "fmt"
	"log/slog"
	"main/internal/app"
	"main/internal/config"
	sl "main/internal/lib/logger"
)

// curl -i -X POST -H "Content-Type: application/json" -d '{"name":"test2","script_file":"testscript2.sh"}' http://localhost:8080/command
// curl -i -X POST -H "Content-Type: application/json" -d '{"name":"test6","script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command/save
// curl -i -X GET -H "Content-Type: application/json" -d '{"name":"","script":""}' http://127.0.0.1:8080/commands/all
// curl -i -X GET -H "Content-Type: application/json" -d '{"id":"43","script":""}' http://127.0.0.1:8080/command/find
// curl -i -X DELETE -H "Content-Type: application/json" -d '{"id":45,"script":""}' http://127.0.0.1:8080/command/delete
// http://localhost:8080/command/find?id=5
// curl -i -X DELETE "http://127.0.0.1:8080/command/delete?id=5"

// curl -i -X POST -H "Content-Type: application/json" -d '{"script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command/save

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	config := config.NewConfig()
	log := sl.SetupLogger(config.Env)
	log.Info("Start app", slog.String("env", config.Env), slog.String("version", "1.0"))
	log.Debug("debug messages are enabled")

	if err := app.Run(config); err != nil { // сделать на Run
		log.Error("failed to start server")
	}
	// log.Info("Start app", slog.String("env", config.Env))

	/*
		storage, err := sqlstore.New(config.StoragePath) // Tuz
		if err != nil {
			log.Error("failed to init storage", sl.Err(err))
			os.Exit(1)
		}

		Tuz
		router := chi.NewRouter()

		router.Post("/url", save.New(log, storage)) //Tuz
		log.Info("starting server", slog.String("address", config.Address))

		srv := &http.Server{
			Addr:         config.Address,
			Handler:      router,
			ReadTimeout:  config.HTTPServer.Timeout,
			WriteTimeout: config.HTTPServer.Timeout,
			IdleTimeout:  config.HTTPServer.IdleTimeout,
		}

		if err := srv.ListenAndServe(); err != nil { // Tuz
			log.Error("failed to start server")
		}
	*/
	log.Info("Server stopped", slog.String("env", config.Env))
}
