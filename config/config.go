package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
	"urlshort.ru/m/logs"
)

type Config struct {
	DB_NAME        string `env:"DB_NAME"`
	LOGGER_LEVEL   string `env:"LOGGER_LEVEL"`
	SECRET_KEY_JWT string `env:"SECRET_KEY_JWT"`
	TIME_ZONE      string `env:"TIME_ZONE"`
}

var ERROR_HANDLER string = "config"

var ConfigAll *Config

// init is a built-in Go function that is automatically called before the main function.
//
// It is commonly used to perform initialization tasks, such as setting the log level.
// This function does not take any parameters and does not return any values.
func init() {
	ConfigAll = getConfig()
	logs.SetLevel(ConfigAll.LOGGER_LEVEL)
	time.Local = time.FixedZone(ConfigAll.TIME_ZONE, 0)
}

// getConfig retrieves the configuration by loading the environment variables from a specific file.
//
// It returns a pointer to a Config struct.
func getConfig() *Config {
	config := &Config{}
	err := godotenv.Load("./config/.env")
	if err != nil {
		slog.Error(ERROR_HANDLER, err)
	}
	config.DB_NAME = os.Getenv("DB_NAME")
	config.LOGGER_LEVEL = os.Getenv("LOGGER_LEVEL")
	config.SECRET_KEY_JWT = os.Getenv("SECRET_KEY_JWT")
	config.TIME_ZONE = os.Getenv("TIME_ZONE")

	if config.LOGGER_LEVEL == "" {
		slog.Error(ERROR_HANDLER, "DEBUG")
	}

	if config.TIME_ZONE == "" {
		config.TIME_ZONE = "Europe/Moscow"
	}

	return config
}
