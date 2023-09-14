package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type TemperatureData struct {
	Temperature float64
	CreatedAt   string
}

var temperatureData []TemperatureData

func main() {
	r := gin.Default()

	r.GET("/temperature", func(c *gin.Context) {
		c.JSON(http.StatusOK, temperatureData[len(temperatureData)-1])
	})

	r.GET("/temperatures", func(c *gin.Context) {
		c.JSON(http.StatusOK, temperatureData)
	})

	r.POST("/temperature", func(c *gin.Context) {
		var json TemperatureData
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		temperatureData = append(temperatureData, json)
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
