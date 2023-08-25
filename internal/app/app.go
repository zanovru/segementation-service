package app

import (
	log "github.com/sirupsen/logrus"
	"segmenatationService/config"
	"segmenatationService/internal/api"
	"segmenatationService/internal/api/handlers"
)

func Run(configPath string) {
	config, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	routing := handlers.NewRouting()
	server := api.NewServer()
	if err = server.Start(routing.InitRoutes(), config); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
