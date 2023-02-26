package routes

import (
	"github.com/gin-gonic/gin"
	"WeatherApi/src/controllers"
)

func Weather(group *gin.RouterGroup) {
	group.GET("/:cityName", controllers.GetWeatherByCityName)
}