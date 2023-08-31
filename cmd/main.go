package main

import (
	log "github.com/sirupsen/logrus"
	"segmenatationService/internal/app"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
}

const configPath = "config/config.yaml"

// @title           Segmentation Service API
// @version         1.0
// @description     This is a Segmentation Service that can create and delete segments, assign them to users and get active segments for user

// @host      localhost:8080
// @BasePath  /api/v1
func main() {
	app.Run(configPath)
}
