package middlewares

import (
	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func RateLimitMiddleware() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)

	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)

		if err != nil {
			ctx.AbortWithStatusJSON(429, gin.H{
				"error": err.Error(),
			})
			return
		} else {
			ctx.Next()
		}
	}
}
