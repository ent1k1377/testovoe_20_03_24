package util

import (
	"os"
)

type Config struct {
	DatabaseDriver string
	DatabaseURL    string
}

func NewConfig() *Config {
	return &Config{
		DatabaseDriver: os.Getenv("DB_DRIVER"),
		DatabaseURL:    os.Getenv("DB_URL"),
	}
}
