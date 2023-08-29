package app

import (
	log "github.com/sirupsen/logrus"
	"segmenatationService/config"
	"segmenatationService/internal/api"
	"segmenatationService/internal/api/handlers"
	"segmenatationService/internal/repository"
	"segmenatationService/internal/repository/postgres"
	"segmenatationService/internal/services"
)

func Run(configPath string) {
	appConfig, err := config.Init(configPath)
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	db, err := postgres.NewPostgresDB(appConfig)
	if err != nil {
		log.Fatalf("Error occured when connecting to PostgreSQL Database: %s", err)
	}

	repos := repository.NewRepository(db)

	service := services.NewService(repos)

	routing := handlers.NewRouting(service)

	server := api.NewServer()
	if err = server.Start(routing.InitRoutes(), appConfig); err != nil {
		log.Fatalf("Server error: %s", err)
	}
}
