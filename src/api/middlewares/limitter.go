package middlewares

import (
	"GOLANG_CLEAN_WEB_API/src/api/helper"
	"log"
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
)

func LimitByRequest() gin.HandlerFunc {
	lmt := tollbooth.NewLimiter(1, nil)
	return func(c *gin.Context) {
		log.Println("Request received:", c.Request.Method, c.Request.URL.Path)

		err := tollbooth.LimitByRequest(lmt, c.Writer, c.Request)
		if err != nil {
			log.Println("Rate limit triggered:", err)
			c.AbortWithStatusJSON(http.StatusTooManyRequests,
				helper.GenerateBaseResponseWithError(nil, false, -100, err))
			return
		}

		log.Println("Request passed limit check, proceeding to next middleware.")
		c.Next()
	}
}
