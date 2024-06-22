package middlewares

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/meiti-x/golang-web-api/api/helper"
)

func RateLimitMiddleware() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)

	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusTooManyRequests, helper.GenerateBaseResponseWithError(nil, false, -100, err))
			return
		} else {
			ctx.Next()
		}
	}
}
