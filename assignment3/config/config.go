package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type dbConfig struct {
	Host     string `envconfig:"DB_HOST"`
	Username string `envconfig:"DB_USERNAME"`
	Password string `envconfig:"DB_PASSWORD"`
	Name     string `envconfig:"DB_NAME"`
	Port     int    `envconfig:"DB_PORT"`
}

type redisConfig struct {
	Host     string `envconfig:"REDIS_ADDRESS"`
	Password string `envconfig:"REDIS_PASSWORD"`
	Port     int    `envconfig:"REDIS_PORT"`
	DB       int    `envconfig:"REDIS_DB"`
}

type Config struct {
	DB    dbConfig
	Redis redisConfig
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error loading .env file: %w", err)
	}
	if err := envconfig.Process("MYAPP", cfg); err != nil {
		return nil, fmt.Errorf("error processing MYAPP env: %w", err)
	}
	return cfg, nil
}
