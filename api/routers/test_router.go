package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/golang-web-api/api/handlers"
)

func TestRouter(r *gin.RouterGroup) {
	h := handlers.NewTestHandler()

	r.GET("/", h.Test)
	r.GET("/users", h.TestUsers)
	r.GET("/users/:id", h.TestHeader)
	r.POST("/users/:id", h.TestBody)
	r.POST("/file", h.TestFile)
}
