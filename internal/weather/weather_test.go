package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/Sankhay/go-api-fetcher/models"
	"github.com/Sankhay/go-api-fetcher/tests"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAll(t *testing.T) {
	t.Run("GetWeatherByCityNameControllers", TestGetWeatherByCityNameControllers)
}

func TestGetWeatherByCityNameControllers(t *testing.T) {
	r := tests.SetupTestMode()

	Init()

	r.GET("/weather/:city", GetCityWeatherByNameControllers)
	t.Run("OK", func(t *testing.T) {
		testGetWeatherByCityNameControllersOK(t, r)
	})

	t.Run("CityNotFound", func(t *testing.T) {
		testGetWeatherByCityNameControllersCityNotFound(t, r)
	})
}

func testGetWeatherByCityNameControllersOK(t *testing.T, r *gin.Engine) {
	w, _ := tests.PerformTestGetRequest(r, "/weather/London")

	assert.Equal(t, http.StatusOK, w.Code)
}

func testGetWeatherByCityNameControllersCityNotFound(t *testing.T, r *gin.Engine) {
	testCityName := "testCity"
	w, body := tests.PerformTestGetRequest(r, fmt.Sprintf(`/weather/%s`, testCityName))

	httpError := models.HttpError{Code: http.StatusNotFound, Msg: fmt.Sprintf(`API returned error for city %s (Status 404): city not found`, testCityName)}
	httpErrorJson, _ := json.Marshal(httpError)
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.JSONEq(t, string(httpErrorJson), body)
}
