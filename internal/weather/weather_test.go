package weather

import (
	"net/http"
	"testing"

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
}

func testGetWeatherByCityNameControllersOK(t *testing.T, r *gin.Engine) {
	w, _ := tests.PerformTestGetRequest(r, "/weather/London")

	assert.Equal(t, http.StatusOK, w.Code)
}
