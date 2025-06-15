package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestMidlleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		apikey := ctx.GetHeader("x-api-key")
		if apikey == "1" {
			ctx.Next()
		}
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"result": "Api key is required",
		})

		return
	}
}
