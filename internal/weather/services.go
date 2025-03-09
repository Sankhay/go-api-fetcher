package weather

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/Sankhay/go-api-fetcher/models"
)

func getCityWeatherByNameServices(cityName string) (*Weather, *models.HttpError) {

	link := fmt.Sprintf("%s/data/2.5/weather?q=%s&appid=%s&units=metric&lang=pt_br", openWeatherApiLink, cityName, openWeatherApiKey)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, link, nil)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to create request for weather city %s: %v", cityName, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		errMsg := fmt.Sprintf("Request failed for city %s: %v", cityName, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		errMsg := fmt.Sprintf("Failed to read response body for city %s: %v", cityName, err)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: err.Error()}
	}

	if resp.StatusCode != http.StatusOK {

		var bodyError struct {
			Message string `json:"message"`
		}

		if err := json.Unmarshal(body, &bodyError); err != nil {
			errMsg := fmt.Sprintf("Failed to parse error response for city %s: %v", cityName, err)
			log.Println(errMsg)
			return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
		}

		errMsg := fmt.Sprintf("API returned error for city %s (Status %d): %s", cityName, resp.StatusCode, bodyError.Message)
		log.Println(errMsg)
		return nil, &models.HttpError{Code: resp.StatusCode, Msg: errMsg}
	}

	var weather Weather

	if err := json.Unmarshal(body, &weather); err != nil {
		errMsg := fmt.Sprintf("Failed to parse json response for city %s: %v", cityName, err)
		log.Print(errMsg)
		return nil, &models.HttpError{Code: http.StatusInternalServerError, Msg: errMsg}
	}

	return &weather, nil
}
