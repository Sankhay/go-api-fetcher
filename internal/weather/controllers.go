package weather

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCityWeatherByName(c *gin.Context) {

	cityName := c.Param("city")

	weatherResponse, err := getCityWeatherByNameService(cityName)

	if err != nil {
		c.JSON(err.Code, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, weatherResponse)
}
