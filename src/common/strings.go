package common

import (
	"GOLANG_CLEAN_WEB_API/src/config"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdefghijklmnopqrstuvwxyz"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
)


func GenerateOtp() string {
    // دریافت تنظیمات برنامه شامل تعداد رقم‌های OTP
    cfg := config.GetConfig()

    // محاسبه کمترین عدد ممکن با توجه به تعداد رقم‌ها (مثلاً اگر ۶ رقم باشد، min برابر 100000 است)
    min := int(math.Pow(10, float64(cfg.Otp.Digits-1)))

    // محاسبه بیشترین عدد ممکن (مثلاً اگر ۶ رقم باشد، max برابر 999999 است)
    max := int(math.Pow(10, float64(cfg.Otp.Digits)) - 1)

    // ایجاد منبع تصادفی براساس زمان فعلی برای افزایش تصادفی بودن
    src := rand.NewSource(time.Now().UnixNano())
    r := rand.New(src)

    // تولید عدد تصادفی در بازه [min, max]
    num := r.Intn(max-min+1) + min

    // تبدیل عدد تولیدشده به رشته و بازگرداندن آن
    return strconv.Itoa(num)
}




func GeneratePassword() string {
	var password strings.Builder
	cfg := config.GetConfig()

	lowerCharSet := "abcdefghijklmnopqrstuvwxyz"
	upperCharSet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet := "!@#$%&*"
	numberSet := "0123456789"
	allCharSet := ""

	if cfg.Password.IncludeLowercase {
		allCharSet += lowerCharSet
	}
	if cfg.Password.IncludeUppercase {
		allCharSet += upperCharSet
	}
	if cfg.Password.IncludeChars {
		allCharSet += specialCharSet
	}
	if cfg.Password.IncludeDigits {
		allCharSet += numberSet
	}

	passwordLength := cfg.Password.MinLength
	if cfg.Password.MaxLength > cfg.Password.MinLength {
		passwordLength = cfg.Password.MaxLength
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < passwordLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteByte(allCharSet[random])
	}

	return password.String()
}

/*
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

}*/
