package services

import (
	"WeatherApi/src/common/models"
	"WeatherApi/src/repositories"
)

func GetWeatherByCityName(cityName string) {
	var weatherApiResponse *models.WeatherApiResponse = repositories.GetWeatherByCityName(cityName)

}
