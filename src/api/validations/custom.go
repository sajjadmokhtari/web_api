package validations

import (
	"errors"
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	Property string `json:"property"`
	Tag      string `json:"tag"`
	Value    string `json:"value"`
	Message  string `json:"message"`
}

func GetValidationErrors(err error) *[]ValidationError {

	var validationError []ValidationError
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		for _, err := range err.(validator.ValidationErrors) {
			var el ValidationError
			el.Property = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			validationError = append(validationError, el)
		}
		return &validationError

	}
	return nil

}

func IranianMobileNumberValidator(fld validator.FieldLevel) bool {

	value, ok := fld.Field().Interface().(string)
	if !ok {
		return false
	}

	res, err := regexp.MatchString(`^09(1[0-9]|2[0-2]|3[0-9]|9[0-9])[0-9]{7}$`, value)
	if err != nil {
		log.Print(err.Error())

	}
	return res

}

/*func PasswordValidator(fld validator.FieldLevel ) bool {

	value , ok := fld.Field().Interface().(string)
	if ok {
		fld.Param()
		return false
	}
	return common.CheckPassword(value)
}*/
