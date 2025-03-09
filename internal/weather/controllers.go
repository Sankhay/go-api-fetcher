package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCityWeatherByNameControllers(c *gin.Context) {

	cityName := c.Param("city")

	weather, networkError := getCityWeatherByNameServices(cityName)

	if networkError != nil {
		c.JSON(networkError.Code, networkError)
		return
	}

	c.JSON(http.StatusOK, weather)
}
