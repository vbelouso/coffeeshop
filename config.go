package main

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Environment string `env:"ENVIRONMENT" envDefault:"local"`
	ServerPort  string `env:"SERVER_PORT" envDefault:":8080"`
	DBDSN       string `env:"DB_DSN" envDefault:"host=localhost port=5432 user=postgres password=coffeeshop dbname=coffeeshop sslmode=disable"`
	//KEYCLOAK_RS256_PUBLIC_KEY string `env:"KEYCLOAK_RS256_PUBLIC_KEY,required"`
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
