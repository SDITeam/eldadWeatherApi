package services

import (
	"WeatherApi/src/common"
	"WeatherApi/src/common/models"
	"WeatherApi/src/repositories"
	"time"

	guuid "github.com/google/uuid"
)

func kelvinToCelsius(kelvinTemp float32) float32 {
	return kelvinTemp - common.KELVIN_TO_CELSIUS_TEMP
}

func GetWeatherByCityName(cityName string) (*models.WeatherData, error) {
	weatherApiResponse, err := repositories.GetWeatherByCityName(cityName)

	if err != nil {
		return nil, err
	}

	return &models.WeatherData{
		RequestId:   "weather-" + guuid.New().String(),
		Temperature: kelvinToCelsius(weatherApiResponse.MainWeather.Temperature),
		TimeStamp:   time.Unix(int64(weatherApiResponse.TimeStamp), 0).Format(time.RFC3339),
	}, nil
}
