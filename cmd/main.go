package main

// !!! MacOS указать в документации
import (
	// "fmt"
	"main/internal/config"

	"fmt"
	"log/slog"
	"main/internal/app"
	sl "main/internal/lib/logger"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	// fmt.Println("Start server")
	config := config.NewConfig()
	log := sl.SetupLogger(config.Env)
	log.Info("Start server", slog.String("env", config.Env))
	log.Debug("Debug messages")

	//  defer db.Close() !!! где-то надо закрыть

	if err := app.Run(config); err != nil { // сделать на Run
		// log.Fatal(err)
		fmt.Println(err)
		// log.Info("Start app", slog.String("env", config.Env))
		log.Error("failed to start server")
	}

	// storage, err := sqlstore.New(config.StoragePath) // Tuz
	// if err != nil {
	// 	log.Error("failed to init storage", sl.Err(err))
	// 	os.Exit(1)
	// }

	// хз
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
	/* // Tuz
	router := chi.NewRouter()

	// router.Post("/url", save.New(log, storage)) //Tuz
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

	log.Error("server stopped")*/
}

// func setupLogger(env string) *slog.Logger { //!!! возможно перенести  другую папку
// 	var log *slog.Logger
// 	switch env {
// 	case envLocal:
// 		log = slog.New(
// 			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
// 		)
// 	case envDev:
// 		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
// 	case envProd:
// 		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
// 	}

// 	return log
// }
