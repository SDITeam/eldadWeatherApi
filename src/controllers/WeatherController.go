package controllers

import (
	"WeatherApi/src/services"
	"net/http"
	"github.com/gin-gonic/gin"
)

func GetWeatherByCityName(context *gin.Context) {
	weatherData, err := services.GetWeatherByCityName(context.Param("cityName"))

	if err != nil {
		context.String(http.StatusNotFound, err.Error())
	} else {
		context.IndentedJSON(http.StatusOK, weatherData)
	}
}
