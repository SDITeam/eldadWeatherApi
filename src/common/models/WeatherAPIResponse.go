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
