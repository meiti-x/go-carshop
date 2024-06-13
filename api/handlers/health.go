package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
}

func NewHealthHanlder() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, "working...")
	return
}

func (h *HealthHandler) Health2(c *gin.Context) {
	c.JSON(http.StatusOK, "working2...")
	return
}
