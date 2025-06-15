package cache

import (
	"context"
	"fmt"
	"time"

	"GOLANG_CLEAN_WEB_API/src/config"

	"github.com/go-redis/redis/v8"
)

var redisClient *redis.Client

func InitRedis(cfg *config.Config) error {
	redisClient = redis.NewClient(&redis.Options{
		Addr:               fmt.Sprintf("%s:%s", cfg.Redis.Host, cfg.Redis.Port), // استفاده از ":" به جای ","
		Password:           cfg.Redis.Password,
		DB:                 0,
		DialTimeout:        cfg.Redis.DialTimeout * time.Second,
		ReadTimeout:        cfg.Redis.ReadTimeout * time.Second,
		WriteTimeout:       cfg.Redis.WriteTimeout * time.Second, // اصلاح نام فیلد (تصحیح اگر نیاز است)
		PoolSize:           cfg.Redis.PoolSize,
		IdleTimeout:        500 * time.Millisecond,                          // از میلی‌ثانیه استفاده شده؛ در صورت نیاز تغییر دهید
		IdleCheckFrequency: cfg.Redis.IdleCheckFrequency * time.Millisecond, // تنظیم واحد زمانی مناسب
	})

	// استفاده از context برای فراخوانی متد Ping
	ctx := context.Background()
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetRedis() *redis.Client {
	return redisClient
}

func CloseRedis() {
	if err := redisClient.Close(); err != nil {
		fmt.Println("Error closing Redis:", err)
	}
}
