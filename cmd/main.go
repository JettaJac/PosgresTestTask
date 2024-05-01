package main

import (
	// "fmt"
	"main/internal/config"
	"log/slog"
	"os"

)
const (
	envLocal = "local"
	envDev = "dev"
	envProd = "prod"	
)

func main() {
	// fmt.Println("Start server")
	config := config.NewConfig()
	log:= settupLogger(config.Env)
	log.Info("Start server", slog.String("env", config.Env))
	log.Debug("Debug messages")
}

func settupLogger(env string) *slog.Logger { //!!! возможно перенести  другую папку
	var log *slog.Logger
	switch env {
		case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case envProd:
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}