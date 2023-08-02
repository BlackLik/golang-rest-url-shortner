package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env:"ENV" envDefault:"development"`
	StoragePath string `yaml:"storage_path" env:"STORAGE_PATH" envDefault:"./storage/storage.db"`
	HttpServer  `yaml:"http_server"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env:"HTTP_ADDRESS" envDefault:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env:"HTTP_TIMEOUT" envDefault:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"HTTP_IDLE_TIMEOUT" envDefault:"60s"`
}

func MustLoad(pathStr ...string) *Config {
	var configPath string
	if len(pathStr) == 0 || pathStr[0] == "" {
		configPath = os.Getenv("CONFIG_PATH")
		if configPath == "" {
			log.Fatal("CONFIG_PATH is not set")
		}
	} else {
		configPath = pathStr[0]
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file not found: %s", configPath)
	}

	var config Config

	if err := cleanenv.ReadConfig(configPath, &config); err != nil {
		log.Fatal(err)
	}

	return &config

}
