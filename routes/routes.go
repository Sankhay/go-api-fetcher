package routes

import (
	"github.com/Sankhay/go-api-fetcher/internal/user"
	"github.com/Sankhay/go-api-fetcher/internal/weather"
	"github.com/gin-gonic/gin"
)

func StartRoutes(r *gin.Engine) {
	r.GET("/api/weather/:city", weather.GetCityWeatherByNameControllers)
	r.GET("/api/user/:id", user.GetUserByIdControllers)
}
