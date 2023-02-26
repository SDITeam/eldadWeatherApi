package controllers

import (
	"WeatherApi/src/services"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeatherByCityName(context *gin.Context) {
	weatherData, err := services.GetWeatherByCityName(context.Param("cityName"))

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, err)
	} else {
		jsonRes, _ := json.Marshal(weatherData)
		fmt.Print(string(jsonRes), "\n\n")
		context.IndentedJSON(http.StatusOK, weatherData)
	}
}
