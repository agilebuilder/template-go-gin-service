package handler

import (
	"net/http"

	"<%= modulePath %>/internal/service"

	"github.com/gin-gonic/gin"
)

// HelloHandler 是示例业务接口的处理器。
type HelloHandler struct {
	svc *service.HelloService
}

// NewHelloHandler 创建 HelloHandler。
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{svc: service.NewHelloService()}
}

// Hello 处理 GET /api/v1/hello?name=xxx。
func (h *HelloHandler) Hello(c *gin.Context) {
	name := c.Query("name")
	c.JSON(http.StatusOK, h.svc.Greet(name))
}
