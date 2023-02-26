package repositories

import (
	"WeatherApi/src/common"
	"WeatherApi/src/common/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func buildWeatherAPIURL(cityName string, apiKey string) string {
	// Replace spaces in city name with "%20"
	formattedCityName := strings.ReplaceAll(cityName, "-", "%20")

	return fmt.Sprintf(common.API_URL, formattedCityName, apiKey)
}

func GetWeatherByCityName(cityName string) *models.WeatherApiResponse {
	var url string = buildWeatherAPIURL(cityName, common.API_KEY)

	res, err := http.Get(url)

	if err != nil {
		return nil
	}

	defer res.Body.Close()

	var responseData models.WeatherApiResponse
	err = json.NewDecoder(res.Body).Decode(&responseData)

	if err != nil {
		fmt.Print(err)
	}

	//fmt.Print("\n\n", responseData, "\n\n")

	// jsonRes, _ := json.Marshal(responseData)
	// fmt.Print(string(jsonRes), "\n\n")

	return &responseData
}
