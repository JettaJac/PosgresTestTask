package main
// !!! MacOS указать в документации
import (
	// "fmt"
	"main/internal/config"
	"main/internal/storage/sqlstore"
	"main/internal/lib/logger"
	// "main/internal/app"
	"log/slog"
	"fmt"
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

	//  defer db.Close() !!! где-то надо закрыть


	// if err := app.Run(config); err != nil { // сделать на Run
	// 	// log.Fatal(err)
	// }

	storage, err := sqlstore.New(config.StoragePath)
	if err!= nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	// id, err := storage.SaveScript("qqqqqqqq", "111")
	// if err!= nil {
	// 	log.Error("failed to save script", sl.Err(err))
	// 	os.Exit(1)
	// }
	// _ = id

	// id, err = storage.SaveScript("qqqqqqqq", "111")
	// if err!= nil {
	// 	log.Error("failed to save script", sl.Err(err))
	// 	os.Exit(1)
	// }
	// _ = id

		fmt.Println(err, storage)
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