// Package handler 存放 HTTP 处理器（控制器层）。
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler 提供健康检查接口。
type HealthHandler struct{}

// NewHealthHandler 创建 HealthHandler。
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check 处理 GET /healthz。
func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
