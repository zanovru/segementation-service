package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "segmenatationService/docs"
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
		v1.POST("/segments", r.createSegment)
		v1.DELETE("/segments/:slug", r.deleteSegmentBySlug)
		v1.POST("/users", r.createUserWithSegments)
		v1.GET("/users/:id/segments", r.getUserSegments)
		v1.POST("/segmentsProb", r.createSegmentWithProbability)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
