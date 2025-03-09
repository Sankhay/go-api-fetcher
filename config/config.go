package config

import (
	"log"

	"github.com/Sankhay/go-api-fetcher/internal/weather"
	"github.com/joho/godotenv"
)

const Port = "8000"

func Init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	weather.Init()
}
