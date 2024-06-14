package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meiti-x/golang-web-api/api/helper"
)

type HealthHandler struct {
}

func NewHealthHanlder() *HealthHandler {
	return &HealthHandler{}
}

func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working...", true, http.StatusAccepted))
	return
}

func (h *HealthHandler) Health2(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working...", true, http.StatusCreated))
	return
}
