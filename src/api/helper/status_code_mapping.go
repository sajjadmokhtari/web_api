package helper

import (
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"net/http"
)

var StatusCodeMapping = map[string]int{
	//otp
	service_errors.OtpExists:   409,
	service_errors.OtpUsed:     409,
	service_errors.OtpNotValid: 400,
	
}
func TranslateErrorStatusCode(err error)int {
	value , ok :=StatusCodeMapping[err.Error()]
	if!ok {
		return http.StatusInternalServerError
	}
	return value
}