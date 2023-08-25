package handlers

import (
	"github.com/gin-gonic/gin"
)

type Routing struct {
}

func NewRouting() *Routing {
	return &Routing{}
}

func (r *Routing) InitRoutes() *gin.Engine {
	router := gin.New()
	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, "Hello World!")
	})
	return router
}
