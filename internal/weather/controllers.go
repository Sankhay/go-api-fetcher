package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCityWeatherByNameControllers(c *gin.Context) {

	cityName := c.Param("city")

	weather, httpError := getCityWeatherByNameServices(cityName)

	if httpError != nil {
		c.JSON(httpError.Code, httpError)
		return
	}

	c.JSON(http.StatusOK, weather)
}
