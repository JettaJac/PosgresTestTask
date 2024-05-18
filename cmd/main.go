package main

// !!! копираование и запуск из своей папки!
// убрать вывод логов при тестах
// приложение работает в докере

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

// Main function
func main() {
	config := config.NewConfig()
	log := sl.SetupLogger(config.Env)
	log.Info("Start app", slog.String("env", config.Env), slog.String("version", "1.0"))
	log.Debug("debug messages are enabled")

	if err := app.Run(config); err != nil {
		log.Error("failed to start server", err)
	}

	log.Info("Server stopped", slog.String("env", config.Env))
	log.Info("Server stopped", slog.String("address", config.Address)) // !!!
}
