package middlewares

import (
	"github.com/gin-gonic/gin"
)

func TestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// token := c.GetHeader("token")

		// if token == "" {
		// 	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		// 		"error": "why are you RUNNING",
		// 	}) //
		// }
		c.Next()
	}
}
