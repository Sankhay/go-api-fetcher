package weather

import (
	"net/http"
	"net/http/httptest"
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

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/weather/London", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}
