package main

import (
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"segmenatationService/internal/app"
)

const configPath = "config/config.yaml"

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("No .env file found %s", err)
	}
}

func main() {
	app.Run(configPath)

}
