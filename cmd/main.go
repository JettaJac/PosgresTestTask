package main

//TODO: прикрутить авторизацию
// TODO: увеличить количество тестов + моки

import (
	"log/slog"
	"main/internal/app"
	"main/internal/config"
	"main/internal/lib/logger"
)

// curl -i -X POST -H "Content-Type: application/json" -d '{"script":"#!/bin/bash\necho \"Hello, World\""}' http://127.0.0.1:8080/command/save
// curl -i -X GET -H "Content-Type: application/json" -d '{"script":""}' http://127.0.0.1:8080/commands/all
// curl -i -X GET -H "Content-Type: application/json" -d '{"script":""}' "http://127.0.0.1:8080/command/find?id=5"
// curl -i -X DELETE -H "Content-Type: application/json" -d '{"script":""}'  "http://127.0.0.1:8080/command/delete?id=2"
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
}
