package controllers

import (
	"log"
	"net/http"

	"github.com/Sankhay/go-api-fetcher/api"
	"github.com/gin-gonic/gin"
)

func GetCityWeatherByName(c *gin.Context) {

	cityName := c.Param("city")

	weatherResponse, err := api.GetCityWeatherByName(cityName)

	if err != nil {
		log.Fatalf(err.Error())
	}

	c.JSON(http.StatusOK, weatherResponse)
}
