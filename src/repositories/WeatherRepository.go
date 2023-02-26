package repositories

import (
	"WeatherApi/src/common"
	"WeatherApi/src/common/models"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func buildWeatherAPIURL(cityName string, apiKey string) string {
	// Replace spaces in city name with "%20"
	formattedCityName := strings.ReplaceAll(cityName, "-", "%20")

	return fmt.Sprintf(common.API_URL, formattedCityName, apiKey)
}

func GetWeatherByCityName(cityName string) (*models.WeatherApiResponse, error) {
	var url string = buildWeatherAPIURL(cityName, common.API_KEY)
	res, err := http.Get(url)

	if err != nil {
		return nil, errors.New("error in getting weather from api")
	}

	defer res.Body.Close();
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.New("error in reading response body")
	}

	var errorResponse models.ErrorResponse
	err = json.Unmarshal(body, &errorResponse)

	if (errorResponse.Code != 200) {
		return nil, errors.New(errorResponse.Message);
	}

	var responseData models.WeatherApiResponse
	err = json.Unmarshal(body, &responseData)

	if err != nil {
		return nil, errors.New("error in reading api response")
	} 

	return &responseData, nil
}
