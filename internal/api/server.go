package api

import (
	"context"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"segmenatationService/config"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(router *gin.Engine, config *config.Config) error {
	s.httpServer = &http.Server{
		Addr:           ":" + config.Server.Port,
		Handler:        router,
		ReadTimeout:    time.Duration(config.Server.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(config.Server.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Infof("Server is running on port %s", config.Server.Port)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	log.Info("Stopping server")
	s.httpServer.Shutdown(ctx)
}
