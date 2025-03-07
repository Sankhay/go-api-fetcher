package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/Sankhay/go-api-fetcher/models"
)

var openWeatherApiKey string

func Init() {

	openWeatherApiKey = os.Getenv("OPEN_WEATHER_API_KEY")

	if openWeatherApiKey == "" {
		log.Fatal("OPEN_WEATHER_API_KEY in .env is required to run the application.")
	}
}

func getCityWeatherByNameService(cityName string) (*WeatherResponse, *models.NetworkError) {

	fmt.Println(openWeatherApiKey)

	link := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", cityName, openWeatherApiKey)

	fmt.Println(link)

	resp, err := http.Get(link)

	if err != nil {
		log.Print(err)
		var networkError models.NetworkError = models.NetworkError{Code: 500, Msg: err.Error()}
		return nil, &networkError
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Print(err)
		var networkError models.NetworkError = models.NetworkError{Code: 500, Msg: err.Error()}
		return nil, &networkError
	}

	if resp.StatusCode != http.StatusOK {

		var bodyError struct {
			Message string `json:"message"`
		}

		if err := json.Unmarshal(body, &bodyError); err != nil {
			log.Print(err)
			var networkError models.NetworkError = models.NetworkError{Code: 500, Msg: err.Error()}
			return nil, &networkError
		}

		var networkError models.NetworkError = models.NetworkError{Code: resp.StatusCode, Msg: bodyError.Message}
		return nil, &networkError
	}

	var weatherResponse WeatherResponse

	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		log.Print(err)
		var networkError models.NetworkError = models.NetworkError{Code: resp.StatusCode, Msg: err.Error()}
		return nil, &networkError
	}

	return &weatherResponse, nil

}
