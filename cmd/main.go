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

func main() {
	app.Run(configPath)
}
