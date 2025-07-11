package routers

import (
	"GOLANG_CLEAN_WEB_API/src/api/handlers"
	"GOLANG_CLEAN_WEB_API/src/api/middlewares"
	"GOLANG_CLEAN_WEB_API/src/config"

	"github.com/gin-gonic/gin"
)

func User(router *gin.RouterGroup, cfg *config.Config) {
	h := handlers.NewUsersHandler(cfg)
	
	// Route with OTP rate limiting
	router.POST("/send-otp", middlewares.OtpLimiter(cfg), h.SendOtp)
	router.POST("/login-by-username", h.LoginByUsername)
	router.POST("/register-by-username", h.RegisterByUsername)
	router.POST("/login-by-mobile", h.RegisterLoginByMobileRequest)
}
