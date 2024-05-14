package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	DatabaseURL string `yaml:"databaseURL" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
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
	fmt.Println(string(configData))
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("Error parsing YAML configuration data %v", err)
	}

	return &config
}
