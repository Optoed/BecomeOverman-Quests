package handlers

import (
	"BecomeOverMan/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BaseHandler struct {
	service *services.BaseService
}

func NewBaseHandler(service *services.BaseService) *BaseHandler {
	return &BaseHandler{service: service}
}

func (h *BaseHandler) CheckConnection(c *gin.Context) {
	if err := h.service.CheckConnection(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "database is available"})
}
