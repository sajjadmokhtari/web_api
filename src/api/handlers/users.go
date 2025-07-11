package handlers

import (
	"GOLANG_CLEAN_WEB_API/src/api/dto"
	"GOLANG_CLEAN_WEB_API/src/api/helper"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	service *services.UserService
}

func NewUsersHandler(cfg *config.Config) *UsersHandler {
	service := services.NewUserService(cfg)
	return &UsersHandler{service: service}
}

// LoginByUsername godoc
// @Summary Login by username
// @Description Login by username
// @Tags users
// @Accept json
// @Produce json
// @param Request body dto.LoginByUsernameRequest true "LoginByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /users/login-by-username [post]
func (h *UsersHandler) LoginByUsername(c *gin.Context) {

	req := new(dto.LoginByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	token, err := h.service.LoginByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return

	}
	// call internal sms.service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))

}

// RegisterByUsername godoc
// @Summary Register by username
// @Description Register by username
// @Tags users
// @Accept json
// @Produce json
// @param Request body dto.RegisterUserByUsernameRequest true "RegisterUserByUsernameRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /users/register-by-username [post]
func (h *UsersHandler) RegisterByUsername(c *gin.Context) {

	req := new(dto.RegisterUserByUsernameRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.RegisterByUsername(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return

	}
	// call internal sms.service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))

}

// RegisterLoginByMobileRequest godoc
// @Summary Register login by mobile request
// @Description Register login by mobile request
// @Tags users
// @Accept json
// @Produce json
// @param Request body dto.RegisterLoginByMobileRequest true "RegisterLoginByMobileRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router /users/login-by-mobile [post]
func (h *UsersHandler) RegisterLoginByMobileRequest(c *gin.Context) {
	fmt.Println("[Start] هندل LoginByMobileRequest دریافت شد")

	req := new(dto.RegisterLoginByMobileRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		fmt.Println("[Error] BindJSON شکست خورد:", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	fmt.Printf("[Bind] mobile=%s, otp=%s\n", req.MobileNumber, req.Otp)

	token, err := h.service.RegisterLoginByMobileRequest(req)
	if err != nil {
		fmt.Println("[Error] RegisterLoginByMobileRequest شکست خورد:", err)
		c.AbortWithStatusJSON(
			helper.TranslateErrorStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return
	}

	fmt.Println("[Success] توکن تولید شد:", token)
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(token, true, 0))
}

// SendOtp godoc
// @Summary Send otp to user
// @Description Send otp to user
// @Tags users
// @Accept json
// @Produce json
// @param Request body dto.GetOtpRequest true "GetOtpRequest"
// @Success 201 {object} helper.BaseHttpResponse "Success"
// @Failure 400 {object} helper.BaseHttpResponse "Failed"
// @Failure 409 {object} helper.BaseHttpResponse "Failed"
// @Router  /users/send-otp [post]
func (h *UsersHandler) SendOtp(c *gin.Context) {
	req := new(dto.GetOtpRequest)
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, helper.GenerateBaseResponseWithValidationError(nil, false, -1, err))
		return
	}
	err = h.service.SendOtp(req)
	if err != nil {
		c.AbortWithStatusJSON(
			helper.TranslateErrorStatusCode(err),
			helper.GenerateBaseResponseWithError(nil, false, -1, err),
		)
		return

	}
	// call internal sms.service
	c.JSON(http.StatusCreated, helper.GenerateBaseResponse(nil, true, 0))
} // SendOtp: این هندلر یک کد تِک‌اُتی‌پی برای کاربر تولید می‌کنه، اعتبار درخواست رو بررسی می‌کنه،
// و بعد از موفقیت، پاسخ مناسبی برمی‌گردونه. در صورت خطا، وضعیت مناسب با پیام خطا ارسال می‌شه.
