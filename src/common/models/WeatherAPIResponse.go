package models

type Weather struct {
	Temperature  float32  `json:"temp"`
}

type TimeStamp struct {
	Dt int  `json:"dt"`
}

type WeatherApiResponse struct {
	MainWeather Weather `json:"main"`
	TimeStamp int `json:"dt"`
}

type ErrorResponse struct {
	Code int `json:"cod"`
	Message string `json:"message"`
}