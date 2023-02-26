package controllers

import (
	"github.com/gin-gonic/gin"
	"WeatherApi/src/services"
)

func GetWeatherByCityName(context *gin.Context) {
	services.GetWeatherByCityName(context.Param("cityName"))
}