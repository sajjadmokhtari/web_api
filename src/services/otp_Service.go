package services

import (
	"GOLANG_CLEAN_WEB_API/src/cache"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constants"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

// OtpService ساختار اصلی سرویس OTP شامل logger، تنظیمات و Redis
type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

// OtpDto ساختار داده‌ای که در Redis ذخیره می‌شود
type OtpDto struct {
	Value string // مقدار کد OTP
	Used  bool   // وضعیت استفاده‌شدن OTP
}

// NewOtpService سازنده‌ی سرویس OTP، logger و redis client را مقداردهی می‌کند
func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg) // راه‌اندازی لاگر
	redis := cache.GetRedis()        // اتصال به Redis
	return &OtpService{logger: logger, cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(MobileNumber string, otp string) error {
	// تولید کلید ذخیره‌سازی براساس شماره موبایل
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, MobileNumber)

	//  Redisآماده‌سازی داده برای ذخیره در
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	// تلاش برای دریافت OTP قبلی از Redis
	res, err := cache.Get[OtpDto](s.redisClient, key)

	// اگر او تی پی   وجود دارد و استفاده نشده، خطای مربوطه را برمی‌گرداند
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}
		// اگر اوتیپی قبلی استفاده شده، خطای مربوطه را برمی‌گرداند
	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}

	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}

	return nil
}

// ValidateOtp اعتبارسنجی کد OTP واردشده برای شماره موبایل مشخص‌شده
func (s *OtpService) ValidateOtp(MobileNumber string, otp string) error {
	// تولید کلید برای بازیابی OTP از Redis
	key := fmt.Sprintf("%s:%s", constants.RedisOtpDefaultKey, MobileNumber)

	// دریافت داده‌ی ذخیره‌شده از Redis
	res, err := cache.Get[OtpDto](s.redisClient, key)

	if err != nil {
		// اگر خطا در دریافت اطلاعات باشد، همان را برمی‌گرداند
		return err
	} else if res.Used {
		// اگر OTP قبلاً استفاده شده باشد
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if res.Value != otp {
		// اگر OTP اشتباه باشد
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else {
		// اگر اوتیپی درست باشد و هنوز استفاده نشده، آن را به عنوان استفاده‌شده علامت‌گذاری کرده و مجدد ذخیره می‌کند
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
	}

	return nil
}
