package config

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	DB       *DB
	BasePath string `env:"BASE_PATH"`
	Address  string `env:"ADDRESS"`
}

type DB struct {
	DBHost string `env:"DB_HOST"`
	DBPort string `env:"DB_PORT"`
	DBUser string `env:"DB_USER"`
	DBPass string `env:"DB_PASS"`
	DBName string `env:"DB_NAME"`
}

func New(envFile string) (*Config, error) {
	appEnv := Config{DB: new(DB)}
	err := godotenv.Load(envFile)
	if err != nil {
		return nil, err
	}

	err = env.Parse(&appEnv)
	if err != nil {
		return nil, err
	}

	return &appEnv, nil
}
