package services

import (
	"WeatherApi/src/repositories"
)

func GetWeatherByCityName(cityName string) {
	repositories.GetWeatherByCityName(cityName)
}