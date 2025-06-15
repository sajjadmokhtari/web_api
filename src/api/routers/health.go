package routers

import (
    "GOLANG_CLEAN_WEB_API/src/api/handlers"

    "github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup) {

    // ایجاد یک نمونه از هندلر
    handler := handlers.NewHealthHandler()

    // تعریف مسیرهای مختلف
    r.GET("/", handler.Health) 
                 

}
