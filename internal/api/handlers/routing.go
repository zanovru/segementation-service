package handlers

import (
	"github.com/gin-gonic/gin"
	"segmenatationService/internal/services"
)

type Routing struct {
	service *services.Service
}

func NewRouting(service *services.Service) *Routing {
	return &Routing{service: service}
}

func (r *Routing) InitRoutes() *gin.Engine {
	router := gin.New()
	v1 := router.Group("/api/v1")
	{
		v1.POST("/segment", r.createSegment)
		v1.DELETE("/segment/:slug", r.deleteSegmentBySlug)
		//v1.POST("/user", r.addUser)
		//v1.GET("/user/:userId", r.getUser)
	}

	return router
}
