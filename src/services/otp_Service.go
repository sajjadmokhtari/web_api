package services

import (
	"GOLANG_CLEAN_WEB_API/src/cache"
	"GOLANG_CLEAN_WEB_API/src/config"
	"GOLANG_CLEAN_WEB_API/src/constans"
	"GOLANG_CLEAN_WEB_API/src/pkg/logging"
	"GOLANG_CLEAN_WEB_API/src/pkg/service_errors"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type OtpService struct {
	logger      logging.Logger
	cfg         *config.Config
	redisClient *redis.Client
}

type OtpDto struct {
	Value string
	Used  bool
}

func NewOtpService(cfg *config.Config) *OtpService {
	logger := logging.NewLogger(cfg)
	redis := cache.GetRedis()
	return &OtpService{logger: logger, cfg: cfg, redisClient: redis}
}

func (s *OtpService) SetOtp(MobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constans.RedisOtpDefaultKey, MobileNumber)
	val := &OtpDto{
		Value: otp,
		Used:  false,
	}

	res, err := cache.Get[OtpDto](s.redisClient, key)
	if err == nil && !res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpExists}

	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	}
	err = cache.Set(s.redisClient, key, val, s.cfg.Otp.ExpireTime*time.Second)
	if err != nil {
		return err
	}
	return nil

}


func (s *OtpService) ValidateOtp(MobileNumber string, otp string) error {
	key := fmt.Sprintf("%s:%s", constans.RedisOtpDefaultKey, MobileNumber)
	res, err := cache.Get[OtpDto](s.redisClient, key)

	if err != nil {

		return err

	} else if err == nil && res.Used {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpUsed}
	} else if err == nil && !res.Used && res.Value != otp {
		return &service_errors.ServiceError{EndUserMessage: service_errors.OtpNotValid}
	} else if err ==
		nil && !res.Used && res.Value == otp {
		res.Used = true
		err = cache.Set(s.redisClient, key, res, s.cfg.Otp.ExpireTime*time.Second)
	}
	return nil

}
