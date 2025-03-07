package config

import (
	"log"

	"github.com/Sankhay/go-api-fetcher/internal/weather"
	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	weather.Init()
}
