package models

type WeatherData struct {
	requestId  string  `json:"report_id"`
	temp float32 `json:"temp"`
	timeStamp string  `json:"time_stamp"`
}