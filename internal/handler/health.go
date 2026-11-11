// Package handler holds HTTP handlers (controller layer).
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthHandler provides the health check endpoint.
type HealthHandler struct{}

// NewHealthHandler creates a HealthHandler.
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// Check handles GET /healthz.
func (h *HealthHandler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
