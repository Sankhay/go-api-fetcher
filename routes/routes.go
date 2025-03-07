package routes

import (
	"github.com/Sankhay/go-api-fetcher/controllers"
	"github.com/gin-gonic/gin"
)

func StartRoutes(r *gin.Engine) {
	r.GET("/api/weather/:city", controllers.GetCityWeatherByName)
	r.GET("/api/user/:id", controllers.GetUserById)
}
