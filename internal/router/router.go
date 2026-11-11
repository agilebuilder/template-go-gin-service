// Package router registers routes and middleware.
package router

import (
	"<%= modulePath %>/internal/handler"
	"<%= modulePath %>/internal/middleware"

	"github.com/gin-gonic/gin"
)

// New builds a Gin engine with all routes and middleware mounted.
func New() *gin.Engine {
	r := gin.New()
	r.Use(middleware.Logger(), middleware.Recovery())

	health := handler.NewHealthHandler()
	r.GET("/healthz", health.Check)

	hello := handler.NewHelloHandler()
	api := r.Group("/api/v1")
	{
		api.GET("/hello", hello.Hello)
	}

	return r
}
