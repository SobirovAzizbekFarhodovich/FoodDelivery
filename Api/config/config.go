package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	AUTH_PORT     string
	COMPANY_PORT  string
	PROGRESS_PORT string

	DB_HOST     string
	DB_PORT     int
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string

	LOG_PATH string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.AUTH_PORT = cast.ToString(getEnv("AUTH_PORT", ":8072"))
	config.COMPANY_PORT = cast.ToString(getEnv("COMPANY_PORT", ":50051"))

	config.DB_HOST = cast.ToString(getEnv("DB_HOST", "gateway"))
	config.DB_PORT = cast.ToInt(getEnv("DB_PORT", 5432))
	config.DB_USER = cast.ToString(getEnv("DB_USER", "azizbek"))
	config.DB_PASSWORD = cast.ToString(getEnv("DB_PASSWORD", "123"))
	config.DB_NAME = cast.ToString(getEnv("DB_NAME", "delivery"))

	config.LOG_PATH = cast.ToString(getEnv("LOG_PATH", "logs/info.log"))

	return config
}

func getEnv(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
