package main

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"dev"`
	ServerPort  string `env:"SERVER_PORT" envDefault:":8080"`
	DBHost      string `env:"DB_HOST" envDefault:"localhost"`
	DBPort      int    `env:"DB_PORT" envDefault:"5432"`
	DBName      string `env:"DB_NAME" envDefault:"coffeeshop"`
	DBUser      string `env:"DB_USER" envDefault:"postgres"`
	DBPassword  string `env:"DB_PASSWORD" envDefault:"coffeeshop"`
}

func InitializeConfig() *Config {
	// loading configuration variables from '.env' file.
	err := godotenv.Load()
	if err != nil {
		log.Println("unable to load .env file:", err)
	}

	cfg := Config{}
	// loading configuration variables from environment variables.
	err = env.Parse(&cfg)
	if err != nil {
		log.Fatal("unable to load environment variables:", err)
	}

	return &cfg
}
