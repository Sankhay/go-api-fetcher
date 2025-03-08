package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Sankhay/go-api-fetcher/models"
)

func getCityWeatherByNameService(cityName string) (*WeatherResponse, *models.NetworkError) {

	link := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s&units=metric&lang=pt_br", openWeatherApiLink, cityName, openWeatherApiKey)

	resp, err := http.Get(link)

	if err != nil {
		log.Println(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	if resp.StatusCode != http.StatusOK {

		var bodyError struct {
			Message string `json:"message"`
		}

		if err := json.Unmarshal(body, &bodyError); err != nil {
			log.Println(err)
			return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
		}

		return nil, &models.NetworkError{Code: resp.StatusCode, Msg: bodyError.Message}
	}

	var weatherResponse WeatherResponse

	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		log.Print(err)
		return nil, &models.NetworkError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	return &weatherResponse, nil

}
