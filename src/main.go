package main

import (
	"time"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	log.WithFields(log.Fields{
		"current_time": time.Now(),
		"name":         "Paweł",
		"surname":      "Czerwieniec",
		"server_port":  80,
	}).Info("Starting info") // Log informacyjny (Zadanie 1a)

	timeService := NewTimeService()

	r := gin.Default()
	r.GET("/", timeService.CheckUserIP) // Definicja podstawowego endpointu (Zadanie 1b)

	err := r.Run("0.0.0.0:80") // Uruchomienie web-servera
	if err != nil {
		return
	}
}
