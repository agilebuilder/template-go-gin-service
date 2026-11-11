package handler

import (
	"net/http"

	"<%= modulePath %>/internal/service"

	"github.com/gin-gonic/gin"
)

// HelloHandler is the handler for the example business endpoint.
type HelloHandler struct {
	svc *service.HelloService
}

// NewHelloHandler creates a HelloHandler.
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{svc: service.NewHelloService()}
}

// Hello handles GET /api/v1/hello?name=xxx.
func (h *HelloHandler) Hello(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, h.svc.Greet(name))
}
