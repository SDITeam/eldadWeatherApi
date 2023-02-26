package repositories

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strings"
	"WeatherApi/src/common"
	"WeatherApi/src/common/models"
)

func buildWeatherAPIURL(cityName string, apiKey string) string {
	// Replace spaces in city name with "%20"
	formattedCityName := strings.ReplaceAll(cityName, "-", "%20")

	return fmt.Sprintf(common.API_URL, formattedCityName, apiKey)
}


func GetWeatherByCityName(cityName string) *models.WeatherData {
	var url string = buildWeatherAPIURL(cityName, common.API_KEY)

	res, err := http.Get(url);

	if err != nil {
		return nil
	}

	defer res.Body.Close()

	var responseMap map[string]interface{}
    err = json.NewDecoder(res.Body).Decode(&responseMap)

	fmt.Print("\n\nBody:\n", responseMap[main])

    if err != nil {
        return nil
    }

	if err != nil {
		return nil
	}

	return nil;
}