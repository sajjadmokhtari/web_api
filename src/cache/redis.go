package cache

import (
	"context"
	"encoding/json"
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



func Set[T any](c *redis.Client, key string, value T, duration time.Duration) error {
    v, err := json.Marshal(value)
    if err != nil {
        return err
    }
    ctx := context.Background()//جدید نوشتم توی فیلم پشتیبانی نمیکرد 
    return c.Set(ctx, key, v, duration).Err()
}



func Get[T any](c *redis.Client, key string) (T, error) {
    var dest T = *new(T)
    ctx := context.Background()//توی فیلم این نیست اضافه کردمش

    v, err := c.Get(ctx, key).Result()
    if err != nil {
        return dest, err
    }
    err = json.Unmarshal([]byte(v), &dest) // توجه کن: &dest به جای dest
    if err != nil {
        return dest, err
    }
    return dest, nil
}

