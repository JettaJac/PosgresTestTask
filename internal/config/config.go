package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/caarlos0/env/v6"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Env          string `yaml:"env" env-default:"local"`
	AuthBase     string `yaml:"authbase" env-default:""`
	NameDataBase string `yaml:"namebase" env-default:"restapi_script"`
	Flags        string `yaml:"flags" env-default:"sslmode=disable"`
	DatabaseURL  string `yaml:"databaseURL" env:"DATABASE_HOST" env-default:"localhost" env-required:"true"`
	HTTPServer   `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:":8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

// NewConfig returns a new config instance
func NewConfig() *Config {

	var configPath string
	flag.StringVar(&configPath, "config-path", "configs/appConfig.yaml", "config file path")

	flag.Parse()
	configData, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Failed to read config file: %s", err)
	}

	var config Config

	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML configuration data %v", err)
	}

	if err := env.Parse(&config); err != nil {
		log.Fatalf("Error parsing environment variables %v", err)
	}

	if os.Getenv("DATABASE_HOST") == "db" {
		config.AuthBase = "user:password"
	}
	config.DatabaseURL = fmt.Sprintf("postgres://%s@%s:5432/%s?%s", config.AuthBase, os.Getenv("DATABASE_HOST"), config.NameDataBase, config.Flags)

	return &config
}
