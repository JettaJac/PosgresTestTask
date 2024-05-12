package config

import (
	"flag"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	DatabaseURL string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func NewConfig() *Config {

	/* альтернатива
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH not set")
	}

	//Check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Config file %s not found", configPath)
	}
	*/
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

/*
// go get github.com/ilyakaznacheev/cleanenv
func NewConfig() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		lof.Fatal("CONFIG_PATH not set")
}

//Check if file exists
if _,err := os.Stat(configPath); os.IsNotExist(err) {
	log.Fatalf("Config file %s not found", configPath)
}

var config Config
if err := cleanenv.ReadConfig(configPath, &config); err!= nil {
	log.Fatalf("Failed to read config file: %s", err)
}
return config
}
*/
