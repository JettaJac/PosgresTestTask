package main

// TODO:
// тестирование
// развертование
// паралелльный запуск
// почистить код
// Документацию
//  инициализировать гит репозиторий
// +? убрать name
// Gorm
// написать комментарии к каждой функции
// Проверить все !!!
// Переименновать RunScript
//возможно стоить прокидывать имя таблице, в которой делаются изменения)б либо прописать ее в конфиге

// TODO:  MacOS указать в документации
import (
	"log/slog"
	"main/internal/app"
	"main/internal/config"
	"main/internal/lib/logger"
)

// curl -i -X POST -H "Content-Type: application/json" -d '{"name":"test6","script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command/save
// curl -i -X GET -H "Content-Type: application/json" -d '{"name":"","script":""}' http://127.0.0.1:8080/commands/all
// curl -i -X GET -H  http://127.0.0.1:8080/command/find?id=5
// curl -i -X DELETE "http://127.0.0.1:8080/command/delete?id=5"
// curl -i -X DELETE -H  http://127.0.0.1:8080/command/delete
// http://localhost:8080/command/find?id=5

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

	log.Info("Server stopped", slog.String("env", config.Env))
}
