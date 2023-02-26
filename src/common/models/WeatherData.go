package models

type WeatherData struct {
	RequestId  string  `json:"report_id"`
	Temperature float32 `json:"temp"`
	TimeStamp string  `json:"time_stamp"`
}