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
	}).Info("Starting info")

	timeService := NewTimeService()

	r := gin.Default()
	r.GET("/", timeService.CheckUserIP)
	err := r.Run("0.0.0.0:80")
	if err != nil {
		return
	}
}
