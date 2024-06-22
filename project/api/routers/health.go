package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/golang-web-api/api/handlers"
)

func Health(r *gin.RouterGroup) {
	handler := handlers.NewHealthHanlder()

	r.GET("/", handler.Health)
	r.GET("2", handler.Health2)
}
