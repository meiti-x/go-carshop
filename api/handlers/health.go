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

// HealthCheck godoc
// @Summary Health Check
// @Description Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/ [get]
func (h *HealthHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working...", true, http.StatusAccepted))
	return
}

// HealthCheck2 godoc
// @Summary Health Check2
// @Description Another Health Check
// @Tags health
// @Accept  json
// @Produce  json
// @Success 200 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Router /v1/health/2 [get]
// @Security AuthBearer
func (h *HealthHandler) Health2(c *gin.Context) {
	if c.GetHeader("token") == "" {
		c.JSON(http.StatusBadRequest, helper.GenerateBaseResponse("need token...", false, http.StatusBadRequest))

	}
	c.JSON(http.StatusOK, helper.GenerateBaseResponse("working...", true, http.StatusCreated))
	return
}
