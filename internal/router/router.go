// Package router 负责注册路由与中间件。
package router

import (
	"<%= modulePath %>/internal/handler"
	"<%= modulePath %>/internal/middleware"

	"github.com/gin-gonic/gin"
)

// New 构建一个挂载了全部路由与中间件的 Gin 引擎。
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
