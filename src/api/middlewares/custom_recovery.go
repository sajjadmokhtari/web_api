package middlewares

import (
	"GOLANG_CLEAN_WEB_API/src/api/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context, err any) {
	if typedErr, ok := err.(error); ok {
		httpResponse := helper.GenerateBaseResponseWithError(nil, false, -500, typedErr)
		c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
		return
	}

	httpResponse := helper.GenerateBaseResponseWithAnyError(nil, false, -500, err)
	c.AbortWithStatusJSON(http.StatusInternalServerError, httpResponse)
}
//  باکمی تغییرات نسبت  به فیلم