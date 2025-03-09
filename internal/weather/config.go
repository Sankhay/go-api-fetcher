package weather

import (
	"log"
	"os"
)

var openWeatherApiKey string

const openWeatherApiLink string = "https://api.openweathermap.org"

func Init() {
	openWeatherApiKey = os.Getenv("OPEN_WEATHER_API_KEY")
	if openWeatherApiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY in .env is required to run the application, and in tests/tests.env to run the tests")
	}
}
