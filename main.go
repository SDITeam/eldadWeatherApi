package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"WeatherApi/src/routes"
	"WeatherApi/src/common"
)

func main() {
	router := gin.Default()
	routes.Weather(router.Group("/weather"))
	router.Run(":" + common.PORT)
	fmt.Print("Server running on port ", common.PORT)
}