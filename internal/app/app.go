package app

import (
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"segmenatationService/config"
	"segmenatationService/internal/api"
	"segmenatationService/internal/api/handlers"
	"segmenatationService/internal/repository"
	"segmenatationService/internal/repository/postgres"
	"segmenatationService/internal/services"
	"syscall"
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
	go func() {
		if err = server.Start(routing.InitRoutes(), appConfig); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server error: %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	server.Shutdown(context.Background())
}
