package middlewares

import (
    "GOLANG_CLEAN_WEB_API/src/api/helper"
    "GOLANG_CLEAN_WEB_API/src/config"
    "GOLANG_CLEAN_WEB_API/src/constants"
    "GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
    "GOLANG_CLEAN_WEB_API/src/services"
    "fmt"
    "net/http"
    "strings"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
    "errors"
)

func Authentication(cfg *config.Config) gin.HandlerFunc {
    var tokenService = services.NewTokenService(cfg)
    return func(c *gin.Context) {
        claimMap := map[string]interface{}{}
        auth := c.GetHeader(constants.AuthorizationHeaderKey)
        token := strings.Split(auth, " ")
        var err error
        if auth == "" {
            err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenRequired}
        } else {
            claimMap, err = tokenService.GetClaims(token[1])
            if err != nil {
                if errors.Is(err, jwt.ErrTokenExpired) {
                    err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenExpired}
                } else {
                    err = &service_errors.ServiceError{EndUserMessage: service_errors.TokenInvalid}
                }
            }
        }
        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, helper.GenerateBaseResponseWithError(
                nil, false,helper.AuthError, err,
            ))
            return
        }
        c.Set(constants.UserIdKey, claimMap[constants.UserIdKey])
        c.Set(constants.FirstNameKey, claimMap[constants.FirstNameKey])
        c.Set(constants.LastNameKey, claimMap[constants.LastNameKey])
        c.Set(constants.UserNameKey, claimMap[constants.UserNameKey])
        c.Set(constants.EmailKey, claimMap[constants.EmailKey])
        c.Set(constants.MobileNumberKey, claimMap[constants.MobileNumberKey])
        c.Set(constants.RoleKey, claimMap[constants.RoleKey])
        c.Set(constants.ExpireTimeKey, claimMap[constants.ExpireTimeKey])
        c.Next()
    }
}


func Authorization(validRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Keys) == 0 {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false, helper.ForbiddenError))
			return

		}
		roleVal := c.Keys[constants.RoleKey]
		fmt.Println(roleVal)
		if roleVal == nil {
			c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false,  helper.ForbiddenError))
			return
		}
		role := roleVal.([]interface{})
		val := map[string]int{}
		for _, item := range role {
			val[item.(string)] = 0

		}
		for _, item := range validRoles {
			if _, ok := val[item]; ok {
				c.Next()
				return
			}
		}
		c.AbortWithStatusJSON(http.StatusForbidden, helper.GenerateBaseResponse(nil, false,  helper.ForbiddenError))

	}
}
