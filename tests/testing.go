package tests

import (
	"main/internal/config"
	"time"
)

func testNewConfig() *config.Config {
	config := &config.Config{}
	config.HTTPServer.Address = "localhost:8080"
	config.StoragePath = "host=localhost dbname=restapi_test sslmode=disable"
	Timeout, _ := time.ParseDuration("4s")
	IdleTimeout, _ := time.ParseDuration("1m")
	config.Env = "local"
	config.Address = "localhost:8080"
	config.HTTPServer.Timeout = Timeout
	config.HTTPServer.IdleTimeout = IdleTimeout
	return config
}
