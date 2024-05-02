
package config

import (
	"time"
	"fmt"
	"flag"
	"io/ioutil"
	// "os"
	"log"
	"gopkg.in/yaml.v3"

)

type Config struct {
	Env string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	DatabaseURL string `yaml:"database_url"` // The database URL to use !!! возможно прописать путь к базе данных здесь)
	HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
  Address string `yaml:"address" env-default:"localhost:8080"`
  Timeout time.Duration `yaml:"timeout" env-default:"4s"` 
  // !!! Возможно не надо, убрать в будущем
  IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}



// // NewConfig creates a new configuration
// func NewConfig() *Config {
// 	return &Config{
// 		BindAddr: ":8080",
// 		LogLevel: "debug",
// 		// Store:    store.NewConfig(),
// 	}
// }


// NewConfig creates a new configuration

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

	configData, err := ioutil.ReadFile(configPath)
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