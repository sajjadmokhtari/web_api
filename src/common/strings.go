package common

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"math"
	"math/rand"
	"strconv"
	"time"
)

/*func GeneratesOtp() string {
	cfg := config.GetConfig()
	rand.Seed(time.Now().UnixNano())

	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)

	var num = rand.Intn(max-min) + min
	return strconv.Itoa(num)
}*/
func GenerateOtp() string {
	cfg := config.GetConfig()

	min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))
	max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)

	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	num := r.Intn(max-min+1) + min
	return strconv.Itoa(num)
}

/*import "GOLANG_CLEAN_WEB_API/src/config"

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)

func CheckPassword(password string) bool {

	cfg := config.GetConfig()
	if len(password) < cfg.Password.MinLenght {
		return false
	}

	if cfg.Password.IncludeChars && !HasLetter(password) {
		return false
	}

	if cfg.Password.IncludeDigits && !HasDigits(password) {
		return false
	}

	if cfg.Password.IncludeLowercase && !HasLower(password) {
		return false
	}

	if cfg.Password.IncludeUppercase && !HasUpper(password) {
		return false
	}
	return true

}

func GeneratePassword() string {

	var password string.Builder
	cfg := config.GetConfig()
	passwordLenght := cfg.Password.MinLenght + 2
	minSpecialChar := 2
	minNum := 3
	if !cfg.Password.IncludeDigits {

		minNum = 0

	}
	minUpperCase := 3
	if !cfg.Password.IncludeUppercase {
		minUpperCase = 0
	}
	minLowerCase := 3
	if !cfg.Password.IncludeLowercase {
		minUpperCase = 0
	}
}*/
